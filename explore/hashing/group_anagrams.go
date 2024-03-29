/*
HASHING

Example 1: 49. Group Anagrams

Given an array of strings strs,
group the anagrams together.

For example, given strs = ["eat","tea","tan","ate","nat","bat"],
return [["bat"],["nat","tan"],["ate","eat","tea"]].
*/

package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)

		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	var result [][]string
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	characters := strings.Split(s, "")

	sort.Strings(characters)

	return strings.Join(characters, "")
}

/*
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)

	fmt.Println(result)
}

*/
