package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadStringInput(prompt string) string {

	Prompt(prompt)
	for {
		var input = ReadAndStripStdin()

		if input != "" {
			return input
		} else {
			fmt.Println("Please enter a value.")
		}
	}
}

func ReadIntInput(prompt string) int {

	Prompt(prompt)
	for {
		var input = ReadAndStripStdin()

		if input == "" {
			fmt.Println("Please enter a value.")

		} else if number, err := strconv.Atoi(input); err != nil {

			fmt.Println("Please enter a whole number.")

		} else
		{
			return number
		}
	}
}

func ReadFloatInput(prompt string) float64 {

	Prompt(prompt)
	for {
		var input = ReadAndStripStdin()

		if input == "" {
			fmt.Println("Please enter a value.")

		} else if number, err := strconv.ParseFloat(input, 64); err != nil {

			fmt.Println("Please enter a number.")

		} else
		{
			return number
		}
	}
}

func ReadBoolInput(prompt string) bool {

	Prompt(prompt)
	for {
		var input = ReadAndStripStdin()

		if input == "" {
			fmt.Println("Please enter a value.")

		} else if number, err := strconv.ParseBool(input); err != nil {

			fmt.Println("Please enter a number.")

		} else
		{
			return number
		}
	}
}

func StripCRLF(input string) string {

	input = strings.Replace(input, "\r\n", "", -1)
	input = strings.Replace(input, "\n", "", -1)
	return input
}

func Prompt(s string) {
	fmt.Println(s)
	fmt.Print("-> ")
}

func ReadAndStripStdin() string {

	var input, _ = reader.ReadString('\n')
	return StripCRLF(input)
}
