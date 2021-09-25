package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

type Circle struct {
	radius float64
}

func (r *Rectangle) area() float64 {
	return r.length * r.breadth
}
func (r *Rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func display(s Shape) {
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	var (
		r      Rectangle
		c      Circle
		inputs []float64
	)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input, _ := strconv.Atoi(s.Text())
		inputs = append(inputs, float64(input))
		if len(inputs) == 3 {
			break
		}
	}
	r.length = inputs[0]
	r.breadth = inputs[1]
	c.radius = inputs[2]

	display(&r)
	display(&c)
}
