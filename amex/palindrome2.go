package main

/*
A palindrome is a ward which reads the same
backward as forward. Some examples of palindromes
are “kayak”, “radar”, “mom”

Write a function
func Solution(N int, K int) string

that given two integers N and K,
returns a palindrome of length N which
consists of K distinct lowercase letters (a-z).

Examples
given N = 5, K= 2  your function may return “abbba”.
there are many other possibilities for example it could also return “zdzdz”.
on the other hand “aaaaa”  is an incorrect answer is it contains only one distinct letter.
given N= 8, K = 3, Your function may return “ppsccspp”.
given N = 3, K = 2, Your function may return “opo”.

Assume that

N  Is an integer within the range [1..200];
K  is an integer within the range [1..26];

creation of the required palindrome is always possible.
in use solution focus on correctness the performance
of your solution will not be the focus of the assessment.

*/

import (
	"fmt"
	"strings"
)

// Solution generates a palindrome of length N with K distinct lowercase letters.
func Solution2(N, K int) string {
	if N <= 0 || K <= 0 || K > 26 {
		// Invalid input
		return ""
	}

	// Create a string with K distinct letters
	letters := "abcdefghijklmnopqrstuvwxyz"[:K]

	// Build the first half of the palindrome
	firstHalf := letters[:N/2]

	// Create the second half by reversing the first half
	secondHalf := reverseString(firstHalf)

	// If N is odd, append the middle letter
	if N%2 == 1 {
		middleLetter := string(letters[K-1])
		return firstHalf + middleLetter + secondHalf
	}

	return firstHalf + secondHalf
}

// reverseString reverses a given string
func reverseString(s string) string {
	var reversed strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		reversed.WriteByte(s[i])
	}
	return reversed.String()
}

func main() {
	result := Solution2(9, 3)
	fmt.Println(result)
}
