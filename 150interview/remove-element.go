package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := len(nums) - 1 // Last index of valid elements in nums
	j := 0             // Pointer for current element to check

	for j <= i {
		if nums[j] == val {
			// Swap with the last valid element
			nums[j] = nums[i]
			i-- // Decrease the size of valid elements
		} else {
			j++ // Move to the next element
		}
	}
	return i + 1 // Return the new length of the array
}

func main() {
	// Example usage
	nums := []int{3, 2, 2, 3}
	val := 3
	newLength := removeElement(nums, val)
	// nums should now be [2, 2] and newLength should be 2
	_ = newLength                         // Use newLength as needed
	fmt.Println(nums[:newLength])         // Output the modified array
	fmt.Println("New length:", newLength) // Output the new length
}
