package main

import (
	"fmt"
	"strings"
)

type stack []rune

func (s stack) Empty() bool { return len(s) == 0 }
func (s *stack) Put(i rune) { (*s) = append((*s), i) }
func (s *stack) Pop() rune {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}

func main() {
	fmt.Println("vim-go")
}

func IsBalanced(input string) bool {
	opening := "{["
	closing := "}]"
	var expected stack

	for _, char := range input {
		if i := strings.IndexRune(opening, char); i != -1 {
			expected.Put([]rune(closing)[i])
		} else if strings.ContainsRune(closing, char) && !expected.Empty() && char != expected.Pop() {
			return false
		}
	}

	if !expected.Empty() {
		return false
	}
	return true
}
