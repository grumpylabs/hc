// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/grumpylabs/hcf/characteristic"
)

const TypeSpeaker = "113"

type Speaker struct {
	*Service

	Mute *characteristic.Mute

	Name   *characteristic.Name
	Volume *characteristic.Volume
}

func NewSpeaker() *Speaker {
	svc := Speaker{}
	svc.Service = New(TypeSpeaker)

	svc.Mute = characteristic.NewMute()
	svc.AddCharacteristic(svc.Mute.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.Volume = characteristic.NewVolume()
	svc.AddCharacteristic(svc.Volume.Characteristic)

	return &svc
}
