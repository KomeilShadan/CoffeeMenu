package main

import (
	"coffeeMenu/coffeeMenu"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"strconv"
)

func main() {
	var err = keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffeeMenu.Menu()

	for {
		var char, key, _ = keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyEsc {
			break
		}

		if key == keyboard.KeyArrowRight {

		}

		i, _ := strconv.Atoi(string(char))

		if _, exists := coffeeMenu.Coffees[i]; !exists {

			fmt.Println("Please select a menu item number.")
		} else {
			fmt.Println("You chose", coffeeMenu.Coffees[i])
		}
	}
	fmt.Println("Program exiting...")
}
