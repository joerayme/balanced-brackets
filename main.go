package main

import (
	"fmt"
	"os"

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
	brackets := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	var expected stack

	for _, char := range input {
		if match, ok := brackets[char]; ok {
			expected.Put(match)
		} else if !expected.Empty() && char != expected.Pop() {
			return false
		}
	}

	if !expected.Empty() {
		return false
	}
	return true
}
