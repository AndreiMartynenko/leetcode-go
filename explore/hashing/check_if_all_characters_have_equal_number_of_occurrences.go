/*
HASHING
Counting
Example 3: 1941. Check if All Characters
Have Equal Number of Occurrences

Given a string s, determine if all
characters have the same frequency.

For example, given s = "abacbc", return true.
All characters appear twice. Given s = "aaabb",
return false. "a" appears 3 times, "b"
appears 2 times. 3 != 2.

Example 1:

Input: s = "abacbc"
Output: true
Explanation: The characters that appear in s are
'a', 'b', and 'c'. All characters occur 2 times in s.

Example 2:

Input: s = "aaabb"
Output: false
Explanation: The characters that
appear in s are 'a' and 'b'.
'a' occurs 3 times while 'b' occurs 2 times,
which is not the same number of times.

*/

package main

import "unicode"

func checkEqualFrequency(s string) bool {
	frequency := make(map[rune]int)

	// Count the frequency of each character in the string
	for _, char := range s {
		if unicode.IsLetter(char) {
			frequency[char]++
		}
	}

	// Check if all characters have the same frequency

	var firstFrequency int
	for _, freq := range frequency {
		if firstFrequency == 0 {
			firstFrequency = freq
		} else if firstFrequency != freq {
			return false
		}

	}

	return true

}
