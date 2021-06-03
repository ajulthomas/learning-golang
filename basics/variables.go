package main

import "fmt"

func main(){
	var name string = "Ajul Thomas"
	fmt.Println("Welcome ", name)

	// shorthand go compiler can infer what is the type of value assigned
	age := 12
	fmt.Println(age);

	// another way of variable declaration and initialisation
	var city = "Banglore"
	fmt.Println(city)

	var signal bool = true
	signal = false
	fmt.Println("signal",signal)

	var count uint8 = 36
	fmt.Println("count",count)

	distance := 250;
	fmt.Println("Distance is ", distance, "kms")
	
	// constants in go
	const Pi float32 = 3.14
	fmt.Println("Value of Pi is",Pi)

	// multiple variable declaration in go
	var (
		a = 10
		b = 20
		c = 30
	)
	fmt.Println("a=",a)
	fmt.Println("b=",b)
	fmt.Println("c=",c)

}