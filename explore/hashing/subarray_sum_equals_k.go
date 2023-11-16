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

}
