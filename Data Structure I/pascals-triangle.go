/*
118. Pascal's Triangle
Given an integer numRows, return the first numRows of Pascal's triangle.
In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

    1
   1 1
  1 2 1
 1 3 3 1
1 4 6 4 1

2 = 1 + 1;
3 = 2 + 1;
4 = 3 + 1;
6 = 3 + 3;
and so on ...

Example:
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

Example:
Input: numRows = 1
Output: [[1]]

https://dev.to/alisabaj/solving-pascal-s-triangle-in-javascript-4e59

*/

package main

func generate(numRows int) [][]int {
	// Initialize an empty slice to hold the result
	result := make([][]int, numRows)

	// Loop through each row
	for i := 0; i < numRows; i++ {
		// Initialize a slice to hold the values for the current row
		row := make([]int, i+1)

		// Set the first and last value of the row to 1
		row[0] = 1
		row[i] = 1

		// Loop through the middle values of the row and calculate their value
		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
			/*
				For each element in the row, this line of code calculates its value
				by adding the values of the two elements directly above it
				in the previous row.
				The values are accessed from the result slice,
				which stores the previously generated rows of the triangle.

				For example,
				to calculate the second element of the third row (index 2),
				we add the values of the second and third elements
				from the second row (index 1 and 2 respectively):
				result[2-1][1-1] + result[2-1][1].
				This will give us the value of the second element in the third row.
			*/
		}

		// Add the current row to the result
		result[i] = row
	}

	// Return the result
	return result
}

func main() {

}
