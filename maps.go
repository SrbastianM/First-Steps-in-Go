package main

import "fmt"

func main() {
	// Create a map
	// To create an map you need to declared the func Make()
	m := make(map[string]int)
	m["key"] = 14
	m["marco"] = 12

	// We pass trow the key and the value
	for k, v := range m {
		fmt.Println(k, v)
	}
	// to found one value, passing another variable (ok) to know if the key exists
	value, ok := m["marco"]
	fmt.Println(value, ok)

}
