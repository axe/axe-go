package axe

import (
	"io"
	"time"
)

type AudioSystem interface { // & GameSystem
	GameSystem

	Listeners() []AudioListener
	Instances() []AudioInstance
	Settings() map[string]AudioSettings
	Sources() []AudioSource
}

type AudioAttenuation func(distance float32, volume float32) float32

// this component can emit audio, and the audio is tied to it. if the component is destroyed the audio is as well.
type AudioSource struct {
	Name      string
	Flags     uint
	Position  SpaceCoord
	Instances []AudioInstance
}

type AudioListener struct {
	Name          string
	Flags         uint
	Position      SpaceCoord
	Direction     SpaceCoord
	Ears          SpaceCoord // Camera position?
	EarsDirection SpaceCoord
	VolumeScale   float32
}

type AudioAmbienceOption struct {
	Data       *AudioData
	Chance     int
	PlayVolume NumericRange[float32]
	PlayLength NumericRange[float32]
}

type AudioAmbience struct {
	PlayChance           int
	Options              []AudioAmbienceOption
	Volume               float32
	MaxPlaysBeforeRepeat int
	TransitionTime       NumericRange[float32]
}

type AudioInstanceSettings struct {
	DirectionSpread     float32 // optional for direction sounds, how big the noise cone is
	StartOnRequest      bool    // if we set the start time when sound playing is requested, or when it actually starts. this plays into if we seek to the position the audio should've been playing by the time it was ready
	Times               int     // -1=Infinite loop
	Attenuate           AudioAttenuation
	IsVirtual           bool
	OnVirtual           AudioVirtualOp
	VirtualOutroTime    time.Duration // how long we fade out when going virtual
	VirtualMinRemaining time.Duration // if we are going virtual and not stopping but have <= this left then we just stop
	Importance          float32       // scales the listener volume by this amount to calculate a priority
}

type AudioInstance struct {
	Name           string
	Flags          uint
	Data           *AudioData
	Position       SpaceCoord // optional for spatial sounds
	Direction      SpaceCoord // optional for directional sounds
	PositionVolume float32    // volume at position
	Speed          float32    // playback rate modifier
	StartTime      time.Time
	PauseTime      time.Time
	Settings       *AudioInstanceSettings // custom or shared
	// relative to current listener
	Distance       float32
	ListenerVolume float32
	LeftVolume     float32
	RightVolume    float32
}

type AudioVirtualOp string

const (
	AudioVirtualOpStop    AudioVirtualOp = "stop"    // when going virtual, stop and ignore the audio
	AudioVirtualOpRestart AudioVirtualOp = "restart" // when coming back from virtual, restart from start
	AudioVirtualOpPause   AudioVirtualOp = "pause"   // when going virtual remember place in audio, and when coming back from virtual start at that place
	AudioVirtualOpPlay    AudioVirtualOp = "play"    // when going virtual remember place in audio and when it comes out of virtual play as if we could hear it the whole time
)

type AudioSettings struct {
	AudioUnloadTime   time.Duration // how long to wait before we unload unused audio
	MaxAudioInstances int           // user specified, otherwise determined by available APIs or trial & error
}

type AudioData struct {
	Length    int
	Streaming bool
	BitRate   int
	Reader    io.ReadSeekCloser
}

type AudioRequest struct {
	MaxWaitTime time.Duration
	Instance    AudioInstance
}
