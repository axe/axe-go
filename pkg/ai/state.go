package axe

import (
	"fmt"
	"math"
)

type StateInput struct {
	ID    int
	Name  string
	Value float32
}

type StateInputMap map[int]StateInput

type StateCondition func(input StateInputMap, dt float32) bool

type StateBlend func(input StateInputMap) float32

type StateIsDone[S any, T any, O any] func(*S, *T, *O) bool

type StateStart[S any, T any, O any] func(*S, *T, *O)

type StateApply[S any, T any, O any] func(*S, []*StateInstance[S, T, O]) bool

type StateTransition[S any, T any, O any] struct {
	Start     int
	End       int
	Condition StateCondition
	Options   O
}

type State[S any, T any, O any] struct {
	ID          int
	Name        string
	State       *T
	Blend       StateBlend
	Transitions map[int][]StateTransition[S, T, O]
	Sub         *StateMachine[S, T, O]
}

type StateMachine[S any, T any, O any] struct {
	States        map[int]State[S, T, O]
	DefaultInputs StateInputMap
	IsDone        StateIsDone[S, T, O]
	Start         StateStart[S, T, O]
	Apply         StateApply[S, T, O]
}

type StateMachineCreate[S any, T any, O any] struct {
	IsDone StateIsDone[S, T, O]
	Start  StateStart[S, T, O]
	Apply  StateApply[S, T, O]
}

type StateInstance[S any, T any, O any] struct {
	Blend   float32
	State   *State[S, T, O]
	Sub     *StateInstance[S, T, O]
	Options *O
}

type StateMachineInstance[S any, T any, O any] struct {
	Machine *StateMachine[S, T, O]
	Subject *S
	Inputs  StateInputMap
	States  []StateInstance[S, T, O]
}

func New[S any, T any, O any](create StateMachineCreate[S, T, O]) *StateMachine[S, T, O] {
	return &StateMachine[S, T, O]{
		States:        make(map[int]State[S, T, O]),
		DefaultInputs: make(StateInputMap),
		IsDone:        create.IsDone,
		Start:         create.Start,
		Apply:         create.Apply,
	}
}

func (s *StateMachine[S, T, O]) AddState(state State[S, T, O]) int {
	if state.ID == 0 {
		state.ID = len(s.States) + 1
	}
	state.Transitions = make(map[int][]StateTransition[S, T, O])
	s.States[state.ID] = state
	return state.ID
}

func (s *StateMachine[S, T, O]) AddInput(name string, defaultValue float32) int {
	i := StateInput{
		ID:    len(s.DefaultInputs) + 1,
		Name:  name,
		Value: defaultValue,
	}
	s.DefaultInputs[i.ID] = i
	return i.ID
}

func (s *StateMachine[S, T, O]) AddTransition(start int, end int, condition StateCondition, options O) {
	if _, exists := s.States[start]; !exists {
		panic(fmt.Sprintf("The state %v does not exist in the state machine.", start))
	}
	t := StateTransition[S, T, O]{
		Start:     start,
		End:       end,
		Condition: condition,
		Options:   options,
	}

	s.States[start].Transitions[start] = append(s.States[start].Transitions[start], t)
}

func (s *StateMachine[S, T, O]) NewInstance(subject *S) StateMachineInstance[S, T, O] {
	return NewInstance(s, subject, cloneMap(s.DefaultInputs))
}

func NewInstance[S any, T any, O any](machine *StateMachine[S, T, O], subject *S, inputs StateInputMap) StateMachineInstance[S, T, O] {
	return StateMachineInstance[S, T, O]{
		Machine: machine,
		Subject: subject,
		Inputs:  inputs,
		States:  make([]StateInstance[S, T, O], 0),
	}
}

func (s *StateMachineInstance[S, T, O]) Set(input int, value float32) {
	if i, ok := s.Inputs[input]; ok {
		i.Value = value
	}
}

func (s *StateMachineInstance[S, T, O]) Get(input int) float32 {
	return s.Inputs[input].Value
}

func (s *StateMachineInstance[S, T, O]) GetState(id int) *StateInstance[S, T, O] {
	for i := range s.States {
		state := &s.States[i]
		if state.State.ID == id {
			return state
		}
	}
	return nil
}

