package coffeeMenu

import "fmt"

var Coffees = map[int]string{
	1: "Cappucino",
	2: "Latte",
	3: "Americano",
	4: "Mocha",
	5: "Macchiato",
	6: "Espresso",
}

func Menu() {
	fmt.Println("Menu")
	fmt.Println("--------")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Esc - Quit the program")
}
