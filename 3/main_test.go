package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "abcadbcbb",
			expected: 4,
		},
		{
			input:    "bbbbb",
			expected: 1,
		},
		{
			input:    "pwwkew",
			expected: 3,
		},
	}

	for _, tt := range tests {
		actual := lengthOfLongestSubstring(tt.input)

		if actual != tt.expected {
			t.Errorf("lengthOfLongestSubstring(%q) = %d, expected %d", tt.input, actual, tt.expected)
		}
	}
}
