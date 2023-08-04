package sustitucion_liskov

import (
	"fmt"
)

type Animal struct {
	LooksLikeAnAnimal bool
}

type Duck struct {
	Animal
}

func (d Duck) Quack() string {
	return "El pato hace Quack!"
}

// Este método es de Duck, no de Animal
func (d Duck) Swim() string {
	return "El pato está nadando"
}

func (d Duck) QuackAndSwim() (string, error) {
	return d.Quack() + d.Swim(), nil
}

type ToyDuck struct {
	Animal
	HaveBatteries bool
}

func (td ToyDuck) Quack() string {
	return "El pato hace Quack!"
}

func MakeDucksWithPrincipleApplied() {
	realDuck := Duck{Animal{true}}

	fmt.Println("El pato real parece animal? : ", realDuck.LooksLikeAnAnimal)
	fmt.Println("El pato real está nadando? : ", realDuck.Swim())
	fmt.Println("El pato real hace Quack? : ", realDuck.Quack())
	realDuckCanQuackAndSwim, err := realDuck.QuackAndSwim()
	if err != nil {
		fmt.Println("El pato real No puede hacer Quack y nadar : ", realDuckCanQuackAndSwim)
	} else {
		fmt.Println("El pato real hace Quack y nada? : ", realDuckCanQuackAndSwim)
	}

	toyDuck := ToyDuck{Animal{true}, true}

	fmt.Println("El pato de juguete parece animal? : ", toyDuck.LooksLikeAnAnimal)
	fmt.Println("El pato de juguete hace Quack? : ", toyDuck.Quack())
	fmt.Println("El pato de juguete tiene baterias? : ", toyDuck.HaveBatteries)

}
