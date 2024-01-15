/*
ARRAYS AND STRINGS
Prefix Sum

2270. Number of Ways to Split Array

Given an integer array nums, find the number
of ways to split the array into two parts so
that the first section has a sum greater than or
equal to the sum of the second section.
The second section should have at least one number.

*/

package main

func waysToSplitArray(nums []int) int {
	n := len(nums)           // Get the length of the input array 'nums'
	prefix := make([]int, n) // Create a slice 'prefix' to store prefix sums with the same length as 'nums'

	prefix[0] = nums[0] // The first element of the prefix sum array is the same as the first element of 'nums'

	// Iterate through the elements of 'nums' starting from the second element
	for i := 1; i < n; i++ {
		// Calculate the current prefix sum by adding the previous prefix sum with the current element of 'nums'
		prefix[i] = prefix[i-1] + nums[i]
	}

	ans := 0 // Initialize variable 'ans' to keep track of the number of ways

	// Iterate through the elements of the prefix sum, except the last one
	for i := 0; i < n-1; i++ {
		leftSection := prefix[i]                // Left section of the split: sum of elements from the start to i
		rightSection := prefix[n-1] - prefix[i] // Right section of the split: sum of elements from i+1 to end

		// If the sum of the left section is greater than or equal to the sum of the right section, increment 'ans'
		if leftSection >= rightSection {
			ans++
		}
	}

	return ans // Return the total number of ways

}

/*
func main() {
	newArr := []int{10, 4, -8, 7}
	result := waysToSplitArray(newArr)
	fmt.Println(result)
}
*/

// 2
