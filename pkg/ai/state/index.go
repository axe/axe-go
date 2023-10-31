package state

import (
	"sort"

	"github.com/axe/axe-go/pkg/id"
)

// The area for state machine identifiers - to keep maps in this package
// as small as fast as possible.
var Area = id.NewArea[uint32, uint16]()

// A transition is taken when the condition returns true.
type Condition[Input any] func(Input) bool

// The amount a state effects a subject based on input can be evaluated
// at state creation or on each update. The effect can be a simple weight
// or a complex object.
type Effector[Input any, Effect any] func(Input) Effect

// A transition in a state machine has an ending and optionally a start.
// A transition without a start is considered global and exists on the machine
// level (as opposed to state) and is possibly evaluated on init or when the
// machine has no active or queued states. A transition can be live or dormant.
// A dormant transition is only evaluated when a state is finished or a machine
// is empty while a live transition is evaluated on each update.
type Transition[Input any, TransitionData any] struct {
	// The start state if any.
	Start id.Identifier
	// The ending state.
	End id.Identifier
	// The condition that determines whether we should take this transition.
	// This is optional and when not specified it's considered to be true.
	Condition func(Input) bool
	// If the transition should be evaluated on each update.
	Live bool
	// The data associated with this transition - used for applying it to the subject.
	Data TransitionData
}

// Based on the input - should the transition be taken?
func (t Transition[Input, TransitionData]) IsReady(input Input) bool {
	return t.Condition == nil || t.Condition(input)
}

// A way to control when (if at all) a states sub-machine is initialized.
type InitializeSub int

const (
	// Initialize the sub machine before we even call Start.
	InitializeSubBeforeStart InitializeSub = iota
	// Initialize the sub after start is called and returns true.
	InitializeSubAfterStart
	// Initialze the sub once we active the state on the machine.
	InitializeSubAfterQueue
	// Never initialize the sub, behave normally and let the first Update handle it.
	InitializeSubNever
)

// The logic and options used by a state machine to process states.
// This allows for fuzzy and finite machines to coexist and also provides
// prioritized multiple state reduction for both finite and fuzzy machines.
// A state has data or it's own sub machine.
type Logic[Subject any, Input any, StateData any, TransitionData any, Effect any] struct {
	// The logic that would return true if the given state is done and the non-live transitions out of it should be evaluted for the next state.
	IsDone func(*Subject, State[Subject, Input, StateData, TransitionData, Effect]) bool
	// Starts the given state on the subject, possible from a transition and possible where an unfinished state needs to be outroed so that it may eventually return true for IsDone and
	// be removed from the active states.
	Start func(*Subject, State[Subject, Input, StateData, TransitionData, Effect], *Transition[Input, TransitionData], *State[Subject, Input, StateData, TransitionData, Effect]) bool
	// The logic invoked to apply the active states on the subject. See other Applied properties.
	Apply func(*Subject, []State[Subject, Input, StateData, TransitionData, Effect])
	// If non-zero, no more than this many states should be sent to be applied to the subject.
	AppliedMax int
	// If non-nil and there's an applied max that will affect the applied states - this function will sort the states so the preferred states are applied.
	AppliedPriority func(a State[Subject, Input, StateData, TransitionData, Effect], b State[Subject, Input, StateData, TransitionData, Effect]) bool
	// If true all active states in all sub machines should be passed to apply - otherwise just the active states in the root machine are sent.
	AppliedDeep bool
	// If true only this number of states can be active at a time. Once a queue of possible active states is determined from transitions
	// only the states that can fit will be added to the active states.
	ActiveMax int
	// If non-nil and there's an active max this will sort the states in the queue to prefer the ones at the top.
	ActivePriority func(a State[Subject, Input, StateData, TransitionData, Effect], b State[Subject, Input, StateData, TransitionData, Effect]) bool
	// If true all states in the machine are always active and no transition logic is done. Only state updating.
	ActiveFully bool
	// If true when active states are removed - should the current order of the active states be preserved or does order not matter?
	ActiveOrdered bool
	// If true after transitions are evaluated in an update - should we add any to the active states so they are available in the apply?
	// If not the active queue will be processed at the start of the next update.
	ProcessQueueImmediately bool
	// Timing for when (if at all) a sub machine should be initialized (active states be determined).
	InitializeSub InitializeSub
}

// A machine definition describes the logic, states, and transitions for a machine.
type MachineDefinition[Subject any, Input any, StateData any, TransitionData any, Effect any] struct {
	Logic       Logic[Subject, Input, StateData, TransitionData, Effect]
	States      id.DenseMap[StateDefinition[Subject, Input, StateData, TransitionData, Effect], uint16, uint8]
	Transitions []Transition[Input, TransitionData]
}

func NewMachineDefinition[Subject any, Input any, StateData any, TransitionData any, Effect any](logic Logic[Subject, Input, StateData, TransitionData, Effect]) MachineDefinition[Subject, Input, StateData, TransitionData, Effect] {
	type MachineDefinitionType = MachineDefinition[Subject, Input, StateData, TransitionData, Effect]
	type StateDefinitionType = StateDefinition[Subject, Input, StateData, TransitionData, Effect]

	return MachineDefinitionType{
		Logic:       logic,
		States:      id.NewDenseMap[StateDefinitionType, uint16, uint8](id.WithArea(Area)),
		Transitions: make([]Transition[Input, TransitionData], 0),
	}
}

