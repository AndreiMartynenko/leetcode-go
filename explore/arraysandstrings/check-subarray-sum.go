package main

/*
ARRAYS AND STRINGS

Prefix Sum

Example 1: Given an integer array nums, an array queries
where queries[i] = [x, y] and an integer limit,
return a boolean array that represents the answer to each query.
A query is true if the sum of the subarray from x to y
is less than limit, or false otherwise.

For example, given nums = [1, 6, 3, 2, 7, 2],
queries = [[0, 3], [2, 5], [2, 4]], and limit = 13,
the answer is [true, false, true]. For each query,
the subarray sums are [12, 14, 12].


To solve this problem using the prefix sum technique,
you can follow these steps:

1. Calculate the prefix sum array for the given nums.
2. For each query [x, y], calculate the sum of the subarray
as prefix_sum[y] - prefix_sum[x-1] (considering prefix_sum[-1] as 0).
3. Compare the sum obtained in step 2 with the given limit
to determine whether it's less than limit.
*/

import "fmt"

func checkSubarraySum(nums []int, queries [][]int, limit int) []bool {
	n := len(nums)

	// Step 1: Calculate prefix sum
	prefixSum := make([]int, n)
	prefixSum[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	// Step 2 and 3: Check each query
	var result []bool
	for _, query := range queries {
		x, y := query[0], query[1]

		//Handle the case where the x == 0 seperatly
		// subarraySum := prefixSum[y] - prefixSum[x-1]
		subarraySum := prefixSum[y]
		if x > 0 {
			subarraySum -= prefixSum[x-1]
		}
		result = append(result, subarraySum < limit)
	}

	return result
}

func main() {
	nums := []int{1, 6, 3, 2, 7, 2}
	queries := [][]int{{0, 3}, {2, 5}, {2, 4}}
	limit := 13

	result := checkSubarraySum(nums, queries, limit)
	fmt.Println(result) // Output: [true false true]
}
