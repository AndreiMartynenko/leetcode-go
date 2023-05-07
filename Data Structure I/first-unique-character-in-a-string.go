/*
387. First Unique Character in a String
Given a string s, find the first non-repeating character in it
and return its index. If it does not exist, return -1.

Examples:

Input: s = "leetcode"
Output: 0

Input: s = "loveleetcode"
Output: 2

Input: s = "aabb"
Output: -1
*/

package main

import "fmt"

// func firstUniqChar(s string) int {
// 	charCount := make(map[rune]int) // Create a map to store the character counts

// 	// Count the occurrence of each character in the string
// 	for _, char := range s {
// 		charCount[char]++ // Increment the count of the current character
// 	}

// 	// Find the first non-repeating character
// 	for i, char := range s {
// 		if charCount[char] == 1 { // If the current character has a count of 1
// 			return i // Return the index of the first non-repeating character
// 		}
// 	}

// 	// No non-repeating character found
// 	return -1
// }

//Another approach

func firstUniqChar(s string) int {
	// Create an array of size 26, initialized with 0s to count occurrences of each character
	var helper [26]int
	for i := 0; i < len(s); i++ {
		// Convert the character to its ASCII code and subtract 97 to get its index in the array
		helper[int(s[i])-97]++
	}

	// Find the first non-repeating character by looping through the string again
	for i := 0; i < len(s); i++ {
		if helper[int(s[i])-97] == 1 {
			return i
		}
	}
	// No non-repeating character found
	return -1
}

func main() {
	var word = "loveleetcode"
	result := firstUniqChar(word)
	fmt.Println(result)

}
