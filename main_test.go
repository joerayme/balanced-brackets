package main

import (
	"testing"
)

func TestIsBalancedReturnsTrueWithEmptyInput(t *testing.T) {
	if IsBalanced("") != true {
		t.Error("Expected IsBalanced to return true for empty input")
	}
}
