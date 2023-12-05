/*
HASHING
Contiguous Array

Given a binary array nums, return the maximum length of
a contiguous subarray with an equal number of 0 and 1.

Example 1:

Input: nums = [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with an equal number of 0 and 1.
Example 2:

Input: nums = [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.

Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.

*/

package main

import "fmt"

func findMaxLength(nums []int) int {
	sumToIndex := make(map[int]int)
	maxLen, sum := 0, 0

	sumToIndex[0] = -1 // Initialize with a sum of 0 at index -1

	for i, num := range nums {
		if num == 0 {
			sum -= 1
		} else {
			sum += 1
		}

		if index, exists := sumToIndex[sum]; exists {
			currentLen := i - index
			if currentLen > maxLen {
				maxLen = currentLen
			}
		} else {
			sumToIndex[sum] = i
		}
	}

	return maxLen

}

func main() {
	// Example 1:
	nums1 := []int{0, 1}
	result1 := findMaxLength(nums1)
	fmt.Println("Example 1 - Maximum Length of Contiguous Subarray:", result1)

	// Example 2:
	nums2 := []int{0, 1, 0}
	result2 := findMaxLength(nums2)
	fmt.Println("Example 2 - Maximum Length of Contiguous Subarray:", result2)
}
