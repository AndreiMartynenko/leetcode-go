/*
HASHING

Counting Elements

Given an integer array arr, count how many
elements x there are, such that x + 1 is also
in arr. If there are duplicates in arr,
count them separately.



Example 1:

Input: arr = [1,2,3]
Output: 2
Explanation: 1 and 2 are counted cause 2
and 3 are in arr.
Example 2:

Input: arr = [1,1,3,3,5,5,7,7]
Output: 0
Explanation: No numbers are counted, cause
there is no 2, 4, 6, or 8 in arr.
*/

package main

func countElements(arr []int) int {
	count := 0
	elements := make(map[int]bool)

	for _, num := range arr {
		elements[num] = true
	}

	for _, num := range arr {
		if elements[num+1] {
			count++
		}
	}

	return count
}

/*
func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{1, 1, 3, 3, 5, 5, 7, 7}

	result1 := countElements(arr1)
	result2 := countElements(arr2)

	fmt.Printf("Number of elements in arr1: %d\n", result1)
	fmt.Printf("Number of elements in arr2: %d\n", result2)

}
*/
