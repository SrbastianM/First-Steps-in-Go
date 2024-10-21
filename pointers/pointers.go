package main

import "fmt"

func main() {
	// to access the memory direction you need to use "&" before the variable
	// and if you want to access the value you need to use "*"
	a := 10
	b := &a
	c := *b

	fmt.Println(a, b, c)
}
