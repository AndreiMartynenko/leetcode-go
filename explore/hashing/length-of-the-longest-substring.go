/*
HASHING
Counting

Example 1: You are given a string s and an integer k.
Find the length of the longest substring that contains
at most k distinct characters.

For example, given s = "eceba" and k = 2, return 3.
The longest substring with at most 2 distinct
characters is "ece".
*/

package main

func find_longest_substring(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}

	charCount := make(map[byte]int)
	start, maxLen := 0, 0

	for end := 0; end < len(s); end++ {
		charCount[s[end]]++

		for len(charCount) > k {
			charCount[s[start]]--
			if charCount[s[start]] == 0 {
				delete(charCount, s[start])
	

}
