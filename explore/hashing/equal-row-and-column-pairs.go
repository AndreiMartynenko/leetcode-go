/*
HASHING
Example 4: 2352. Equal Row and Column Pairs

Given an n x n matrix grid, return the number of pairs (R, C)
where R is a row and C is a column, and R and C are equal if
we consider them as 1D arrays.
*/

package main

func numEqualPairs(grid [][]int) int {
	n := len(grid)
	count := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				for l := 0; l < n; l++ {
					if grid[i][j] == grid[k][l] && i != k && j != l {
						count++
					}
				}
			}
		}
	}

}
