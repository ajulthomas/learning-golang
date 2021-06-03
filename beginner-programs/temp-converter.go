/*

Farenheit to celsius converter
-------------------------------

Program that converts from Fahrenheit
into Celsius. (C = (F - 32) * 5/9)

*/

package main

import "fmt"

func main() {
	fmt.Println("Please enter tempurature(in F) : ");
	var tempF float64;
	fmt.Scanf("%f", &tempF);
	tempC := farenheitToCelsius(tempF);
	fmt.Println("Temperature in Degree Celsius is ", tempC);
}

/**
Converts tempurature in Farenheit to Degree Celsius
@input f { float64 } Tempurature in Farenheit
*/
func farenheitToCelsius(f float64) float64 {
	var c float64 = ((f-32)*5)/9;
	return c;
}