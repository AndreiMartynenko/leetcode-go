/*
ARRAYS AND STRINGS
Sliding Window

Maximum Average Subarray I

You are given an integer array nums
consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal
to k that has the maximum average value and
return this value. Any answer with a calculation
error less than 10-5 will be accepted.



Example 1:

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

Example 2:

Input: nums = [5], k = 1
Output: 5.00000
*/

package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	currSum := 0
	maxSum := math.MinInt64

	for i := 0; i < len(nums); i++ {
		currSum += nums[i]

		if i >= k-1 {
			maxSum = max(currSum, maxSum)
			currSum -= nums[i-(k-1)]
		}
	}

	return float64(maxSum) / float64(k)
}

/*
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

*/