// Add the state to the definition.
func (md MachineDefinition[Subject, Input, StateData, TransitionData, Effect]) AddState(s StateDefinition[Subject, Input, StateData, TransitionData, Effect]) {
	md.States.Set(s.ID, s)
}

// Add the transition to the definition. This should only be done after the states are added, otherwise it will panic.
func (md MachineDefinition[Subject, Input, StateData, TransitionData, Effect]) AddTransition(t Transition[Input, TransitionData]) {
	if !md.States.Has(t.End) {
		panic("machine definition has no end state for transition")
	}
	if t.Start.Exists() {
		if !md.States.Has(t.Start) {
			panic("machine definition has no start state for transition")
		}
		start := md.States.Get(t.Start)
		start.Transitions = append(start.Transitions, t)
	} else {
		md.Transitions = append(md.Transitions, t)
	}
}

// A state definition is identifiable, has data, may have a way to compute it's effect (or has a constant effect),
// has transitions out of it, and possibly a sub-machine.
type StateDefinition[Subject any, Input any, StateData any, TransitionData any, Effect any] struct {
	ID             id.Identifier
	Data           StateData
	EffectGet      func(Input) Effect
	EffectConstant Effect
	EffectLive     bool
	Transitions    []Transition[Input, TransitionData]
	Sub            *MachineDefinition[Subject, Input, StateData, TransitionData, Effect]
}

func (d StateDefinition[Subject, Input, StateData, TransitionData, Effect]) GetEffect(input Input) Effect {
	if d.EffectGet == nil {
		return d.EffectConstant
	}
	return d.EffectGet(input)
}
func (d StateDefinition[Subject, Input, StateData, TransitionData, Effect]) IsEffectLive() bool {
	return d.EffectGet != nil && d.EffectLive
}

type State[Subject any, Input any, StateData any, TransitionData any, Effect any] struct {
	Definition *StateDefinition[Subject, Input, StateData, TransitionData, Effect]
	Effect     Effect
	Sub        *Machine[Subject, Input, StateData, TransitionData, Effect]
}

func (a *State[Subject, Input, StateData, TransitionData, Effect]) Update(subject *Subject, input Input) {
	if a.Sub != nil {
		a.Sub.Update(subject, input)
	} else if a.Definition.IsEffectLive() {
		a.Effect = a.Definition.EffectGet(input)
	}
}

func (a State[Subject, Input, StateData, TransitionData, Effect]) IsDone(subject *Subject, logic Logic[Subject, Input, StateData, TransitionData, Effect]) bool {
	if a.Sub != nil {
		if len(a.Sub.ActiveQueue) > 0 {
			return false
		}
		for _, subState := range a.Sub.Active.Values() {
			if !subState.IsDone(subject, logic) {
				return false
			}
		}
		return true
	} else {
		return logic.IsDone != nil && logic.IsDone(subject, a)
	}
}

type Machine[Subject any, Input any, StateData any, TransitionData any, Effect any] struct {
	Definition  *MachineDefinition[Subject, Input, StateData, TransitionData, Effect]
	Active      id.DenseMap[State[Subject, Input, StateData, TransitionData, Effect], uint16, uint8]
	ActiveQueue []State[Subject, Input, StateData, TransitionData, Effect]
	Applicable  []State[Subject, Input, StateData, TransitionData, Effect]
}

func NewMachine[Subject any, Input any, StateData any, TransitionData any, Effect any](def *MachineDefinition[Subject, Input, StateData, TransitionData, Effect]) Machine[Subject, Input, StateData, TransitionData, Effect] {
	type MachineType = Machine[Subject, Input, StateData, TransitionData, Effect]
	type StateType = State[Subject, Input, StateData, TransitionData, Effect]

	return MachineType{
		Definition:  def,
		Active:      id.NewDenseMap[StateType, uint16, uint8](id.WithArea(Area)),
		ActiveQueue: make([]StateType, 0),
		Applicable:  make([]StateType, 0),
	}
}

func (m *Machine[Subject, Input, StateData, TransitionData, Effect]) queueActive(def *StateDefinition[Subject, Input, StateData, TransitionData, Effect], subject *Subject, input Input, trans *Transition[Input, TransitionData], outro *State[Subject, Input, StateData, TransitionData, Effect]) bool {
	logic := m.Definition.Logic

	state := State[Subject, Input, StateData, TransitionData, Effect]{
		Definition: def,
		Effect:     def.GetEffect(input),
	}
	if def.Sub != nil {
		sub := NewMachine(def.Sub)
		state.Sub = &sub
		if logic.InitializeSub == InitializeSubBeforeStart {
			state.Sub.Init(subject, input)
		}
	}
	if logic.Start(subject, state, trans, outro) {
		if logic.InitializeSub == InitializeSubAfterStart && state.Sub != nil {
			state.Sub.Init(subject, input)
		}
		m.ActiveQueue = append(m.ActiveQueue, state)
		return true
	}
	return false
}

