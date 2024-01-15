/*
ARRAY AND STRINGS
Sliding Window

Example 1: Given an array of positive integers
nums and an integer k, find the length of the
longest subarray whose sum is less than or equal to k.

*/

package main

import (
	"fmt"
	"math"
)

func lengthOfTheLongestSubarray(nums []int, k int) int {
	left, right, curr := 0, 0, 0
	maxLength := 0

	for right < len(nums) {
		curr += nums[right]

		for curr > k {
			curr -= nums[left]
			left++
		}

		maxLength = int(math.Max(float64(maxLength), float64(right-left+1)))

		right++
	}

	return maxLength
}

func main() {
	newArr := []int{3, 1, 2, 7, 4, 2, 1, 1, 5}
	result := lengthOfTheLongestSubarray(newArr, 8)
	fmt.Println(result)
}
