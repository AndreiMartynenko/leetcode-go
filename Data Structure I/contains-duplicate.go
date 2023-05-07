/*217. Contains Duplicate
Given an integer array nums, return true if any value appears at least
twice in the array, and return false if every element is distinct.
Input: nums = [1,2,3,1]
Output: true
Input: nums = [1,2,3,4]
Output: false
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

package main

import "fmt"

func containsDuplicate(nums []int) bool {
	// Create a map to store the occurrence of each number
	set := make(map[int]bool)

	// Loop through the numbers in the array
	for _, num := range nums {
		// If the number already exists in the map, it's a duplicate
		if set[num] {
			return true
		} else {
			// If it doesn't exist, add it to the map
			set[num] = true
		}
	}
	// If no duplicates were found, return false
	return false
}

func main() {
	arr1 := []int{1, 2, 3, 1}
	result := containsDuplicate(arr1)
	fmt.Println(result)
}
