/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order,
and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.
The final sorted array should not be returned by the function,
but instead be stored inside the array nums1.
To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements
that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

Example 1
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]
Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

Example 2
Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]
Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].

Example 3
Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]
Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

Related topics: Array, Two Pointers, Sorting
*/

package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// Initialize three pointers:
	// - p1: points to the last element in nums1
	// - p2: points to the last element in nums2
	// - p: points to the last element in the merged array
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	// While both arrays have elements left
	for p1 >= 0 && p2 >= 0 {
		// Compare the last elements of nums1 and nums2
		// and add the larger one to the end of the merged array
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// If there are remaining elements in nums2,
	// add them to the beginning of the merged array
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}

	return nums1
}

func main() {
	var arr1 = []int{1, 2, 3, 0, 0, 0}
	var num1 = 3
	var arr2 = []int{2, 5, 6}
	var num2 = 3

	result := merge(arr1, num1, arr2, num2)
	fmt.Println(result)
}
