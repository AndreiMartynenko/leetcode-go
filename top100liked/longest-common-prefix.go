/*
14. Longest Common Prefix

Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Related Topics
String, Trie
*/

// Trie
package main

import (
	"fmt"
	"strings"
)

// longestCommonPrefix finds the longest common prefix string amongst an array of strings.
// If there is no common prefix, it returns an empty string.
func longestCommonPrefix(strs []string) string {
	// If the array is nil or empty, return an empty string
	if strs == nil || len(strs) == 0 {
		return ""
	}
	// Initialize prefix as the first string in the array
	var prefix = strs[0]
	// Iterate through the array starting from the second string
	for i := 1; i < len(strs); i++ {
		// While the prefix is not a substring of the current string,
		// remove the last character from the prefix and check again
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			// If the prefix becomes empty, there is no common prefix, return an empty string
			if prefix == "" {
				return ""
			}
		}
	}
	// Return the longest common prefix found
	return prefix
}

func main() {
	var arr = []string{"flower", "flow", "flight"}
	var result = longestCommonPrefix(arr)
	fmt.Println(result)
}
