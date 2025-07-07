package main

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
