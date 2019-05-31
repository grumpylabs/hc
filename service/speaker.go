// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/grumpylabs/hcf/characteristic"
)

const TypeSpeaker = "113"

type Speaker struct {
	*Service

	Mute *characteristic.Mute
}

func NewSpeaker() *Speaker {
	svc := Speaker{}
	svc.Service = New(TypeSpeaker)

	svc.Mute = characteristic.NewMute()
	svc.AddCharacteristic(svc.Mute.Characteristic)

	return &svc
}
