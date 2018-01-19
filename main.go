package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func IsBalanced(input string) bool {
	var expected rune

	for _, char := range input {
		if char == '{' {
			expected = '}'
		} else if expected != 0 && char != expected {
			return false
		} else if char == expected {
			expected = 0
		}
	}

	if expected != 0 {
		return false
	}
	return true
}
