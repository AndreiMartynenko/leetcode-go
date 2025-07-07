package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1     // Last index of valid elements in nums1
	j := n - 1     // Last index of nums2
	k := m + n - 1 // Last index of nums1 (including extra space)

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

// func main() {
// 	// Example usage
// 	nums1 := []int{1, 2, 3, 0, 0, 0}
// 	m := 3
// 	nums2 := []int{2, 5, 6}
// 	n := 3
// 	merge(nums1, m, nums2, n)
// 	// nums1 should now be [1, 2, 2, 3, 5, 6]
// }
