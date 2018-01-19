package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func IsBalanced(input string) bool {
	var expected rune

	for _, char := range input {
		if expected != 0 && char != expected {
			return false
		}
		if char == '{' {
			expected = '}'
		}
	}

	if expected != 0 {
		return false
	}
	return true
}
