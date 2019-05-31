// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/grumpylabs/hcf/characteristic"
)

const TypeTemperatureSensor = "8A"

type TemperatureSensor struct {
	*Service

	CurrentTemperature *characteristic.CurrentTemperature
}

func NewTemperatureSensor() *TemperatureSensor {
	svc := TemperatureSensor{}
	svc.Service = New(TypeTemperatureSensor)

	svc.CurrentTemperature = characteristic.NewCurrentTemperature()
	svc.AddCharacteristic(svc.CurrentTemperature.Characteristic)

	return &svc
}
