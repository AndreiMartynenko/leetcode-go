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

// Solution generates a palindrome of length N with K distinct lowercase letters.
func Solution(N, K int) string {
	if N <= 0 || K <= 0 || K > 26 {
		// Invalid input
		return ""
	}

	// Create a string with K distinct letters
	letters := "abcdefghijklmnopqrstuvwxyz"[:K]

	result := make([]byte, N)
	i, j := 0, N-1

	for i <= j {
		result[i], result[j] = letters[i%K], letters[i%K]
		i++
		j--
	}

	return string(result)
}


// Different solutions

package solution

// Solution generates a palindrome of length N with K distinct lowercase letters.
func Solution(N, K int) string {
	if N <= 0 || K <= 0 || K > 26 {
		// Invalid input
		return ""
	}

	// Create a string with K distinct letters
	letters := "abcdefghijklmnopqrstuvwxyz"[:K]

	// Build the first half of the palindrome
	firstHalf := letters[:N/2]

	// Create the second half by reversing the first half and taking N/2 or (N+1)/2 characters based on N being even or odd
	secondHalf := reverseString(firstHalf)[:N/2 + N%2]

	// If N is odd, remove the first character from the second half to avoid duplication
	if N%2 == 1 {
		secondHalf = secondHalf[1:]
	}

	return firstHalf + secondHalf
}

// reverseString reverses a given string
func reverseString(s string) string {
	var reversed string
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}


