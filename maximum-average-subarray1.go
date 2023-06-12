package main

import (
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	currSum := 0
	maxSum := math.MinInt64

	for i := 0; i < len(nums); i++ {
		currSum += nums[i]

		if i >= k-1 {
			maxSum = max(currSum, maxSum)
			currSum -= nums[i-(k-1)]
		}
	}

	return float64(maxSum) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	nums := []int{1, 12, -5, -6, 50, 3}
// 	k := 4
// 	result := findMaxAverage(nums, k)
// 	fmt.Println(result) // Output: 12.75
// }
