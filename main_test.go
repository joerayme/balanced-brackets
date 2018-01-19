package main

import (
	"testing"
)

func TestIsBalancedReturnsTrueWithEmptyInput(t *testing.T) {
	if IsBalanced("") != true {
		t.Error("Expected IsBalanced to return true for empty input")
	}
}

func TestIsBalancedReturnsFalseWithSingleBracket(t *testing.T) {
	if IsBalanced("{") != false {
		t.Error("Expected IsBalanced to return false when single bracket is passed")
	}
}

func TestIsBalancedReturnsTrueWhenNoBrackets(t *testing.T) {
	if IsBalanced("a") != true {
		t.Error("Expected IsBalanced to return true when no brackets are passed")
	}
}
