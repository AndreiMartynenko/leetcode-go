/*

HASHING

Jewels and Stones

You're given strings jewels representing the types of stones that are jewels,
and stones representing the stones you have. Each character in stones is a type
of stone you have. You want to know how many of the stones you have are also jewels.

Letters are case sensitive, so "a" is considered a different type of stone from "A".

Example 1:

Input: jewels = "aA", stones = "aAAbbbb"
Output: 3

Example 2:

Input: jewels = "z", stones = "ZZ"
Output: 0

Constraints:

1 <= jewels.length, stones.length <= 50
jewels and stones consist of only English letters.
All the characters of jewels are unique.

*/

package main

func numJewelsInStones(jewels string, stones string) int {
	jewelSet := make(map[rune]bool)

	// Create a set of jewels
	for _, jewel := range jewels {
		jewelSet[jewel] = true
	}

	count := 0

	// Count the stones that are also jewels
	for _, stone := range stones {
		if jewelSet[stone] {
			count++
		}
	}

	return count
}

/*
func main() {
	// Example 1
	jewels1 := "aA"
	stones1 := "aAAbbbb"
	result1 := numJewelsInStones(jewels1, stones1)
	fmt.Printf("Output for Example 1: %d\n", result1)

	// Example 2
	jewels2 := "z"
	stones2 := "ZZ"
	result2 := numJewelsInStones(jewels2, stones2)
	fmt.Printf("Output for Example 2: %d\n", result2)
}

*/
