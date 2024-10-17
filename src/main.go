package main

import (
	"fmt"
)

func main() {
	const pi float64 = 3.14
	fmt.Println("Hello, World!", pi)
	// how to  variables
	// var num int = 1
	// var num int

	// zero values when not initialized the value is 0 to in - flot64
	// Empty string to string, and false to bool
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// Short Exercise whem a variable is initialize like := go declare and assign the value
	const base int = 10
	const baseRectangle int = 20
	const secondBaseTrapezoid int = 30
	const height int = 10

	areaRectangle := baseRectangle * height
	areaTrapezoid := (secondBaseTrapezoid + base) * height / 2
	area := base * base

	fmt.Println(area, areaRectangle, areaTrapezoid)
}
