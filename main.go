package main

import (
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/SOLIDExamplesGo/segregacion_interfaz"
)

func main() {
	/*
		//Principio de Responsabilidad única
		responsabilidad_unica.ProductionLineWithoutPrinciple()
		responsabilidad_unica.ProductionLineWithPrinciple()

		//Principio de Abierto/Cerrado
		abierto_cerrado.MakeMenu()
		abierto_cerrado.MakeMenuWithoutPrincipleApplied()

		//Principio Sustitución de Liskov
		sustitucion_liskov_principle_not_applied.MakeDucksWithOutPrincipleApplied()
		sustitucion_liskov.MakeDucksWithPrincipleApplied()


	*/
	//Segregación de la interfaz
	segregacion_interfaz.MakeDucksWithPrincipleApplied()
}
