package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubstring(s string) int {
	left, maxLength, zeroCount := 0, 0, 0

	for right, char := range s {
		if char == '0' {
			zeroCount++
		}

		for zeroCount > 1 {
			if s[left] == '0' {
				zeroCount--
			}
			left++
		}

		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

func main() {
	s := "1101100111"
	fmt.Println("Length of the longest substring containing only '1':", longestSubstring(s))
}
