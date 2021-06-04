/**
FizzBuzz Algorithm

Print Fizz if divisible by 3
Print Buzz if divisible by 5
Print FizzBuzz if divisible by 3 and 5

*/

package main

import "fmt"

func main(){
	printFizzBuzz(1,100);
}

func printFizzBuzz(l int32, u int32) {
	for i:=l; i<=u; i++ {
		if i%3==0 && i%5==0 { 
			fmt.Println(i,": FizzBuzz")
		} else if i%3==0 {
			fmt.Println(i,": Fizz")
		} else if i%5==0 {
			fmt.Println(i,": Buzz")
		}
	}
}