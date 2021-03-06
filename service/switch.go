// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/grumpylabs/hcf/characteristic"
)

const TypeSwitch = "49"

type Switch struct {
	*Service

	On *characteristic.On

	Name *characteristic.Name
}

func NewSwitch() *Switch {
	svc := Switch{}
	svc.Service = New(TypeSwitch)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	return &svc
}
