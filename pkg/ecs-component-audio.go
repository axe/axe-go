package axe

type AudioEmitter struct {
	Source   AudioSource
	Pending  []AudioRequest
	Ambience *AudioAmbience
}

func (audio *AudioEmitter) Play(req AudioRequest) {
	audio.Pending = append(audio.Pending, req)
}

var AUDIO = DefineComponent("audio", AudioEmitter{})
