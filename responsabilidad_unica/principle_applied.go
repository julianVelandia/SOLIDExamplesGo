package responsabilidad_unica

import (
	"fmt"
	"time"
)

func ProductionLineWithPrinciple() {
	fmt.Println("Se inicia la línea de ensamblaje")
	for currentCar := 0; currentCar < NumberCarsToProduce; currentCar++ {
		formatProcessMessage(true, "ensamblaje del carro", currentCar)
		paintProcess(currentCar)
		tireProcess(currentCar)
		interiorProcess(currentCar)
		formatProcessMessage(false, "ensamblaje del carro", currentCar)
	}
	fmt.Println("Se finaliza la línea de ensamblaje")
}

func paintProcess(currentCar int) {
	formatProcessMessage(true, "pintura de la carrocería", currentCar)
	time.Sleep(TimePaintProcess * time.Second)
	formatProcessMessage(false, "pintura de la carrocería", currentCar)
}

func tireProcess(currentCar int) {
	formatProcessMessage(true, "ensamble de los neumáticos", currentCar)
	for currentTire := 0; currentTire < 4; currentTire++ {
		fmt.Print("\t")
		formatProcessMessage(true, fmt.Sprintf("Ensamble del neumático %d", currentTire+1), currentCar)
		time.Sleep(TimeTireProcess * time.Second)
	}
	formatProcessMessage(false, "ensamble de los neumáticos", currentCar)
}

func interiorProcess(currentCar int) {
	formatProcessMessage(true, "ensamble del interior", currentCar)
	time.Sleep(TimeInteriorProcess * time.Second)
	formatProcessMessage(false, "ensamble del interior", currentCar)
}

func formatProcessMessage(isStart bool, process string, currentCar int) {
	var prefix string
	if isStart {
		prefix = "Inicia"
	} else {
		prefix = "Finaliza"
	}
	fmt.Printf("%s proceso de %s, Carro: %v\n", prefix, process, currentCar+1)
}
