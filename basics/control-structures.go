package main

import "fmt"

func main() {

	defer fmt.Println("Testing Defer Statements, First Defer Statement")
	defer fmt.Println("Last Defer Statement")

	// switch without condition is a clean way to execute long if else chains
	// evaluates from top to bottom
	switch {
	case 1 < 2:
		defer fmt.Println("To test how Defer Statements work inside a switch Statement")
		fmt.Println("can evaluate expression in a case")
	case false:
		fmt.Println("Prints first true statement")
	case true:
		fmt.Println("Prints first true statement")
	default:
		fmt.Println("Default condition gets executed if non of the cases evaluates to true")
	}

	fmt.Println("Prints numbers from 1 to 10 using for loop")
	printRange(1, 10, "")
	fmt.Println("Prints odd numbers from 1 to 10")
	printRange(1, 20, "odd")
	fmt.Println("Prints even numbers from 1 to 30")
	printRange(1, 30, "even")
}

func printRange(l int32, u int32, filterType string) {
	for i := l; i <= u; i++ {
		switch filterType {
		case "odd":
			if i%2 != 0 {
				fmt.Println(i)
			}
		case "even":
			if i%2 == 0 {
				fmt.Println(i)
			}
		default:
			fmt.Println(i)
		}
	}
}
