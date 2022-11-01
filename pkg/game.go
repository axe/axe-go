package axe

import "time"

type GameSystem interface {
	Init(game *Game) error
	Update(game *Game)
	Destroy()
}

type GameState struct {
	StartTime   time.Time
	UpdateTimer Timer
	DrawTimer   Timer
}

func (state *GameState) FromSettings(settings GameSettings) {
	if settings.FixedDrawFrequency > 0 {
		state.DrawTimer.Frequency = settings.FixedDrawFrequency
	}
	if settings.FixedUpdateFrequency > 0 {
		state.UpdateTimer.Frequency = settings.FixedUpdateFrequency
	}
}

type GameSettings struct {
	Name                 string
	EnableDebug          bool
	FixedUpdateFrequency time.Duration
	FixedDrawFrequency   time.Duration
	FirstStage           string
}

type Game struct {
	Debug    DebugSystem
	Assets   AssetSystem
	Windows  WindowSystem
	Graphics GraphicsSystem
	Input    InputSystem
	Actions  InputActionSets
	Audio    AudioSystem
	Events   EventSystem
	Stages   StageManager
	State    GameState
	Settings GameSettings
	Running  bool
}

func NewGame(settings GameSettings) *Game {
	state := GameState{}
	state.FromSettings(settings)

	return &Game{
		Settings: settings,
		State:    state,
		Assets:   NewAssetSystem(),
		Actions:  NewInputActionSets(),
		Stages:   NewStageManager(),
		Audio:    &NoAudioSystem{},
		Windows:  &NoWindowSystem{},
		Graphics: &NoGraphicsSystem{},
		Input:    &NoInputSystem{},
	}
}

func (game *Game) Run() error {
	err := game.Init()
	if err != nil {
		return err
	}
	defer game.Destroy()

	for game.Running {
		err = game.Tick()
		if err != nil {
			return err
		}
	}

	return nil
}

func (game *Game) Init() error {
	err := game.Assets.Init(game)
	if err != nil {
		return err
	}
	err = game.Debug.Init(game)
	if err != nil {
		return err
	}
	err = game.Windows.Init(game)
	if err != nil {
		return err
	}
	err = game.Graphics.Init(game)
	if err != nil {
		return err
	}
	err = game.Input.Init(game)
	if err != nil {
		return err
	}
	game.Actions.Init(game.Input)

	err = game.Audio.Init(game)
	if err != nil {
		return err
	}
	err = game.Events.Init(game)
	if err != nil {
		return err
	}
	err = game.Stages.Init(game)
	if err != nil {
		return err
	}

	game.Running = true
	game.State.StartTime = time.Now()
	game.State.UpdateTimer.Reset()
	game.State.DrawTimer.Reset()
	return nil
}

func (game *Game) Destroy() {
	game.Stages.Destroy()
	game.Events.Destroy()
	game.Audio.Destroy()
	game.Assets.Destroy()
	game.Input.Destroy()
	game.Graphics.Destroy()
	game.Windows.Destroy()
	game.Debug.Destroy()
}

func (game *Game) Tick() error {
	doUpdate := game.State.UpdateTimer.Tick()

	game.Windows.Update(game)
	game.Input.Update(game)
	game.Actions.Update(game.Input)
	game.Assets.Update(game)
	if doUpdate {
		game.Stages.Update(game)
	}
	game.Audio.Update(game)
	game.Debug.Update(game)

	doDraw := game.State.DrawTimer.Tick()
	if doDraw {
		game.Graphics.Update(game)
	}

	sleepUpdate := game.State.UpdateTimer.NextTick()
	sleepDraw := game.State.DrawTimer.NextTick()
	if sleepUpdate > 0 && sleepDraw > 0 {
		sleep := sleepUpdate
		if sleepDraw < sleep {
			sleep = sleepDraw
		}
		if sleep > 0 {
			time.Sleep(sleep)
		}
	}

	return nil
}

type NoAudioSystem struct{}

var _ AudioSystem = &NoAudioSystem{}

func (audio *NoAudioSystem) Init(game *Game) error { return nil }
func (audio *NoAudioSystem) Update(game *Game)     {}
func (audio *NoAudioSystem) Destroy()              {}

type NoWindowSystem struct{}

var _ WindowSystem = &NoWindowSystem{}

func (windows *NoWindowSystem) MainWindow() Window                     { return nil }
func (windows *NoWindowSystem) GetScreens() []Screen                   { return nil }
func (windows *NoWindowSystem) Events() *Listeners[WindowSystemEvents] { return nil }
func (windows *NoWindowSystem) Init(game *Game) error                  { return nil }
func (windows *NoWindowSystem) Update(game *Game)                      {}
func (windows *NoWindowSystem) Destroy()                               {}

type NoGraphicsSystem struct{}

var _ GraphicsSystem = &NoGraphicsSystem{}

func (gr *NoGraphicsSystem) Init(game *Game) error { return nil }
func (gr *NoGraphicsSystem) Update(game *Game)     {}
func (gr *NoGraphicsSystem) Destroy()              {}

type NoInputSystem struct{}

var _ InputSystem = &NoInputSystem{}

func (in *NoInputSystem) Devices() []*InputDevice               { return nil }
func (in *NoInputSystem) Inputs() []*Input                      { return nil }
func (in *NoInputSystem) InputTime() time.Time                  { return time.Time{} }
func (in *NoInputSystem) Get(name string) *Input                { return nil }
func (in *NoInputSystem) Points() []*InputPoint                 { return nil }
func (in *NoInputSystem) Events() *Listeners[InputSystemEvents] { return nil }
func (in *NoInputSystem) Init(game *Game) error                 { return nil }
func (in *NoInputSystem) Update(game *Game)                     {}
func (in *NoInputSystem) Destroy()                              {}
