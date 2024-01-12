/*
ARRAYS AND STRINGS
Two Pointers

Example 1: Given a string s, return true
if it is a palindrome, false otherwise.

A string is a palindrome if it reads the
same forward as backward. That means, after
reversing it, it is still the same string.
For example: "abcdcba", or "racecar".

*/

package main

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/*
func main() {
	newString := "racecar"
	result := isPalindrome(newString)
	fmt.Println(result)
}
*/
