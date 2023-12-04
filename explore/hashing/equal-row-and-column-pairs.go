/*
HASHING
Example 4: 2352. Equal Row and Column Pairs

Given an n x n matrix grid, return the number of pairs (R, C)
where R is a row and C is a column, and R and C are equal if
we consider them as 1D arrays.
*/

package main

import (
	"fmt"
	"strings"
)

func convertToKey(arr []int) string {
	var key strings.Builder

	for _, num := range arr {
		key.WriteString(fmt.Sprintf("%d,", num))
	}

	return key.String()
}

func equalPairs(grid [][]int) int {
	dic := make(map[string]int)

	// Count occurrences for each row
	for _, arr := range grid {
		key := convertToKey(arr)
		dic[key]++
	}

	dic2 := make(map[string]int)

	// Count occurrences for each column
	for col := 0; col < len(grid[0]); col++ {
		var currentCol []int
		for row := 0; row < len(grid); row++ {
			currentCol = append(currentCol, grid[row][col])
		}

		key := convertToKey(currentCol)
		dic2[key]++
	}

	ans := 0

	// Calculate the number of equal pairs
	for key, val := range dic {
		ans += val * dic2[key]
	}

	return ans
}
