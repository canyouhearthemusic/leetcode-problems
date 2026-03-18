package main

import (
	"fmt"
	"testing"
)

func TestFindLongestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "babad",
			expected: "bab",
		},
		{
			input:    "cbbd",
			expected: "bb",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i+1), func(t *testing.T) {
			actual := longestPalindrome(tt.input)

			if tt.expected != actual {
				t.Errorf("got %v; wanted %v", actual, tt.expected)
			}
		})
	}
}
