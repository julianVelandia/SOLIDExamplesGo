package sustitucion_liskov_principle_not_applied

import (
	"errors"
	"fmt"
)

type Animal struct {
	LooksLikeAnAnimal bool
}

func (a Animal) Swim() string {
	return "El animal está nadando"
}

type Duck struct {
	Animal
}

func (d Duck) Quack() string {
	return "El pato hace Quack!"
}

func (d Duck) Swim() string {
	return "El pato está nadando"
}

func (d Duck) QuackAndSwim() (string, error) {
	return fmt.Sprintf(d.Quack(), d.Swim()), nil
}

type ToyDuck struct {
	Animal
}

func (td ToyDuck) Quack() string {
	return "El pato hace Quack!"
}

func (td ToyDuck) Swim() string {
	panic("Un pato de juguete no puede nadar")
}

func (td ToyDuck) QuackAndSwim() (string, error) {
	//Retornará un error
	errTxt := "un pato de juguete no puede nadar, es un comportamiento inesperado"
	return errTxt, errors.New(errTxt)
}

func MakeDucksWithOutPrincipleApplied() {
	realDuck := Duck{}

	fmt.Println("El pato real parece animal? : ", realDuck.LooksLikeAnAnimal)
	fmt.Println("El pato real está nadando? : ", realDuck.Swim())
	fmt.Println("El pato real hace Quack? : ", realDuck.Quack())
	realDuckCanQuackAndSwim, err := realDuck.QuackAndSwim()
	if err != nil {
		fmt.Println("El pato real No puede hacer Quack y nadar : ", realDuckCanQuackAndSwim)
	} else {
		fmt.Println("El pato real hace Quack y nada? : ", realDuckCanQuackAndSwim)
	}

	toyDuck := ToyDuck{}

	fmt.Println("El pato de juguete parece animal? : ", toyDuck.LooksLikeAnAnimal)
	// El método nadar en toyDuck me retorna Panic
	fmt.Println("El pato de juguete hace Quack? : ", toyDuck.Quack())

	toyDuckCanQuackAndSwim, err := toyDuck.QuackAndSwim()
	if err != nil {
		fmt.Println("El pato de juguete No puede hacer Quack y nadar : ", err.Error())
	} else {
		fmt.Println("El pato de juguete hace Quack y nada? : ", toyDuckCanQuackAndSwim)
	}
}
