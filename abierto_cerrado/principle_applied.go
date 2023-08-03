package abierto_cerrado

import "fmt"

type Combo interface {
	GetDescription() string
}

type Hamburger struct {
	hamburgerType string
	combo         Combo
}

type CocaColaAndFries struct{}

func (c *CocaColaAndFries) GetDescription() string {
	return "CocaCola y Papas a la Francesa"
}

type QuatroAndOnionRings struct{}

func (c *QuatroAndOnionRings) GetDescription() string {
	return "Quatro y Anillos de cebolla"
}

type Lemonade struct{}

func (c *Lemonade) GetDescription() string {
	return "Limonada"
}

func MakeMenu() {
	HamburgerComboCocaCola := Hamburger{
		hamburgerType: "Hamburguesa Clásica",
		combo:         &CocaColaAndFries{},
	}

	HamburgerComboQuatro := Hamburger{
		hamburgerType: "Hamburguesa Clásica",
		combo:         &QuatroAndOnionRings{},
	}

	HamburgerComboLimonada := Hamburger{
		hamburgerType: "Hamburguesa Clásica",
		combo:         &Lemonade{},
	}

	fmt.Println("Menú")
	fmt.Println("1. HamburgerComboCocaCola: ")
	fmt.Println("Consta de: ", HamburgerComboCocaCola.hamburgerType, ", en combo con ", HamburgerComboCocaCola.combo.GetDescription())

	fmt.Println("2. HamburgerComboQuatro: ")
	fmt.Println("Consta de: ", HamburgerComboQuatro.hamburgerType, ", en combo con ", HamburgerComboQuatro.combo.GetDescription())

	fmt.Println("3. HamburgerComboLimonada: ")
	fmt.Println("Consta de: ", HamburgerComboLimonada.hamburgerType, ", en combo con ", HamburgerComboLimonada.combo.GetDescription())

}
