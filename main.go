package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func IsBalanced(input string) bool {
	if input == "{" {
		return false
	}

	return true
}
