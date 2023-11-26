/*
HASHING

Example 3: 2342. Max Sum of a Pair With Equal Sum of Digits

Given an array of integers nums, find the maximum value of nums[i] + nums[j],
where nums[i] and nums[j] have the same digit sum (the sum of their individual digits).
Return -1 if there is no pair of numbers with the same digit sum.
*/

package main

import "strconv"

func maximumSum(nums []int) int {

	digitSumMap := make(map[int][]int)
	maxSum := -1

	for i, num := range nums {

		digitSum := calculateDigitSum(num)

		if indices, found := digitSumMap[digitSum]; found {
			for _, index := range indices {

				currentSum := nums[index] + num
				if currentSum > maxSum {
					maxSum = currentSum
				}
			}
		}

		digitSumMap[digitSum] = append(digitSumMap[digitSum], i)
	}

	return maxSum
}

func calculateDigitSum(num int) int {
	sum := 0
	strNum := strconv.Itoa(num)

	for _, digit := range strNum {
		digitValue, _ := strconv.Atoi(string(digit))
		sum += digitValue
	}

	return sum
}

/*
func main() {
	nums := []int{18, 43, 36, 13, 7}
	result := maxSumOfPairWithEqualDigitSum(nums)

	fmt.Println(result)
}

*/
