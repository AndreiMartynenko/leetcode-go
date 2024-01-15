/*
ARRAYS AND STRINGS
Prefix Sum

Minimum Value to Get Positive Step by Step Sum

Given an array of integers nums,
you start with an initial positive value startValue.

In each iteration, you calculate the step by step
sum of startValue plus elements in nums (from left to right).

Return the minimum positive value of startValue
such that the step by step sum is never less than 1.

Example 1:

Input: nums = [-3,2,-3,4,2]
Output: 5
Explanation: If you choose startValue = 4, in the third iteration your step by step sum is less than 1.
step by step sum
startValue = 4 | startValue = 5 | nums
  (4 -3 ) = 1  | (5 -3 ) = 2    |  -3
  (1 +2 ) = 3  | (2 +2 ) = 4    |   2
  (3 -3 ) = 0  | (4 -3 ) = 1    |  -3
  (0 +4 ) = 4  | (1 +4 ) = 5    |   4
  (4 +2 ) = 6  | (5 +2 ) = 7    |   2

  Example 2:

Input: nums = [1,2]
Output: 1
Explanation: Minimum start value should be positive.

Example 3:

Input: nums = [1,-2,-3]
Output: 5
*/

package main

// Approach 1: Brute Force
func minStartValue1(nums []int) int {
	// Start with startValue = 1.
	startValue := 1

	//While we haven't found the first valid startValue
	for {
		total := startValue
		isValid := true
		// The step-by-step total "total" equals startValue at the beginning.
		// Use boolean parameter "isValid" to record whether the total
		// is larger than or equal to 1.

		// Iterate over the array "nums".
		for i := 0; i < len(nums); i++ {
			total += nums[i]
			if total < 1 {
				isValid = false
				break
			}
		}

		// If "total" is less than 1, we shall try a larger startValue,
		// we mark "isValid" as "false" and break the current iteration.
		if total < 1 {
			if isValid {
				return startValue
			}

		}
		// If "isVaild" is true, meaning "total" is never less than 1 in the
		// iteration, therefore we return this "startValue". Otherwise, we
		// go ahead and try "startValue" + 1 as the new "startValue".
		startValue++
	}
}

// Approach 2: Binary Search

// Function to find the minimum valid start value
func minStartValue2(nums []int) int {
	// Let n be the length of the array "nums", m be the absolute value
	// of the lower boundary of the element. In this question we have m = 100.
	n := len(nums)
	m := 100

	// Set left and right boundaries according to left = 1, right = m * n + 1.
	left := 1
	right := m*n + 1

	// Binary search loop
	for left < right {
		// Get the middle index "middle" of the two boundaries, let the start value
		// be "middle". The initial step-by-step total "total" equals to middle as well.
		// Use boolean parameter "isValid" to record whether the total
		// is greater than or equal to 1.
		middle := (left + right) / 2
		total := middle
		isValid := true

		// Iterate over the array "nums".
		for _, num := range nums {
			// In each iteration, calculate "total" plus the element "num" in the array.
			total += num

			// If "total" is less than 1, we shall try a larger start value,
			// we mark "isValid" as "false" and break the current iteration.
			if total < 1 {
				isValid = false
				break
			}
		}

		// Check if middle is valid, and reduce the search space by half.
		if isValid {
			right = middle
		} else {
			left = middle + 1
		}
	}

	// When the left and right boundaries coincide, we have found
	// the target value, that is, the minimum valid startValue.
	return left
}

// Approach 3: Prefix total

func minStartValue3(nums []int) int {
	// minVal will store the minimum step-by-step total,
	// while total will keep track of the current step-by-step total.
	minVal, total := 0, 0

	// Iterate over the elements in the nums array.
	for i := 0; i < len(nums); i++ {
		// Update the current step-by-step total.
		total += nums[i]

		// Keep track of the minimum step-by-step total.
		if total < minVal {
			minVal = total
		}
	}

	// Ensure the minimum step-by-step total is at least 1 by adjusting the startValue.
	return -minVal + 1
}
