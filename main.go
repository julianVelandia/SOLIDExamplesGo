package main

import (
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/SOLIDExamplesGo/abierto_cerrado"
	"github.com/julianVelandia/EDteam/SOLIDyHexagonal/SOLIDExamplesGo/responsabilidad_unica"
)

func main() {
	//Principio de Responsabilidad única
	responsabilidad_unica.ProductionLineWithoutPrinciple()
	responsabilidad_unica.ProductionLineWithPrinciple()

	//Principio de Abierto/Cerrado
	abierto_cerrado.MakeMenu()
	abierto_cerrado.MakeMenuWithoutPrincipleApplied()

	//Principio Sustitución de Liskov

}
