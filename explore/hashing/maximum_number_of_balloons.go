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

}