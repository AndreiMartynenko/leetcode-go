/*
1. Two Sum
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int) // Create a map to store the index of each number as we iterate over the nums slice.
	for i, num := range nums {    // Iterate over each number in the nums slice.
		complement := target - num             // Calculate the complement of the current number with respect to the target value.
		if j, ok := indexMap[complement]; ok { // Check if the complement exists as a key in the indexMap.
			return []int{j, i} // If so, return a slice containing the indices of the complement and the current number.
		}
		indexMap[num] = i // Otherwise, store the index of the current number in the indexMap.
	}
	return nil // If no two numbers add up to the target value, return nil.
}

func main() {
	arg1 := []int{2, 7, 11, 15}  // Define the input slice of numbers.
	arg2 := 9                    // Define the target value.
	result := twoSum(arg1, arg2) // Call the twoSum function with the input and target values.
	fmt.Println(result)          // Print the result of the twoSum function.
}
