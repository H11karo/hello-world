package main

import (
	"fmt"
)

func main() {

	var n1, n2 int64
	var o string

	fmt.Println("Hello, this is a calculator")
	fmt.Println("First Number:")
	fmt.Scanln(&n1)

	fmt.Println("+,-,*,/")
	fmt.Scanln(&o)

	fmt.Println("Second Number:")
	fmt.Scanln(&n2)

	switch o {

	case "+":
		fmt.Printf("%d %s %d = %d", n1, or, n2, n1+n2)

	case "-":
		fmt.Printf("%d %s %d = %d", n1, o, n2, n1-n2)

	case "*":
		fmt.Printf("%d %s %d = %d", n1, o, n2, n1*n2)

	case "/":
		if n2 == 0 {
			fmt.Println("That's not possible, dumbass")
		} else {
			fmt.Printf("%d %s %d = %d", n1, o, n2, n1/n2)
		}
	default:
		fmt.Println("Please choose a valid operator")
	}

}
