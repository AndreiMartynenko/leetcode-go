/*
Example 2: 2260. Minimum Consecutive Cards to Pick Up

Given an integer array cards, find the length of the
shortest subarray that contains at least one duplicate.
If the array has no duplicates, return -1.
*/

package main

func minimumCardPickup(cards []int) int {
	cardIndex := make(map[int]int)
	minLength := len(cards) + 1

}
