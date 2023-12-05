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

}
