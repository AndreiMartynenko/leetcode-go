/*
HASHING

Maximum Number of Balloons
Given a string text, you want to use the characters of text to form
as many instances of the word "balloon" as possible.

You can use each character in text at most once.
Return the maximum number of instances that can be formed.

Example 1:

Input: text = "nlaebolko"
Output: 1
Example 2:

Input: text = "loonbalxballpoon"
Output: 2
Example 3:

Input: text = "leetcode"
Output: 0


Constraints:

1 <= text.length <= 104
text consists of lower case English letters only.
*/

package main

import (
	"math"
)

func maxNumberOfBalloons(text string) int {

	balloonCounts := map[rune]int{
		'b': 0,
		'a': 0,
		'l': 0,
		'o': 0,
		'n': 0,
	}

	for _, char := range text {
		if count, exists := balloonCounts[char]; exists {
			balloonCounts[char] = count + 1
		}
	}

	// Adjust the counts for 'l' and 'o' as they need to appear twice in the word "balloon"
	balloonCounts['l'] /= 2
	balloonCounts['o'] /= 2

	// Find the minimum count among the characters in "balloon"
	minCount := math.MaxInt32
	for _, count := range balloonCounts {
		if count < minCount {
			minCount = count
		}
	}

	return minCount

}

/*

func main() {
	// Example 1
	text1 := "nlaebolko"
	result1 := maxNumberOfBalloons(text1)
	fmt.Println(result1) // Output: 1

	// Example 2
	text2 := "loonbalxballpoon"
	result2 := maxNumberOfBalloons(text2)
	fmt.Println(result2) // Output: 2

	// Example 3
	text3 := "leetcode"
	result3 := maxNumberOfBalloons(text3)
	fmt.Println(result3) // Output: 0
}

*/
