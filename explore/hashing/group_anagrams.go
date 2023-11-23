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
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)

		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}


