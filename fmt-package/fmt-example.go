package main

import (
	"fmt"
)

func main() {
	fmt.Print("ajul thomas")
	fmt.Printf("ajul thomas \n")
	fmt.Println("ajul thomas")

	fmt.Print("Please enter two numbers: ")
	var (
		num1 int
		num2 int
	)
	fmt.Scan(&num1, &num2)
	fmt.Printf("num1 %T = %v, num2 %T = %v \n", num1, num1, num2, num2)
}
