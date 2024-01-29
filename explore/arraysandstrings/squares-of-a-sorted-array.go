/*
ARRAYS AND STRINGS
Two Pointers

Squares of a Sorted Array

Given an integer array nums
sorted in non-decreasing order,
return an array of the squares
of each number sorted in non-decreasing order.



Example 1:

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Example 2:

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]
*/

package main

func sortedSquares(nums []int) []int {
	// Square each number in the array
	squaredNums := make([]int, len(nums))
	for i, num := range nums {
		squaredNums[i] = num * num
	}

	// Sort the squared numbers using insertion sort
	for i := 1; i < len(squaredNums); i++ {
		key := squaredNums[i]
		j := i - 1
		for j >= 0 && squaredNums[j] > key {
			squaredNums[j+1] = squaredNums[j]
			j = j - 1
		}
		squaredNums[j+1] = key
	}

	return squaredNums
}

/*
func main() {
	newArr := []int{-4, -1, 0, 3, 10}
	result := sortedSquares(newArr)
	fmt.Println(result)
}
*/
