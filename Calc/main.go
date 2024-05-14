package main

import (
	"fmt"
)

func main() {

	var number1, number2 int64
	var operator string

	fmt.Println("Hello, this is a calculator")
	fmt.Println("First Number:")
	fmt.Scanln(&number1)

	fmt.Println("+,-,*,/")
	fmt.Scanln(&operator)

	fmt.Println("Second Number:")
	fmt.Scanln(&number2)

	switch o {

	case "+":
		fmt.Printf("%d %s %d = %d", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%d %s %d = %d", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%d %s %d = %d", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0 {
			fmt.Println("That's not possible, dumbass")
		} else {
			fmt.Printf("%d %s %d = %d", number1, operator, number2, number1/number2)
		}
	default:
		fmt.Println("Please choose a valid operator")
	}

}
