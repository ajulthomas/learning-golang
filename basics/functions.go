package main

import "fmt"

func main() {
	var (
		num1 = 10
		num2 = 12
	)
	fmt.Printf("%v + %v = %v \n", num1, num2, sum(5, 10))
	divident := 58
	divisor := 6
	q, r := divide(divident, divisor)
	fmt.Printf("%d / %d gives Quotient = %f, Remainder = %f\n", divident, divisor, q, r)
	fmt.Printf("%d ^ 2 = %d", num1, sqr(num1))
}

// normal functions
func sum(x int, y int) int {
	return x + y
}

// multiple return values
func divide(dividend, divisor int) (float64, float64) {
	quotient := float64(dividend / divisor)
	reminder := float64(dividend % divisor)
	return quotient, reminder
}

// naked return
func sqr(x int) (xSqr int) {
	xSqr = x * x
	return
}

/**
## Go's basic types are

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

## Zero Values

 Variables declared without an explicit initial value are given their zero value.

The zero value is:

    0 for numeric types,
    false for the boolean type, and
    "" (the empty string) for strings.

*/
