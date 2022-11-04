package axe

import (
	"strings"

	"github.com/axe/axe-go/pkg/ui"
)

type StageWindow struct {
	Name       string
	Title      string
	Placement  ui.Placement
	Fullscreen bool
}

type StageManager struct {
	Stages  map[string]*Stage
	Current *Stage
	Next    *Stage

	events *Listeners[StageManagerEvents]
}

var _ GameSystem = &StageManager{}

func NewStageManager() StageManager {
	return StageManager{
		Stages:  make(map[string]*Stage),
		Current: nil,
		Next:    nil,
		events:  NewListeners[StageManagerEvents](),
	}
}

func (sm *StageManager) Add(stage *Stage) {
	sm.Stages[strings.ToLower(stage.Name)] = stage
}

func (sm *StageManager) Events() *Listeners[StageManagerEvents] {
	return sm.events
}

func (sm *StageManager) Init(game *Game) error {
	sm.Update(game)
	return nil
}

func (sm *StageManager) Update(game *Game) {
	sm.initStage(game)

	if sm.Next != nil {
		if !sm.Next.HasStartedLoading() {
			sm.triggerTransitionStart()
		}
		sm.Next.Load(game)
		if sm.Next.IsLoaded() {
			sm.Next.Start(game)
			sm.triggerTransitionEnd()
			if sm.Current != nil {
				sm.Current.Unload(sm.Next)
			}
			sm.Current = sm.Next
			sm.Next = nil
		}
	}

	if sm.Current != nil {
		sm.Current.Update(game)
	}
}

func (sm *StageManager) triggerTransitionStart() {
	sm.events.Trigger(func(listener StageManagerEvents) bool {
		handled := false
		if listener.StageExiting != nil {
			listener.StageExiting(sm.Current, sm.Next)
			handled = true
		}
		if listener.StageStarting != nil {
			listener.StageStarting(sm.Next)
			handled = true
		}
		return handled
	})
}

func (sm *StageManager) triggerTransitionEnd() {
	sm.events.Trigger(func(listener StageManagerEvents) bool {
		handled := false
		if listener.StageStarted != nil {
			listener.StageStarted(sm.Next)
			handled = true
		}
		if listener.StageExited != nil {
			listener.StageExited(sm.Current, sm.Next)
			handled = true
		}
		return handled
	})
}

func (sm *StageManager) Destroy() {
	if sm.Next != nil {
		sm.Next.Unload(nil)
		sm.Next = nil
	}
	if sm.Current != nil {
		sm.Current.Unload(nil)
		sm.Current = nil
	}
}

func (sm *StageManager) initStage(game *Game) {
	if sm.Current == nil && sm.Next == nil && game.Settings.FirstStage != "" {
		sm.Set(game.Settings.FirstStage)
	}
}

func (sm *StageManager) Get(name string) *Stage {
	return sm.Stages[strings.ToLower(name)]
}

func (sm *StageManager) Set(name string) bool {
	if sm.Current != nil && sm.Current.Name == name {
		return true
	}
	if sm.Next != nil && sm.Next.Name == name {
		return true
	}
	next := sm.Get(name)
	if next == nil {
		return false
	}
	if sm.Next != nil {
		sm.Next.Unload(next)
		sm.Next = nil
	}
	sm.Next = next
	return true
}

type StageManagerEvents struct {
	StageStarting func(current *Stage)
	StageStarted  func(current *Stage)
	StageExiting  func(current *Stage, next *Stage)
	StageExited   func(previous *Stage, current *Stage)
}

// type SceneType interface {
// 	~Scene2f | ~Scene3f
// }

type Stage struct {
	Name    string
	Assets  []AssetRef
	Windows []StageWindow
	Scenes2 []Scene2f
	Scenes3 []Scene3f
	Views2  []View2f
	Views3  []View3f
	Actions InputActionSets

	assets []*Asset
}

func (stage Stage) HasStartedLoading() bool {
	return stage.assets != nil
}

func (stage *Stage) Load(game *Game) {
	if stage.assets == nil {
		stage.assets = game.Assets.AddMany(stage.Assets)
	}
	for _, asset := range stage.assets {
		if !asset.LoadStatus.Started {
			go game.Debug.LogError(asset.Load()) // TODO handle
		}
		if !asset.LoadStatus.IsSuccess() {
			continue
		}
		if !asset.ActivateStatus.Started {
			game.Debug.LogError(asset.Activate()) // TODO handle
		}
	}
}

func (stage *Stage) IsLoaded() bool {
	if stage.assets != nil {
		for _, asset := range stage.assets {
			if !asset.ActivateStatus.IsSuccess() {
				return false
			}
		}
	}
	return true
}

func (stage *Stage) Unload(activeStage *Stage) {
	if stage.assets != nil {
		unusedAssets := make(map[string]*Asset)
		for i := range stage.assets {
			asset := stage.assets[i]
			unusedAssets[asset.Ref.URI] = asset
		}
		if activeStage != nil && activeStage.assets != nil {
			for _, asset := range activeStage.assets {
				delete(unusedAssets, asset.Ref.URI)
			}
		}
		for _, asset := range unusedAssets {
			asset.Unload() // TODO error?
		}
	}
	stage.assets = nil
}

func (stage *Stage) Start(game *Game) {
	stage.Actions.Init(game.Input)

	for _, scene := range stage.Scenes2 {
		game.Debug.LogError(scene.Init(game))
	}
	for _, scene := range stage.Scenes3 {
		game.Debug.LogError(scene.Init(game))
	}

	for _, view := range stage.Views2 {
		game.Debug.LogError(view.Init(game))
	}
	for _, view := range stage.Views3 {
		game.Debug.LogError(view.Init(game))
	}
}

func (stage *Stage) Update(game *Game) {
	// get input/actions
	stage.Actions.Update(game.Input)

	// handle input/actions
	// update movement
	// update space
	// update collisions
	// update space
	for _, scene := range stage.Scenes2 {
		scene.Update(game)
	}
	for _, scene := range stage.Scenes3 {
		scene.Update(game)
	}

	// update camera
	for _, view := range stage.Views2 {
		view.Update(game)
	}
	for _, view := range stage.Views3 {
		view.Update(game)
	}
}
