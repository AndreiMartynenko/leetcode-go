/*
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

func minStartValue(nums []int) int {
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
