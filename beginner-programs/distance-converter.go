/**

Distance Converter
------------------
program that converts from feet into meters. (1 ft = 0.3048 m)

*/

package main

import "fmt"

func main() {
	fmt.Println("Enter distance(in feet) : ");
	var distF float64;
	fmt.Scanf("%f",&distF);
	distM := feetToMeters(distF);
	fmt.Println("Distance in meters is ", distM)
}

func feetToMeters(f float64) float64 {
	const FACTOR float64 = 0.3048;
	m := f * FACTOR;
	return m;
}