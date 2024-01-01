/*
HASHING

Counting

Example 5: 1248. Count Number of Nice Subarrays

Given an array of integers nums and an integer k.
A continuous subarray is called nice if there
are k odd numbers on it.

Return the number of nice sub-arrays.



Example 1:

Input: nums = [1,1,2,1,1], k = 3
Output: 2
Explanation: The only sub-arrays with 3 odd numbers are [1,1,2,1] and [1,2,1,1].

Example 2:

Input: nums = [2,4,6], k = 1
Output: 0
Explanation: There is no odd numbers in the array.

Example 3:

Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
Output: 16

*/

package main

func numberOfSubarrays(nums []int, k int) int {

	count := 0
	curr := 0
	oddCount := make(map[int]int)

	// Initialize the map with 0: 1, considering an empty prefix
	oddCount[0] = 1

	for _, num := range nums {
		// Update curr based on whether num is odd or even
		curr += num % 2

		// Check if (curr - k) exists in the map, if yes, increment the count
		if val, ok := oddCount[curr-k]; ok {
			count += val
		}

		// Increment the count of the current odd count in the map
		oddCount[curr]++
	}

	return count
}

/*
	func main() {
		nums := []int{1, 1, 2, 1, 1}
		k := 3

		result := countNiceSubarrays(nums, k)
		fmt.Println(result) // Output: 2
	}

*/
