package main

import "fmt"

func main() {
	// Arrays - sequence of elements of same type with fixed length
	var arr [5]string
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)

	// calculate length of array
	fmt.Println("Length of array = ", len(arr))

	// go provides another form of initialiation and declaration
	newArr := [5]int{1, 2, 3, 4, 5}
	total := 0
	for _, num := range newArr {
		total += num
	}
	fmt.Println("sum of", newArr, " = ", total)
}
