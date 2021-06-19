package main

import "fmt"

func main() {
	// slices dynamically sized flexible view into an array
	// slice is a segment of an array
	var a []string
	b := make([]int, 5)

	array := [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	// creating a slice from an array
	var c []int = array[2:4]
	// can omit the limit values
	d := array[3:8]
	e := array[:5]
	f := array[6:]
	g := make([]int, len(f))
	copy(g, f)

	a = append(a, "someValue", "againAnotherValue", "anotherValue")

	fmt.Println(a)

	b = append(b, c...)
	fmt.Printf(" A = %v\n B = %v\n C = %v\n D = %v\n E = %v\n F = %v\n G = %v\n", a, b, c, d, e, f, g)

	fmt.Printf("Looping through A :\n")
	for _, item := range a {
		fmt.Println(item)
	}

}
