package main

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	for end := 0; end < len(s); end++ {
		if charIndex[s[end]] >= start {
			start = charIndex[s[end]] + 1
		}
		charIndex[s[end]] = end
		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}

	return maxLength
}
