package ai

// when passing & storing Subject, use ptr
// when storing Input, use ptr
// when storing Data use ptr

type StateBlend[Data any] struct {
	Data   *Data
	Amount float32
}

type StateTransitions[Input any] struct {
	Start     int
	End       int
	Condition func(input *Input) bool
}

type State[Data any, Input any] func(input *Input) StateBlend[Data]

type StateMachineState[Data any, Input any] struct {
	ID          int
	State       State[Data, Input]
	Transitions []*StateTransitions[Input]
}

func StateConstant[Data any](data *Data, blend float32) State[Data, any] {
	stateBlend := StateBlend[Data]{Data: data, Amount: blend}
	return func(input *any) StateBlend[Data] {
		return stateBlend
	}
}

func StateSub[Subject any, Data any, Input any](machine StateMachine[Subject, Data, Input]) State[Data, Input] {
	return func(input *Input) StateBlend[Data] {
		return StateBlend[Data]{}
	}
}

type StateMachineLogic[Subject any, Data any, Input any] struct {
	IsDone func(subject *Subject, state StateBlend[Data]) bool
	Start  func(subject *Subject, state StateBlend[Data])
	Apply  func(subject *Subject, blends []StateBlend[Data]) bool
}

type StateMachine[Subject any, Data any, Input any] struct {
	Logic       StateMachineLogic[Subject, Data, Input]
	States      map[int]StateMachineState[Data, Input]
	Transitions []*StateTransitions[Input]
}

func NewStateMachine[Subject any, Data any, Input any](logic StateMachineLogic[Subject, Data, Input]) *StateMachine[Subject, Data, Input] {
	return &StateMachine[Subject, Data, Input]{
		Logic:       logic,
		States:      make(map[int]StateMachineState[Data, Input]),
		Transitions: make([]*StateTransitions[Input], 0),
	}
}

func (sm *StateMachine[Subject, Data, Input]) AddState(id int, state State[Data, Input]) {
	sm.States[id] = StateMachineState[Data, Input]{
		ID:          id,
		State:       state,
		Transitions: make([]*StateTransitions[Input], 0),
	}
}

func (sm *StateMachine[Subject, Data, Input]) AddTransition(transition StateTransitions[Input]) {
	if state, ok := sm.States[transition.Start]; ok {
		state.Transitions = append(state.Transitions, &transition)
	}
	sm.Transitions = append(sm.Transitions, &transition)
}

func (sm *StateMachine[Subject, Data, Input]) For(subject *Subject, input *Input, initialStates []int) *StateMachineInstance[Subject, Data, Input] {
	inst := &StateMachineInstance[Subject, Data, Input]{
		Machine: sm,
		Subject: subject,
		Input:   input,
		States:  make(map[int]StateInstance[Data, Input]),
	}

	for _, id := range initialStates {
		if state, ok := sm.States[id]; ok {
			inst.States[id] = StateInstance[Data, Input]{
				State: &state,
			}
		}
	}

	return inst
}

type StateInstance[Data any, Input any] struct {
	StateBlend StateBlend[Data]
	State      *StateMachineState[Data, Input]
}

type StateMachineInstance[Subject any, Data any, Input any] struct {
	Machine *StateMachine[Subject, Data, Input]
	Subject *Subject
	Input   *Input
	States  map[int]StateInstance[Data, Input]
}

func (inst StateMachineInstance[Subject, Data, Input]) Update() {
	for _, state := range inst.States {
		for _, transition := range state.State.Transitions {
			if state.State.ID == transition.Start {
				proceed := false
				if transition.Condition != nil {
					proceed = transition.Condition(inst.Input)
				} else {
					proceed = inst.Machine.Logic.IsDone(inst.Subject, state.StateBlend)
				}
				if proceed {
					if next, ok := inst.Machine.States[transition.End]; ok {
						state.StateBlend = next.State(inst.Input)
						state.State = &next
						inst.Machine.Logic.Start(inst.Subject, state.StateBlend)
					}
				}
			}
		}
	}
	inst.Apply()
}

func (inst StateMachineInstance[Subject, Data, Input]) Apply() {
	inst.Machine.Logic.Apply(inst.Subject, inst.GetActive())
}

func (inst StateMachineInstance[Subject, Data, Input]) GetActive() []StateBlend[Data] {
	active := make([]StateBlend[Data], 0)
	for _, state := range inst.States {
		state.StateBlend = state.State.State(inst.Input)
		if state.StateBlend.Amount != 0 {
			active = append(active, state.StateBlend)
		}
	}
	return active
}