func (s *StateMachineInstance[S, T, O]) Update(dt float32) {
	for i := range s.States {
		state := &s.States[i]
		if transitions, ok := state.State.Transitions[state.State.ID]; ok {
			for _, transition := range transitions {
				proceed := false
				if transition.Condition != nil {
					proceed = transition.Condition(s.Inputs, dt)
				} else {
					proceed = s.Machine.IsDone(s.Subject, state.State.State, &transition.Options)
				}
				if proceed {
					if next, nextExists := s.Machine.States[transition.End]; nextExists {
						state.State = &next
						state.Options = &transition.Options
						if next.Sub != nil {
							// state.Sub = NewInstance[S, T, O](next.Sub, s.Subject, s.Inputs)
						}
						s.Machine.Start(s.Subject, state.State.State, &transition.Options)
					} else {
						panic(fmt.Sprintf("The next state %v does not exist in the state machine.", transition.End))
					}
				}
			}
		}
	}
	active := make([]*StateInstance[S, T, O], 0)
	for i := range s.States {
		state := &s.States[i]
		state.Blend = state.State.Blend(s.Inputs)
		if state.Blend > 0 {
			active = append(active, state)
		}
	}
	s.Machine.Apply(s.Subject, active)
}

func cloneMap[K comparable, V any](s map[K]V) map[K]V {
	d := make(map[K]V)
	for k, v := range s {
		d[k] = v
	}
	return d
}

func t() {
	type Animator struct{}
	type Animation struct{}
	type AnimationOptions struct{}

	sm := New(StateMachineCreate[Animator, Animation, AnimationOptions]{})
	var (
		// GROUNDED   = sm.AddInput("grounded", 1)
		// HIT_LEDGE  = sm.AddInput("hitLedge", 0)
		// LEDGE_OPEN = sm.AddInput("ledgeOpen", 0)
		// FALLING    = sm.AddInput("falling", 0)
		JUMP   = sm.AddInput("jump", 0)
		MOVE_X = sm.AddInput("moveX", 0)
		MOVE_Y = sm.AddInput("moveY", 0)
	)

	subStateOptions := StateMachineCreate[Animator, Animation, AnimationOptions]{
		IsDone: func(*Animator, *Animation, *AnimationOptions) bool {
			// return animator->GetState(animation).done;
			return false
		},
		Start: func(*Animator, *Animation, *AnimationOptions) {
			// animator->Transition(animation, options);
		},
		Apply: func(*Animator, []*StateInstance[Animator, Animation, AnimationOptions]) bool {
			return false
		},
	}

	grounded := New(subStateOptions)

	grounded.AddState(State[Animator, Animation, AnimationOptions]{
		Name:  "idle",
		State: nil, // Idle Animation
		Blend: func(input StateInputMap) float32 {
			return float32(math.Abs(float64(input[MOVE_X].Value * input[MOVE_Y].Value)))
		},
	})
	grounded.AddState(State[Animator, Animation, AnimationOptions]{
		Name:  "forward",
		State: nil, // Forward Animation
		Blend: func(input StateInputMap) float32 {
			return float32(math.Max(0, float64(input[MOVE_Y].Value)))
		},
	})

	ledge := New(subStateOptions)

	var STATE_GROUNDED = sm.AddState(State[Animator, Animation, AnimationOptions]{
		Name: "grounded",
		Sub:  grounded,
	})
	/*var STATE_LEDGE = */ sm.AddState(State[Animator, Animation, AnimationOptions]{
		Name: "ledge",
		Sub:  ledge,
	})
	/*var STATE_LEDGE_GRAB = */ sm.AddState(State[Animator, Animation, AnimationOptions]{
		Name:  "ledgeGrab",
		State: nil, // Ledge Grab animation
	})
	/*var STATE_LEDGE_PULL_UP = */ sm.AddState(State[Animator, Animation, AnimationOptions]{
		Name:  "ledgePullUp",
		State: nil, // Ledge Pull Up animation
	})
	var STATE_AIR_JUMP = sm.AddState(State[Animator, Animation, AnimationOptions]{
		Name:  "airJump",
		State: nil, // Air Jump animation
	})

	sm.AddTransition(STATE_GROUNDED, STATE_AIR_JUMP, func(input StateInputMap, dt float32) bool {
		return input[JUMP].Value > 0
	}, AnimationOptions{})
}
