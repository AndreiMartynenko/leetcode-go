/*
ARRAYS AND STRINGS
Two Pointers

Reverse string
Write a function that reverses a string.
The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
Example 2:

Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]


Constraints:

1 <= s.length <= 105
s[i] is a printable ascii character.
*/

package main

import "fmt"

/*
func reverseString(s []byte) {
	leftCharacter := 0
	rightCharacter := len(s) - 1
	for leftCharacter < rightCharacter {
		// Swap the characters at leftCharacter and rightCharacter positions
		s[leftCharacter], s[rightCharacter] = s[rightCharacter], s[leftCharacter]

		// Move the leftCharacter pointer towards the right
		leftCharacter++

		// Move the rightCharacter pointer towards the left
		rightCharacter--
	}
}

*/

func reverseString(s []byte) []byte {
	left := 0
	right := len(s) - 1
	//swapedArr := []string{}

	//i, j := 0, 0

	for left < right {
		s[left], s[right] = s[right], s[left]
		fmt.Println(left, right)
		left++
		right--

	}

	return s
}

func main() {
	var testWord = "Hello World!"
	var newWord = reverseString([]byte(testWord))
	fmt.Println(string(newWord))
}
