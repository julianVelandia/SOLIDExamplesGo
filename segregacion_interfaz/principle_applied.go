package segregacion_interfaz

import (
	"fmt"
)

type Animal interface {
	LooksLikeAnAnimal() bool
}

type Duck interface {
	Animal
	Quack() string
}

type SwimmingAnimals interface {
	Animal
	Swim() string
}

type RealDuck struct{}

func (rd RealDuck) LooksLikeAnAnimal() bool {
	return true
}

func (rd RealDuck) Swim() string {
	return "El pato está nadando"
}

func (rd RealDuck) Quack() string {
	return "El pato hace Quack!"
}

func (rd RealDuck) QuackAndSwim() string {
	return rd.Quack() + rd.Swim()
}

type ToyDuck struct {
	HaveBatteries bool
}

func (td ToyDuck) LooksLikeAnAnimal() bool {
	return true
}

func (td ToyDuck) Quack() string {
	return "El pato de juguete hace Quack!"
}

func MakeDucksWithPrincipleApplied() {
	realDuck := RealDuck{}

	fmt.Println("El pato real parece animal? : ", realDuck.LooksLikeAnAnimal())
	fmt.Println("El pato real está nadando? : ", realDuck.Swim())
	fmt.Println("El pato real hace Quack? : ", realDuck.Quack())
	fmt.Println("El pato real hace Quack y nada? : ", realDuck.QuackAndSwim())

	toyDuck := ToyDuck{true}

	fmt.Println("El pato de juguete parece animal? : ", toyDuck.LooksLikeAnAnimal())
	fmt.Println("El pato de juguete hace Quack? : ", toyDuck.Quack())
	fmt.Println("El pato de juguete tiene baterias? : ", toyDuck.HaveBatteries)

}
