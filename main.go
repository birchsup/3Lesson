package main

import (
	"fmt"
)

func main() {

	var a float32
	var b float32
	var c string
	var result float32

	fmt.Println("введите число а")
	fmt.Scanln(&a)

	fmt.Println("введите число b")
	fmt.Scanln(&b)

	fmt.Println("введите операцию ")
	fmt.Scanln(&c)

	switch c {
	case "+":
		result = a + b
		fmt.Printf("%.2f", result)
		break
	case "-":
		result = a - b
		fmt.Printf("%.2f", result)
	case "*":
		result = a * b
		fmt.Printf("%.2f", result)
	case "/":
		if b == 0.0 {
			fmt.Println("Divide by Zero Situation")
		} else {
			result = a / b
			fmt.Printf("%.2f", result)
		}

	default:
		fmt.Println("Invalid Operation")
	}

}
