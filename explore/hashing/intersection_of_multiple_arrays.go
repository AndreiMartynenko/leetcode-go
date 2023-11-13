/*
HASHING
Counting
Example 2: 2248. Intersection of Multiple Arrays

Given a 2D array nums that contains n
arrays of distinct integers,
return a sorted array containing all
the numbers that appear in all n arrays.

For example, given nums = [[3,1,2,4,5],[1,2,3,4],[3,4,5,6]],
return [3, 4]. 3 and 4 are the only numbers that are in all arrays.
*/

package main

import (
	"sort"
)

func intersect(nums [][]int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	// Initialize the result with the numbers from the first array
	result := make(map[int]int)
	for _, num := range nums[0] {
		result[num]++
	}

	// Iterate through the rest of the arrays
	for i := 1; i < len(nums); i++ {
		currentMap := make(map[int]int)

		// Count the frequency of each number in the current array
		for _, num := range nums[i] {
			currentMap[num]++
		}
		// Update the result map to only include numbers present in both arrays
		for num, count := range result {
			if currentCount, ok := currentMap[num]; ok {
				// Choose the minimum count to account for duplicates
				result[num] = min(count, currentCount)
			} else {
				// If the number is not present in the current array, remove it from the result
				delete(result, num)
			}
		}
	}

	// Convert the result map to a sorted array
	var intersection []int
	for num := range result {
		intersection = append(intersection, num)
	}

	// Sort the array before returning
	sort.Ints(intersection)
	return intersection

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
func main() {
	nums := [][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}}
	result := intersect(nums)
	fmt.Println(result) // Output: [3 4]
}

*/
