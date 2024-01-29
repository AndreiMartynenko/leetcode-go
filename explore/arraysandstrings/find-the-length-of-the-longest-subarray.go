/*
ARRAYS AND STRINGS
Sliding Window

Example 1: Given an array of positive integers
nums and an integer k, find the length of the
longest subarray whose sum is less than or equal to k.

*/

package main

import "fmt"

/*
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

*/

func longestSubArraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	left, currSum, maxLength := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		currSum += nums[right]

		for currSum > k {
			currSum -= nums[left]
			left++
		}
		maxLength = max(maxLength, right-left+1)

	}
	return maxLength

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	newArr := []int{3, 1, 2, 7, 4, 2, 1, 1, 5}
	result := longestSubArraySum(newArr, 8)
	fmt.Println(result)
}
