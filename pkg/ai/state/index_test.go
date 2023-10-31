package state_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/axe/axe-go/pkg/ai/state"
	"github.com/axe/axe-go/pkg/id"
	"golang.org/x/exp/constraints"
)

var (
	Jump          = state.Bool(false)
	OnGround      = state.Bool(true)
	GrabbingLedge = state.Bool(false)
	PullLedge     = state.Bool(false)
	ForwardSpeed  = state.Float(0)
	SideSpeed     = state.Float(0)
	FallingSpeed  = state.Float(0)
)

type Animator struct{ animating map[id.Identifier]float32 }

func (a Animator) String() string {
	out := strings.Builder{}
	for id, weight := range a.animating {
		if weight > 0 {
			if out.Len() > 0 {
				out.WriteString(", ")
			}
			out.WriteString(fmt.Sprintf("%s(weight=%.1f)", id.String(), weight))
		}
	}
	return out.String()
}

type Animation = id.Identifier
type TransitionOptions = string
type AnimationInput = state.UserData
type BlendWeight = float32

type State = state.State[Animator, AnimationInput, Animation, TransitionOptions, BlendWeight]
type StateDefinition = state.StateDefinition[Animator, AnimationInput, Animation, TransitionOptions, BlendWeight]
type Machine = state.Machine[Animator, AnimationInput, Animation, TransitionOptions, BlendWeight]
type MachineDefinition = state.MachineDefinition[Animator, AnimationInput, Animation, TransitionOptions, BlendWeight]
type Logic = state.Logic[Animator, AnimationInput, Animation, TransitionOptions, BlendWeight]
type Transition = state.Transition[AnimationInput, TransitionOptions]

var AnimationLogic = Logic{
	Start: func(subject *Animator, state State, trans *Transition, outro *State) bool {
		if outro != nil {
			delete(subject.animating, outro.Definition.Data)
		}
		subject.animating[state.Definition.Data] = state.Effect
		return true
	},
	Apply: func(subject *Animator, active []State) {
		for _, state := range active {
			subject.animating[state.Definition.Data] = state.Effect
		}
	},
	IsDone: func(subject *Animator, state State) bool {
		_, exists := subject.animating[state.Definition.Data]
		return !exists
	},
	InitializeSub: state.InitializeSubAfterStart,
}

func NewDefinition() MachineDefinition {
	return state.NewMachineDefinition(AnimationLogic)
}

func Abs[V constraints.Signed | constraints.Float](value V) V {
	if value < 0 {
		return -value
	}
	return value
}

func Min[V constraints.Signed | constraints.Float](a V, b V) V {
	if a < b {
		return a
	}
	return b
}

func Max[V constraints.Signed | constraints.Float](a V, b V) V {
	if a > b {
		return a
	}
	return b
}

// Effect functions
func IdleEffect(in AnimationInput) BlendWeight {
	return 1.0 - Max(Abs(ForwardSpeed.Get(in)), Abs(SideSpeed.Get(in)))
}
func ForwardEffect(in AnimationInput) BlendWeight {
	return Max(0, ForwardSpeed.Get(in))
}
func BackwardEffect(in AnimationInput) BlendWeight {
	return Max(0, -ForwardSpeed.Get(in))
}
func RightEffect(in AnimationInput) BlendWeight {
	return Max(0, SideSpeed.Get(in))
}
func LeftEffect(in AnimationInput) BlendWeight {
	return Max(0, -SideSpeed.Get(in))
}

// Transition conditions
func Jumped(in AnimationInput) bool       { return Jump.Get(in) }
func IsFalling(in AnimationInput) bool    { return FallingSpeed.Get(in) < 0 && !OnGround.Get(in) }
func Grounded(in AnimationInput) bool     { return OnGround.Get(in) }
func LedgeGrabbed(in AnimationInput) bool { return GrabbingLedge.Get(in) }
func LedgeLetGo(in AnimationInput) bool   { return !GrabbingLedge.Get(in) || OnGround.Get(in) }
func LedgePulled(in AnimationInput) bool  { return PullLedge.Get(in) }

