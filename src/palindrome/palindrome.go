package palindrome

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {
	var reverse string
	var convertString string

	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	convertString += strings.ToLower(reverse)
	if convertString == s {
		fmt.Println("Is palindrome")
		return true
	} else {
		fmt.Println("Not a palindrome")
		return false
	}
}
