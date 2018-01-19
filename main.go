package main

import (
	"fmt"
	"os"
	"strings"

	"bufio"
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
	s := bufio.NewScanner(os.Stdin)

	// We don't care about the first line of input, it's just how many lines to expect
	s.Scan()

	for s.Scan() {
		if IsBalanced(s.Text()) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func IsBalanced(input string) bool {
	opening := "{[("
	closing := "}])"
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
