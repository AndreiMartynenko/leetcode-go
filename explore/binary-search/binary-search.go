/*
Example 1: 704. Binary Search

You are given an array of integers nums which is sorted in ascending order,
and an integer target. If target exists in nums, return its index. Otherwise,
return -1.

*/

package main

import "fmt"

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := ((left + right) / 2)

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	newArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarySearch(newArr, 6)
	fmt.Println(result)
}
