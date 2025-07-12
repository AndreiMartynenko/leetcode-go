package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// Initialize a pointer for the position of the last unique element
	lastUniqueIndex := 0
	// Iterate through the array starting from the second element
	for i := 1; i < len(nums); i++ {
		// If the current element is different from the last unique element
		if nums[i] != nums[lastUniqueIndex] {
			// Move the last unique index forward
			lastUniqueIndex++
			// Update the next unique element in the array
			nums[lastUniqueIndex] = nums[i]
		}
	}
	// Return the length of the array with unique elements
	return lastUniqueIndex + 1
}

func main() {
	// Example usage
	nums := []int{1, 1, 2, 2, 3, 4, 4}
	length := removeDuplicates(nums)
	// The first 'length' elements of nums will be unique
	// nums[:length] should now be [1, 2, 3, 4]
	// length should be 4
	_ = length                 // Use length as needed
	_ = nums[:length]          // Use the modified array as needed
	fmt.Println(nums[:length]) // Uncomment to see the result
}
