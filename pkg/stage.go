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

type Stage struct {
	Name    string
	Assets  []AssetRef
	Windows []StageWindow
	Scenes2 []Scene2f
	Scenes3 []Scene3f
	Views2  []View2f
	Views3  []View3f
	Actions InputActionSets

	pendingAssets map[string]*Asset
	loadedAssets  map[string]*Asset
}

func (stage Stage) HasStartedLoading() bool {
	return stage.pendingAssets != nil
}

func (stage *Stage) Load(game *Game) {
	if stage.pendingAssets == nil {
		stage.pendingAssets = game.Assets.AddManyMap(stage.Assets)
		stage.loadedAssets = make(map[string]*Asset, len(stage.pendingAssets))
	}
	for key, asset := range stage.pendingAssets {
		// If loading hasn't started, kick it off
		if !asset.LoadStatus.Started {
			game.Debug.LogDebug("Loading asset: %s\n", asset.Ref.URI)
			go game.Debug.LogError(asset.Load()) // TODO handle
		}
		// If loading is not done, check the next asset
		if !asset.LoadStatus.Done {
			continue
		}
		game.Debug.LogDebug("Asset loaded: %s\n", asset.Ref.URI)
		// The asset is loaded, remove it from pending and add it to loaded.
		delete(stage.pendingAssets, key)
		stage.loadedAssets[key] = asset
		// The asset is loaded, if any need to follow add them
		if len(asset.Next) > 0 {
			for _, nextRef := range asset.Next {
				nextAsset := game.Assets.Add(nextRef)

				if stage.loadedAssets[nextAsset.Ref.URI] == nil {
					stage.pendingAssets[nextAsset.Ref.URI] = nextAsset
				}
			}
		}
	}
	for _, asset := range stage.loadedAssets {
		// If the asset has not been activated, activate it
		if !asset.ActivateStatus.Started {
			game.Debug.LogDebug("Activating asset: %s\n", asset.Ref.URI)
			game.Debug.LogError(asset.Activate()) // TODO handle
		}
	}
}

func (stage *Stage) IsLoaded() bool {
	if stage.pendingAssets != nil {
		return len(stage.pendingAssets) == 0
	}
	return false
}

func (stage *Stage) Unload(activeStage *Stage) {
	if stage.loadedAssets != nil {
		unusedAssets := make(map[string]*Asset)
		for i := range stage.loadedAssets {
			asset := stage.loadedAssets[i]
			unusedAssets[asset.Ref.URI] = asset
		}
		if activeStage != nil && activeStage.loadedAssets != nil {
			for _, asset := range activeStage.loadedAssets {
				delete(unusedAssets, asset.Ref.URI)
			}
		}
		for _, asset := range unusedAssets {
			asset.Unload() // TODO error?
		}
	}
	stage.loadedAssets = nil
	stage.pendingAssets = nil
}

func (stage *Stage) Start(game *Game) {
	stage.Actions.Init(game.Input)

	for i := range stage.Scenes2 {
		game.Debug.LogError(stage.Scenes2[i].Init(game))
	}
	for i := range stage.Scenes3 {
		game.Debug.LogError(stage.Scenes3[i].Init(game))
	}

	for i := range stage.Views2 {
		game.Debug.LogError(stage.Views2[i].Init(game))
	}
	for i := range stage.Views3 {
		game.Debug.LogError(stage.Views3[i].Init(game))
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
	for i := range stage.Scenes2 {
		stage.Scenes2[i].Update(game)
	}
	for i := range stage.Scenes3 {
		stage.Scenes3[i].Update(game)
	}

	// update camera
	for i := range stage.Views2 {
		stage.Views2[i].Update(game)
	}
	for i := range stage.Views3 {
		stage.Views3[i].Update(game)
	}
}
