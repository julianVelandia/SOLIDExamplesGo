package responsabilidad_unica

import (
	"fmt"
	"time"
)

const (
	NumberCarsToProduce = 5
	TimePaintProcess    = 4
	TimeTireProcess     = 1
	TimeInteriorProcess = 4
)

func ProductionLineWithoutPrinciple() {
	fmt.Println("Se inicia la línea de ensamblaje")
	for currentCar := 0; currentCar < NumberCarsToProduce; currentCar++ {
		fmt.Printf("Se inicia el ensamblaje del carro: %d\n", currentCar+1)
		fmt.Printf("Inicia proceso de pintura de la carrocería, Carro: %d\n", currentCar+1)
		time.Sleep(TimePaintProcess * time.Second)
		fmt.Printf("Finaliza proceso de pintura de la carrocería, Carro: %d\n", currentCar+1)
		fmt.Printf("Inicia proceso de ensamble de los neumáticos, Carro: %d\n", currentCar+1)
		for currentTire := 0; currentTire < 4; currentTire++ {
			fmt.Printf("\t Ensamble del neumático %d, del Carro: %d\n", currentTire+1, currentCar+1)
			time.Sleep(TimeTireProcess * time.Second)
		}
		fmt.Printf("Finaliza proceso de ensamble de los neumáticos, Carro: %d\n", currentCar+1)

		fmt.Printf("Inicia proceso de ensamble del interior, Carro: %d\n", currentCar+1)
		time.Sleep(TimeInteriorProcess * time.Second)
		fmt.Printf("Finaliza proceso de ensamble del interior, Carro: %d\n", currentCar+1)
		fmt.Printf("Se finaliza el ensamblaje del carro: %d\n", currentCar+1)
	}
	fmt.Println("Se finaliza la línea de ensamblaje")
}

func main() {
	ProductionLineWithoutPrinciple()
}
