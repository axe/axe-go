package input

import (
	"fmt"
	"sort"

	"github.com/axe/axe-go/pkg/ds"
	"github.com/axe/axe-go/pkg/react"
	"github.com/axe/axe-go/pkg/util"
)

type Action struct {
	Name             string
	Enabled          react.Value[bool]
	Trigger          Trigger
	Data             Data
	Triggered        bool
	OverridePriority int
}

func NewAction(name string, trigger Trigger) *Action {
	return &Action{
		Name:      name,
		Trigger:   trigger,
		Enabled:   react.Val(true),
		Data:      Data{},
		Triggered: false,
	}
}

func (action *Action) String() string {
	return fmt.Sprintf("%s: %.1f", action.Name, action.Data.Value)
}

func (action *Action) Init(inputs InputSystem) {
	if action.Trigger != nil {
		action.Trigger.Init(inputs)
	}
}

func (action *Action) Update(inputs InputSystem) {
	if action.Enabled.Get() && action.Trigger != nil {
		data, triggered := action.Trigger.Update(inputs)

		if triggered && len(data.Inputs) > 0 {
			for _, in := range data.Inputs {
				if in.Action != nil {
					triggered = false
					break
				}
			}
			if triggered {
				for _, in := range data.Inputs {
					in.Action = action
				}
			}
		}

		action.Data = data
		action.Triggered = triggered
	}
}

func (action Action) Priority() int {
	if action.OverridePriority > 0 || action.Trigger == nil {
		return action.OverridePriority
	}
	return action.Trigger.InputCount()
}

type Actions []*Action

func (actions Actions) Less(i, j int) bool {
	return actions[i].Priority() > actions[j].Priority()
}

func (actions Actions) Swap(i, j int) {
	t := actions[i]
	actions[i] = actions[j]
	actions[j] = t
}

func (actions Actions) Len() int {
	return len(actions)
}

func (actions Actions) Sort() {
	sort.Sort(actions)
}

type ActionSet struct {
	Name      string
	Enabled   react.Value[bool]
	Actions   Actions
	Triggered ds.CircularQueue[*Action]
}

func NewActionSet(name string) *ActionSet {
	return &ActionSet{
		Name:      name,
		Enabled:   react.Val(true),
		Actions:   make(Actions, 0, 64),
		Triggered: ds.NewCircularQueue[*Action](64),
	}
}

func CreateActionSet(name string, actions map[string]Trigger) *ActionSet {
	set := NewActionSet(name)
	for name, trigger := range actions {
		set.Add(NewAction(name, trigger))
	}
	return set
}

func (set *ActionSet) Init(inputs InputSystem) {
	for _, action := range set.Actions {
		action.Init(inputs)
	}

	set.Actions.Sort()
}

func (set *ActionSet) Update(inputs InputSystem) {
	set.Triggered.Clear()

	if !set.Enabled.Get() {
		return
	}
	if set.Actions != nil {
		for i := range set.Actions {
			action := set.Actions[i]
			action.Update(inputs)

			if action.Triggered {
				set.Triggered.Push(action)
			}
		}
	}
}

func (set *ActionSet) Add(action *Action) {
	set.Actions = append(set.Actions, action)
}

func (set *ActionSet) Iterator() ds.Iterator[*Action] {
	return set.Triggered.Iterator()
}

type ActionHandler func(action *Action)

type ActionSets struct {
	Sets    map[string]*ActionSet
	Handler ActionHandler
}

func NewActionSets() ActionSets {
	return ActionSets{
		Sets:    make(map[string]*ActionSet),
		Handler: nil,
	}
}

type ActionSetsInput = map[string]map[string]Trigger

func CreateActionSets(actionSets ActionSetsInput) ActionSets {
	sets := NewActionSets()
	for name, actions := range actionSets {
		sets.Sets[name] = CreateActionSet(name, actions)
	}
	return sets
}

func (sets *ActionSets) Init(inputs InputSystem) {
	if sets.Sets == nil {
		return
	}
	for _, set := range sets.Sets {
		set.Init(inputs)
	}
}

func (sets *ActionSets) Update(inputs InputSystem) {
	if sets.Sets == nil {
		return
	}
	for _, input := range inputs.Inputs() {
		input.Action = nil
	}
	for _, set := range sets.Sets {
		set.Update(inputs)
		if sets.Handler != nil {
			triggeredIterator := set.Triggered.Iterator()
			for triggeredIterator.HasNext() {
				triggered := triggeredIterator.Next()
				sets.Handler(*triggered)
			}
		}
	}
}

func (sets *ActionSets) Add(set *ActionSet) {
	sets.Sets[set.Name] = set
}

func (sets *ActionSets) Iterable() ds.Iterable[*Action] {
	return ds.NewMultiIterable(
		util.SliceMap(util.MapValues(sets.Sets), func(s *ActionSet) ds.Iterable[*Action] {
			return s
		}),
	)
}
