/*

HASHING

Longest Substring Without Repeating Characters

*/

package main

// Sliding Window approach
func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	for end := 0; end < len(s); end++ {
		if idx, found := charIndex[s[end]]; found && idx >= start {
			start = idx + 1
		}
		charIndex[s[end]] = end
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
