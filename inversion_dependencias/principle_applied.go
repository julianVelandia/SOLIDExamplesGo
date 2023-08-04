package inversion_dependencias

import "fmt"

type Swim interface {
	Swim() string
}

type SwimmingAnimal struct{}

func (sa *SwimmingAnimal) Swim() string {
	return "El animal está nadando"
}

type RealDuck struct {
	moveThroughTheWater Swim
}

func NewRealDuck(moveThroughTheWater Swim) *RealDuck {
	return &RealDuck{moveThroughTheWater: moveThroughTheWater}
}

func Wire() {
	swimmingAnimal := &SwimmingAnimal{}
	duck := NewRealDuck(swimmingAnimal)

	fmt.Println(duck.moveThroughTheWater.Swim())
}
