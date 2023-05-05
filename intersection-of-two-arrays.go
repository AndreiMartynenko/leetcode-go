package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	var arr []int
	for i := 0; i < len(nums2); i++ {
		num := nums2[i]
		for j := 0; j < len(nums1); j++ {
			if nums1[j] == num {
				arr = append(arr, num)
				nums1 = append(nums1[:j], nums1[j+1:]...)
				break
			}
		}
	}
	return arr
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{4, 5, 6, 7, 8}
	result := intersect(arr1, arr2)
	fmt.Println(result)
}
