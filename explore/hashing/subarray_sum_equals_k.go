/*
HASHING
Counting
Example 4: 560. Subarray Sum Equals K

Given an integer array nums and an integer k,
find the number of subarrays whose sum is equal to k.

Example 1:

Input: nums = [1,1,1], k = 2
Output: 2
Example 2:

Input: nums = [1,2,3], k = 3
Output: 2

*/

package main

func subarraySum(nums []int, k []int) int {

	count := 0
	sum := 0
	sumCount := make(map[int]int)

	// Initialize the map with the sum 0, as it represents an empty subarray
	sumCount[0] = 1

	/*
			In the given problem, we are looking for subarrays whose sum is equal to k.
			When calculating cumulative sums, if at any point the current sum is equal to k,
			it means we have found a subarray whose sum is directly equal to k.

		By setting sumCount[0] = 1 initially, we are essentially saying that
		we have seen the cumulative sum of 0 once (representing an empty subarray).
		This is necessary because, during the iteration through the array,
		if at any point the current sum (sum) equals k, we can increment the count
		by the value stored in sumCount[0], indicating that we have
		found a subarray with the desired sum.

		In summary, sumCount[0] = 1 handles the case where a subarray with sum k
		is encountered directly, and it ensures that such cases are correctly
		counted during the subsequent iterations.

	*/

	for _, num := range nums {
		sum += num

		// Check if (sum - k) exists in the map, if yes, increment the count
		if val, ok := sumCount[sum-k]; ok {
			count += val
		}

		// Increment the count of the current sum in the map
		sumCount[sum]++
	}

	return count

}

/*

func main() {
	nums := []int{1, 2, 3, 4, 1, 1, 1}
	k := 5

	result := subarraySum(nums, k)
	fmt.Println(result) // Output: 2
}

*/
