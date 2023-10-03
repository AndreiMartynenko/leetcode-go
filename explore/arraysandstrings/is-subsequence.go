/*
Two Pointers
392. Is Subsequence.
Given two strings s and t, return true if s
is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is
formed from the original string by deleting some
(can be none) of the characters without disturbing
the relative positions of the remaining characters.
(i.e., "ace" is a subsequence of "abcde" while "aec" is not).



Example 1:

Input: s = "abc", t = "ahbgdc"
Output: true

Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false

*/

package main

func isSubsequence(s string, t string) bool {
	// Initialize two pointers, i and j, to track positions in s and t
	i, j := 0, 0

	// Iterate through both strings
	for i < len(s) && j < len(t) {
		// If the characters match, move the pointer in s
		if s[i] == t[j] {
			i++
		}
		// Move the pointer in t regardless
		j++
	}

	// If we reached the end of s, it means s is a subsequence of t
	return i == len(s)
}

/*
func main() {
	s := "ace"
	t := "abcde"
	result := isSubsequence(s, t)
	fmt.Println(result)
}
*/
