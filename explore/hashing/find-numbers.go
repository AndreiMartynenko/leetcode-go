/*
Hashing

Example 3: Given an integer array nums,
find all the numbers x in nums that
satisfy the following: x + 1 is not in nums,
and x - 1 is not in nums.

*/

package main

func findNumbers(nums []int) []int {

	seen := make(map[int]bool)

	// First pass: populate the map with all numbers in the array
	for _, num := range nums {
		seen[num] = true
	}

}
