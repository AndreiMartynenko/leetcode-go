package main

/*
Given two strings ransomNote and magazine, return true
if ransomNote can be constructed by using the letters
from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.

Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false

Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false

Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true


Constraints:

1 <= ransomNote.length, magazine.length <= 105
ransomNote and magazine consist of lowercase English letters.

*/

func canConstruct(ransomNote string, magazine string) bool {
	// rune i used here as I am not sure about Upper or Lowercase of the input words
	// If I know for sure I can use bytes instead
	magazineCount := make(map[rune]int)

	for _, char := range magazine {
		magazineCount[char]++
	}

	for _, char := range ransomNote {
		if magazineCount[char] == 0 {

			return false
		}

		magazineCount[char]--
	}

	return true

}
