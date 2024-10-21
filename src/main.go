package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	var reverse string

	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	if reverse == s {
		fmt.Println("Is palindrome")
		return true
	} else {
		fmt.Println("Not a palindrome")
		return false
	}
}
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

	// PRINTF %d = integer %s = string, %v = variable
	nombre := "Pedro"
	edad := 25
	fmt.Printf("Mi nombre es %s y tengo %d anios \n", nombre, edad)
	// Short Exercise: How to use For
	counter := 10
	for counter >= 0 {
		if counter == 0 {
			break
		}
		counter--
		fmt.Println(counter)
	}

	// password and username validation and is pair or not
	const pass string = "1234"
	const user string = "test"
	isPair := func(n int) bool {
		if n%2 == 0 {
			return true
		} else {
			return false
		}
	}

	validation := func(user string, pass string) bool {

		if user == "test" && pass == "1234" {
			return true
		} else {
			return false
		}
	}
	fmt.Println(validation(user, pass))
	fmt.Println(isPair(2))

	slice := []string{"test", "pedro", "gaga"}
	for i, e := range slice {
		fmt.Println(i, e)
	}
	isPalindrome("hello")
	isPalindrome("kayak")
}
