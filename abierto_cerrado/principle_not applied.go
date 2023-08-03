package abierto_cerrado

import "fmt"

type HamburgerCocaCola struct {
	hamburgerType string
}

func NewHamburgerCocaCola(hamburgerType string) *HamburgerCocaCola {
	return &HamburgerCocaCola{hamburgerType: hamburgerType}
}

func (hc *HamburgerCocaCola) GetDescription() string {
	return fmt.Sprintf("%s en combo con CocaCola y Papas a la Francesa", hc.hamburgerType)
}

type HamburgerQuatro struct {
	hamburgerType string
}

func NewHamburgerQuatro(hamburgerType string) *HamburgerQuatro {
	return &HamburgerQuatro{hamburgerType: hamburgerType}
}

func (hq *HamburgerQuatro) GetDescription() string {
	return fmt.Sprintf("%s en combo con Quatro y Anillos de cebolla", hq.hamburgerType)
}

type HamburgerLemonate struct {
	hamburgerType string
}

func NewHamburgerLemonate(hamburgerType string) *HamburgerLemonate {
	return &HamburgerLemonate{hamburgerType: hamburgerType}
}

func (hl *HamburgerLemonate) GetDescription() string {
	return fmt.Sprintf("%s en combo con Limonada", hl.hamburgerType)
}

func MakeMenuWithoutPrincipleApplied() {
	hamburgerComboCocaCola := NewHamburgerCocaCola("Hamburguesa Clásica")
	hamburgerComboQuatro := NewHamburgerQuatro("Hamburguesa Clásica")
	hamburgerComboLimonada := NewHamburgerLemonate("Hamburguesa Clásica")

	fmt.Println("Menú")
	fmt.Println("1. HamburgerComboCocaCola:")
	fmt.Println("Consta de:", hamburgerComboCocaCola.GetDescription())

	fmt.Println("2. HamburgerComboQuatro:")
	fmt.Println("Consta de:", hamburgerComboQuatro.GetDescription())

	fmt.Println("3. HamburgerComboLimonada:")
	fmt.Println("Consta de:", hamburgerComboLimonada.GetDescription())
}
