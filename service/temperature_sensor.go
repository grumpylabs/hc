// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/grumpylabs/hcf/characteristic"
)

const TypeTemperatureSensor = "8A"

type TemperatureSensor struct {
	*Service

	CurrentTemperature *characteristic.CurrentTemperature

	StatusActive     *characteristic.StatusActive
	StatusFault      *characteristic.StatusFault
	StatusLowBattery *characteristic.StatusLowBattery
	StatusTampered   *characteristic.StatusTampered
	Name             *characteristic.Name
}

func NewTemperatureSensor() *TemperatureSensor {
	svc := TemperatureSensor{}
	svc.Service = New(TypeTemperatureSensor)

	svc.CurrentTemperature = characteristic.NewCurrentTemperature()
	svc.AddCharacteristic(svc.CurrentTemperature.Characteristic)

	svc.StatusActive = characteristic.NewStatusActive()
	svc.AddCharacteristic(svc.StatusActive.Characteristic)

	svc.StatusFault = characteristic.NewStatusFault()
	svc.AddCharacteristic(svc.StatusFault.Characteristic)

	svc.StatusLowBattery = characteristic.NewStatusLowBattery()
	svc.AddCharacteristic(svc.StatusLowBattery.Characteristic)

	svc.StatusTampered = characteristic.NewStatusTampered()
	svc.AddCharacteristic(svc.StatusTampered.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	return &svc
}
