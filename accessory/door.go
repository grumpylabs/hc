package accessory

import (
	"github.com/grumpylabs/hcf/service"
)

// Door represents the Door type accessory
type Door struct {
	*Accessory
	Door *service.Door
}

// NewDoor returns a door accessory containing one door service.
func NewDoor(info Info) *Door {
	acc := Door{}
	acc.Accessory = New(info, TypeDoor)
	acc.Door = service.NewDoor()

	acc.AddService(acc.Door.Service)

	return &acc
}
