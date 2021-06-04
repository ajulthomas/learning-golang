package main

import "fmt"

func main(){

	fmt.Println("Prints numbers from 1 to 10 using for loop")
	printRange(1,10,"")
	fmt.Println("Prints odd numbers from 1 to 10")
	printRange(1,20,"odd")
	fmt.Println("Prints even numbers from 1 to 30")
	printRange(1,30,"even")
}

func printRange(l int32, u int32, filterType string) {
	for i:=l; i<=u; i++ {
		switch filterType {
		case "odd": if i % 2 != 0 { fmt.Println(i) }
		case "even": if i % 2 == 0 { fmt.Println(i) }
		default: fmt.Println(i)	
		}
	}
}