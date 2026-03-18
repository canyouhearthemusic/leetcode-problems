package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	maxLen := 0
	left := 0

	// abcadbc
	for right := 0; right < len(s); right++ {
		if idx, exists := lastSeen[s[right]]; exists && idx >= left {
			left = idx + 1
		}

		lastSeen[s[right]] = right

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
