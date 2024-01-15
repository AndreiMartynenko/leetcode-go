/*
ARRAYS AND STRINGS
Prefix Sum

K Radius Subarray Averages

You are given a 0-indexed array nums of n integers, and an integer k.
The k-radius average for a subarray of nums centered at some index i
with the radius k is the average of all elements in nums between
the indices i - k and i + k (inclusive). If there are less than
k elements before or after the index i, then the k-radius average is -1.

Build and return an array avgs of length n where avgs[i] is
the k-radius average for the subarray centered at index i.

The average of x elements is the sum of the x elements divided by x,
using integer division. The integer division truncates toward zero,
which means losing its fractional part.

For example, the average of four elements 2, 3, 1, and
5 is (2 + 3 + 1 + 5) / 4 = 11 / 4 = 2.75, which truncates to 2.

Example 1

Input: nums = [7,4,3,9,1,8,5,2,6], k = 3
Output: [-1,-1,-1,5,4,4,-1,-1,-1]
Explanation:
- avg[0], avg[1], and avg[2] are -1 because there are
less than k elements before each index.
- The sum of the subarray centered at index 3 with
radius 3 is: 7 + 4 + 3 + 9 + 1 + 8 + 5 = 37.
  Using integer division, avg[3] = 37 / 7 = 5.
- For the subarray centered at index 4,
avg[4] = (4 + 3 + 9 + 1 + 8 + 5 + 2) / 7 = 4.
- For the subarray centered at index 5,
avg[5] = (3 + 9 + 1 + 8 + 5 + 2 + 6) / 7 = 4.
- avg[6], avg[7], and avg[8] are -1 because there are
less than k elements after each index.

Example 2:
Input: nums = [100000], k = 0
Output: [100000]
Explanation:
- The sum of the subarray centered at index 0
with radius 0 is: 100000.
  avg[0] = 100000 / 1 = 100000.

Example 3:

Input: nums = [8], k = 100000
Output: [-1]
Explanation:
- avg[0] is -1 because there are less than k
elements before and after index 0.

*/

package main

// Approach 1: Prefix Sum

func getAverages1(nums []int, k int) []int {
	// When a single element is considered then its average will be the number itself only.
	if k == 0 {
		return nums
	}

	n := len(nums)
	averages := make([]int, n)
	for i := range averages {
		averages[i] = -1
	}

	// Any index will not have 'k' elements in its left and right.
	if 2*k+1 > n {
		return averages
	}

	// Generate 'prefix' array for 'nums'.
	// 'prefix[i + 1]' will be sum of all elements of 'nums' from index '0' to 'i'.
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(nums[i])
	}

	// We iterate only on those indices which have at least 'k' elements in their left and right.
	// i.e. indices from 'k' to 'n - k'
	for i := k; i < (n - k); i++ {
		leftBound := i - k
		rightBound := i + k
		subArraySum := prefix[rightBound+1] - prefix[leftBound]
		average := int(subArraySum / int64(2*k+1))
		averages[i] = average
	}

	return averages
}

// Approach 2: Sliding Window

func getAverages2(nums []int, k int) []int {
	// When a single element is considered then its average will be the number itself only.
	if k == 0 {
		return nums
	}

	n := len(nums)
	averages := make([]int, n)
	for i := range averages {
		averages[i] = -1
	}

	// Any index will not have 'k' elements in its left and right.
	if 2*k+1 > n {
		return averages
	}

	// First get the sum of the first window of the 'nums' array.
	windowSum := 0
	for i := 0; i < 2*k+1; i++ {
		windowSum += nums[i]
	}
	averages[k] = windowSum / (2*k + 1)

	// Iterate on the rest of the indices which have at least 'k' elements
	// on their left and right sides.
	for i := 2*k + 1; i < n; i++ {
		// We remove the discarded element and add the new element to get the current window sum.
		// 'i' is the index of the newly inserted element, and
		// 'i - (window size)' is the index of the last removed element.
		windowSum = windowSum - nums[i-(2*k+1)] + nums[i]
		averages[i-k] = windowSum / (2*k + 1)
	}

	return averages
}
