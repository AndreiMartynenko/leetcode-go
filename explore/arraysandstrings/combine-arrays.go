/*
ARRAYS AND STRINGS
Two Pointers

Example 3: Given two sorted integer arrays arr1 and arr2,
return a new array that combines both of them and is also sorted.
*/

package main

func combineArray(arr1 []int, arr2 []int) []int {
	answer := []int{}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			answer = append(answer, arr1[i])
			i++
		} else {
			answer = append(answer, arr2[j])
			j++
		}
	}
	/*
		for i < len(arr1) {
			answer = append(answer, arr1[i])
			i++
		}
		for j < len(arr2) {
			answer = append(answer, arr2[j])
			j++
		}
	*/
	for _, num := range arr1[i:] {
		answer = append(answer, num)
		
	}

	for _, num := range arr2[j:] {
		answer = append(answer, num)
	}
	return answer
}

/*
func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	result := combineArray(arr1, arr2)
	fmt.Println(result)
}
*/
