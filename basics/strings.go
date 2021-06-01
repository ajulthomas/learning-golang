package main

import "fmt"

func main(){
	fmt.Println("Welcome Gophers :) !!!")
	fmt.Println("Third Character of \"Welcome Gophers\" is ",string("Welcome Gophers"[3]))
	fmt.Println("length of \"Welcome Gophers\" is",len("Welcome Gophers"))
	fmt.Println("Double quotes\tcan\r\n contain \n special characters \n\r")
	fmt.Println(`Welcome Gophers
	this is a 
	multiline string
	literal
	`)
}