func TestFakeAnimator(t *testing.T) {
	grounded := NewDefinition()
	grounded.Logic.ActiveFully = true
	grounded.AddState(StateDefinition{ID: id.Get("idle"), Data: id.Get("idle"), EffectGet: IdleEffect, EffectLive: true})
	grounded.AddState(StateDefinition{ID: id.Get("forward"), Data: id.Get("forward"), EffectGet: ForwardEffect, EffectLive: true})
	grounded.AddState(StateDefinition{ID: id.Get("backward"), Data: id.Get("backward"), EffectGet: BackwardEffect, EffectLive: true})
	grounded.AddState(StateDefinition{ID: id.Get("left"), Data: id.Get("strafeLeft"), EffectGet: LeftEffect, EffectLive: true})
	grounded.AddState(StateDefinition{ID: id.Get("right"), Data: id.Get("strageRight"), EffectGet: RightEffect, EffectLive: true})

	ledge := NewDefinition()
	ledge.Logic.ActiveFully = true
	ledge.AddState(StateDefinition{ID: id.Get("idle"), Data: id.Get("ledgeIdle"), EffectGet: IdleEffect, EffectLive: true})
	ledge.AddState(StateDefinition{ID: id.Get("forward"), Data: id.Get("ledgeUp"), EffectGet: ForwardEffect, EffectLive: true})
	ledge.AddState(StateDefinition{ID: id.Get("backward"), Data: id.Get("ledgeDown"), EffectGet: BackwardEffect, EffectLive: true})
	ledge.AddState(StateDefinition{ID: id.Get("left"), Data: id.Get("ledgeLeft"), EffectGet: LeftEffect, EffectLive: true})
	ledge.AddState(StateDefinition{ID: id.Get("right"), Data: id.Get("ledgeRight"), EffectGet: RightEffect, EffectLive: true})

	def := NewDefinition()
	def.Logic.ProcessQueueImmediately = true
	def.AddState(StateDefinition{ID: id.Get("grounded"), Sub: &grounded})
	def.AddState(StateDefinition{ID: id.Get("ledge"), Sub: &ledge})
	def.AddState(StateDefinition{ID: id.Get("ledgeGrab"), Data: id.Get("ledgeGrab"), EffectConstant: 1.0})
	def.AddState(StateDefinition{ID: id.Get("ledgeDrop"), Data: id.Get("ledgeDrop"), EffectConstant: 1.0})
	def.AddState(StateDefinition{ID: id.Get("ledgePullUp"), Data: id.Get("ledgePullUp"), EffectConstant: 1.0})
	def.AddState(StateDefinition{ID: id.Get("jumping"), Data: id.Get("jumping"), EffectConstant: 1.0})
	def.AddState(StateDefinition{ID: id.Get("falling"), Data: id.Get("falling"), EffectConstant: 1.0})
	def.AddState(StateDefinition{ID: id.Get("landing"), Data: id.Get("landing"), EffectConstant: 1.0})
	def.AddTransition(Transition{End: id.Get("grounded"), Condition: Grounded})
	def.AddTransition(Transition{End: id.Get("falling"), Condition: IsFalling})
	def.AddTransition(Transition{Start: id.Get("grounded"), End: id.Get("jumping"), Condition: Jumped, Live: true, Data: "player jumped"})
	def.AddTransition(Transition{Start: id.Get("grounded"), End: id.Get("falling"), Condition: IsFalling, Live: true, Data: "player walked off edge"})
	def.AddTransition(Transition{Start: id.Get("jumping"), End: id.Get("falling"), Condition: IsFalling, Live: true, Data: "player reached peak of jump, starting to fall"})
	def.AddTransition(Transition{Start: id.Get("falling"), End: id.Get("landing"), Condition: Grounded, Live: true, Data: "player hit ground, land"})
	def.AddTransition(Transition{Start: id.Get("landing"), End: id.Get("grounded"), Data: "land to grounded right away"})
	def.AddTransition(Transition{Start: id.Get("grounded"), End: id.Get("ledgeGrab"), Condition: LedgeGrabbed, Live: true, Data: "player walked up to ledge"})
	def.AddTransition(Transition{Start: id.Get("ledgeGrab"), End: id.Get("ledge"), Data: "player now climbing"})
	def.AddTransition(Transition{Start: id.Get("ledge"), End: id.Get("ledgePullUp"), Condition: LedgePulled, Live: true, Data: "player reached top of edge and wants to go up"})
	def.AddTransition(Transition{Start: id.Get("ledgePullUp"), End: id.Get("grounded"), Data: "finished getting to top of ledge, now grounded"})
	def.AddTransition(Transition{Start: id.Get("ledge"), End: id.Get("landing"), Condition: LedgeLetGo, Live: true, Data: "player was climbing but let go"})

	animator := &Animator{
		animating: make(map[id.Identifier]float32),
	}
	input := state.NewUserData(Jump, OnGround, GrabbingLedge, PullLedge, ForwardSpeed, SideSpeed, FallingSpeed)

	machine := state.NewMachine(&def)
	machine.Init(animator, input)

	// [ 0- 4] idle
	// [ 5- 9] walk half speed
	// [10-19] run full speed
	// [20-20] jump
	// [21-29] go up and fall down
	// [30-30] land
	// [31-31] continue running
	// [32-32] grab ledge
	// [33-33] climb up and left
	// [34-34] hang still on edge
	// [35-35] pull yourself up edge
	// [36-xx] you pulled yourself up and are standing idly
	for i := 0; i < 40; i++ {

		if i == 5 { // walk
			ForwardSpeed.Set(input, 0.5)
		} else if i == 10 { // run
			ForwardSpeed.Set(input, 1.0)
		} else if i == 20 { // jump
			Jump.Set(input, true)
			FallingSpeed.Set(input, 1.0)
			OnGround.Set(input, false)
		} else if i > 20 && i < 30 { // up & down
			Jump.Set(input, false)
			FallingSpeed.Set(input, FallingSpeed.Get(input)-0.2)
		} else if i == 30 { // land
			OnGround.Set(input, true)
			FallingSpeed.Set(input, 0.0)
		} else if i == 32 { // climb up and left
			GrabbingLedge.Set(input, true)
			OnGround.Set(input, false)
			SideSpeed.Set(input, -1.0)
			ForwardSpeed.Set(input, 1.0)
		} else if i == 34 { // pause climbing
			SideSpeed.Set(input, 0.0)
			ForwardSpeed.Set(input, 0.0)
		} else if i == 35 { // pull self up
			PullLedge.Set(input, true)
		} else if i == 36 { // standing on ledge that I climbed up
			PullLedge.Set(input, false)
			OnGround.Set(input, true)
			GrabbingLedge.Set(input, false)
		}

		// LogLevel = i >= 33 && i <= 36 ? debug : none;

		machine.Update(animator, input)
		machine.Apply(animator)

		fmt.Printf("Frame[%d]: %v\n", i, animator)
	}
}
