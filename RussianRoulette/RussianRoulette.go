package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
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
	input = strings.ToLower(input)

	switch input {

	case "yes":
		fmt.Println("Good luck!")
		random()

		if randNum == 6 {
			fmt.Println("You lost!")
		} else {
			fmt.Println("You survived!")
		}

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
