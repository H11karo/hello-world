package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var randNum int

func main() {
	var input string

GameStart:

	fmt.Println("Hello and welcome to this Russian Roulette session.")
	fmt.Println("There are six Chambers, but only one is loaded")
	fmt.Println("Ready to try your luck?")
	fmt.Println("To start type in 'YES'")
	fmt.Println("To cancel type in 'NO'")
	fmt.Scanln(&input)

	switch input {

	case "YES":
		fmt.Println("Good luck!")
		random()

		switch randNum {

		case 6:
			fmt.Println("You lost!")

		default:
			fmt.Println("You survived!")
		}

	case "Yes":
		fmt.Println("Good luck!")
		random()

		switch randNum {

		case 6:
			fmt.Println("You lost!")

		default:
			fmt.Println("You survived!")
		}

	case "yes":
		fmt.Println("Good luck!")
		random()

		switch randNum {

		case 6:
			fmt.Println("You lost!")

		default:
			fmt.Println("You survived!")
		}

		///////////////////////////////////////////////////////////////////////////////////////////////////////

	case "NO":
		fmt.Println("Boooooring")
		os.Exit(0)

	case "No":
		fmt.Println("Boooooring")
		os.Exit(0)

	case "no":
		fmt.Println("Boooooring")
		os.Exit(0)

	default:
		fmt.Println("Input not accepted, try again")
		goto GameStart

	}

}

func random() {
	rand.Seed(time.Now().UnixNano())
	randNum = rand.Intn(7-1) + 1
	fmt.Println("Your Number is", randNum)
}
