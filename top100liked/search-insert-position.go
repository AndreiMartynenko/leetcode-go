/*
Given a sorted array of distinct integers and a target value,
return the index if the target is found.
If not, return the index where it would be
if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Input: nums = [1,3,5,6], target = 5
Output: 2

Input: nums = [1,3,5,6], target = 2
Output: 1

Input: nums = [1,3,5,6], target = 7
Output: 4

Related Topics
Array, Binary Search
*/

package main

func searchInsert(nums []int, target int) int {
	// Initialize the lower and upper bounds of the search range
	var lo = 0
	var hi = len(nums)

	// Perform binary search
	for i := 0; i < len(nums); i++ {
		// Check if the lower bound is still less than the upper bound
		if lo < hi {
			// Calculate the middle index of the current search range
			var mid = lo + ((hi - lo) / 2)

			// Compare the target value with the middle element
			if target > nums[mid] {
				// If the target is greater than the middle element,
				// adjust the lower bound to search in the right half
				lo = mid + 1
			} else {
				// If the target is less than or equal to the middle element,
				// adjust the upper bound to search in the left half
				hi = mid
			}
		}
	}

	// Return the lower bound as the index where the target should be inserted
	return lo
}