func (m *Machine[Subject, Input, StateData, TransitionData, Effect]) Transitions(transitions []Transition[Input, TransitionData], subject *Subject, input Input, onlyLive bool, outro *State[Subject, Input, StateData, TransitionData, Effect]) int {
	transitioned := 0

	for _, trans := range transitions {
		if onlyLive && !trans.Live {
			continue
		}
		active := m.Active.Ptr(trans.End)
		if active != nil {
			continue
		}
		if trans.IsReady(input) {
			def := m.Definition.States.Ptr(trans.End)
			if m.queueActive(def, subject, input, &trans, outro) {
				transitioned++
			}
		}
	}
	return transitioned
}

func (m *Machine[Subject, Input, StateData, TransitionData, Effect]) Init(subject *Subject, input Input) {
	if !m.Active.Empty() || len(m.ActiveQueue) > 0 {
		return
	}
	logic := m.Definition.Logic
	transitions := m.Definition.Transitions
	hasTransitions := len(transitions) > 0

	if logic.ActiveFully && !hasTransitions {
		for _, def := range m.Definition.States.Values() {
			m.queueActive(&def, subject, input, nil, nil)
		}
	} else if hasTransitions {
		m.Transitions(transitions, subject, input, false, nil)
	}
}

func (m *Machine[Subject, Input, StateData, TransitionData, Effect]) ProcessQueue(subject *Subject, input Input) {
	logic := m.Definition.Logic

	if len(m.ActiveQueue) > 0 {
		if logic.ActiveMax > 0 {
			remainingSpace := logic.ActiveMax - m.Active.Len()
			if remainingSpace >= len(m.ActiveQueue) {

			} else if remainingSpace > 0 {
				if logic.ActivePriority != nil {
					sort.Slice(m.ActiveQueue, func(i, j int) bool {
						a := m.ActiveQueue[i]
						b := m.ActiveQueue[j]
						return logic.ActivePriority(a, b)
					})
				}
				m.ActiveQueue = m.ActiveQueue[:remainingSpace]
			} else {
				m.ActiveQueue = m.ActiveQueue[:0]
			}
		}
		if len(m.ActiveQueue) > 0 {
			for _, active := range m.ActiveQueue {
				m.Active.Set(active.Definition.ID, active)

				if logic.InitializeSub == InitializeSubAfterQueue && active.Sub != nil {
					active.Sub.Init(subject, input)
				}
			}
			m.ActiveQueue = m.ActiveQueue[:0]
		}
	}
}

func (m *Machine[Subject, Input, StateData, TransitionData, Effect]) UpdateActive(subject *Subject, input Input) {
	logic := m.Definition.Logic

	if logic.ActiveFully {
		for _, state := range m.Active.Values() {
			state.Update(subject, input)
		}
	} else {
		for _, state := range m.Active.Values() {
			done := state.IsDone(subject, m.Definition.Logic)
			if !done {
				state.Update(subject, input)
			}

			transitioned := m.Transitions(state.Definition.Transitions, subject, input, !done, &state)

			if transitioned > 0 && !done {
				done = true
			}

			if done {
				m.Active.Remove(state.Definition.ID, logic.ActiveOrdered)
			}
		}
	}
}

func (m *Machine[Subject, Input, StateData, TransitionData, Effect]) Update(subject *Subject, input Input) {
	logic := m.Definition.Logic

	if !logic.ActiveFully {
		hasState := m.Active.Len() > 0 || len(m.ActiveQueue) > 0
		m.Transitions(m.Definition.Transitions, subject, input, hasState, nil)
	}

	m.ProcessQueue(subject, input)
	m.UpdateActive(subject, input)

	if logic.ProcessQueueImmediately && len(m.ActiveQueue) > 0 {
		m.ProcessQueue(subject, input)
	}
}

func (m *Machine[Subject, Input, StateData, TransitionData, Effect]) Apply(subject *Subject) {
	if m.Active.Empty() {
		return
	}

	logic := m.Definition.Logic

	m.Applicable = m.Applicable[:0]
	if m.Definition.Logic.AppliedDeep {
		var addApplicable func(state State[Subject, Input, StateData, TransitionData, Effect])

		addApplicable = func(state State[Subject, Input, StateData, TransitionData, Effect]) {
			if state.Sub != nil {
				for _, subState := range state.Sub.Active.Values() {
					addApplicable(subState)
				}
			} else {
				m.Applicable = append(m.Applicable, state)
			}
		}
	} else {
		for _, state := range m.Active.Values() {
			m.Applicable = append(m.Applicable, state)
		}
	}

	if logic.AppliedMax > 0 && logic.AppliedMax < m.Active.Len() {
		if logic.AppliedPriority != nil {
			sort.Slice(m.Applicable, func(i, j int) bool {
				a := m.Applicable[i]
				b := m.Applicable[j]
				return logic.AppliedPriority(a, b)
			})
		}
		m.Applicable = m.Applicable[:logic.AppliedMax]
	}

	m.Definition.Logic.Apply(subject, m.Applicable)
}
