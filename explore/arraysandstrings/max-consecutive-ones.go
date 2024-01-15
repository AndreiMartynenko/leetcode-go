/*
ARRAY AND STRINGS
Sliding Window

Max Consecutive Ones III
Given a binary array nums and an integer k,
return the maximum number of consecutive 1's
in the array if you can flip at most k 0's.



Example 1:

Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

Example 2:

Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

*/

package main

func longestOnes(nums []int, k int) int {
	// Initialize variables to keep track of maximum consecutive 1s, left pointer,
	// and count of zeros in the current window.
	maxOnes := 0
	left := 0
	zeroCount := 0

	// Iterate through the array using a right pointer.
	for right, num := range nums {
		// If we encounter a 0, increment zeroCount.
		if num == 0 {
			zeroCount++
		}

		// If zeroCount exceeds k, we need to adjust the window.
		// Move the left pointer to the right until zeroCount is less than or equal to k.
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		// Update maxOnes with the maximum window size encountered.
		maxOnes = max(maxOnes, right-left+1)
	}

	// Return the maximum number of consecutive 1s.
	return maxOnes
}

// Helper function to return the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
