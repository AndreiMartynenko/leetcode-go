/*
Given an integer array nums, return true if any value
appears at least twice in the array, and return false if every element is distinct.
*/

package main

import "fmt"

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, num := range nums {
		if set[num] {
			return true
		} else {
			set[num] = true
		}
	}
	return false
}

func main() {
	arr1 := []int{1, 2, 3, 1}
	result := containsDuplicate(arr1)
	fmt.Println(result)
}
