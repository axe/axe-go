package axe

import "github.com/axe/axe-go/pkg/ecs"

type AudioEmitter struct {
	Source   AudioSource
	Pending  []AudioRequest
	Ambience *AudioAmbience
}

func (audio *AudioEmitter) Play(req AudioRequest) {
	audio.Pending = append(audio.Pending, req)
}

var AUDIO = ecs.DefineComponent("audio", AudioEmitter{})
