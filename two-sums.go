/*
1. Two Sum
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
*/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := indexMap[complement]; ok {
			return []int{j, i}
		}
		indexMap[num] = i
	}
	return nil
}

func main() {
	arg1 := []int{2, 7, 11, 15}
	arg2 := 9
	result := twoSum(arg1, arg2)
	fmt.Println(result)
}
