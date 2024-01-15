/*
ARRAYS AND STRINGS
Two Pointers

Example 2: Given a sorted array of unique integers
and a target integer, return true if there exists
a pair of numbers that sum to target, false otherwise.
This problem is similar to Two Sum.
(In Two Sum, the input is not sorted).

For example, given nums = [1, 2, 4, 6, 8, 9, 14, 15]
and target = 13, return true because 4 + 9 = 13.


*/

package main

func sumToTarget(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	for left <= right {
		current := (nums[left] + nums[right])
		if current == target {
			return true
		}
		if current > target {
			right--
		} else {
			left++
		}

	}
	return false
}

/*
func main() {
	newArr := []int{1, 2, 4, 6, 8, 9, 14, 15}
	result := sumToTarget(newArr, 13)
	fmt.Println(result)
}
*/
