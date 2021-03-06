// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/grumpylabs/hcf/characteristic"
)

const TypeFan = "40"

type Fan struct {
	*Service

	On *characteristic.On

	RotationDirection *characteristic.RotationDirection
	RotationSpeed     *characteristic.RotationSpeed
	Name              *characteristic.Name
}

func NewFan() *Fan {
	svc := Fan{}
	svc.Service = New(TypeFan)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	svc.RotationDirection = characteristic.NewRotationDirection()
	svc.AddCharacteristic(svc.RotationDirection.Characteristic)

	svc.RotationSpeed = characteristic.NewRotationSpeed()
	svc.AddCharacteristic(svc.RotationSpeed.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	return &svc
}
