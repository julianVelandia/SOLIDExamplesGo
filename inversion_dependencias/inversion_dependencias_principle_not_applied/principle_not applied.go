package inversion_dependencias_principle_not_applied

import (
	"fmt"
)

type SwimmingAnimal struct{}

func (sa *SwimmingAnimal) Swim() string {
	return "El animal está nadando"
}

type RealDuck struct {
	MoveThroughTheWater SwimmingAnimal
}

func BadWire() {
	swimmingAnimal := SwimmingAnimal{}

	duck := RealDuck{swimmingAnimal}

	fmt.Println(duck.MoveThroughTheWater)
}
