package main

import "fmt"

func main() {

	hashMap := map[int]int{
		1: 2,
		5: 3,
		7: 2,
	}

	//hashMap := make(map[int]int)

	//check if the key exists
	value, ok := hashMap[1]
	_, ok := hashMap[9]

	//Accessing the value of a key
	value := hashMap[5]

	//Adding or updating a new key-value pair:
	hashMap[5] = 6

	//If the key does not exist, it will be added to the map.
	hashMap[9] = 15

	//Deleting a key-value pair:
	delete(hashMap, 9)

	//Get size of the map
	size := len(hashMap)

	//Iterating over a map
	for key, value := range hashMap {
		fmt.Println(key, value)
	}
}

/*
### Example 1:

You have a string `s` and an integer `k`. Find the length of the longest substring that contains no more than `k` distinct characters.

For example, if `s = "eceba"` and `k = 2`, the answer is 3. The longest substring with no more than 2 distinct characters is `"ece"`.

*/

//Check if the key exists in the map
//value, ok := hashMap[key]
//if ok {
//	fmt.Println(value)

//func lengthOfLongestSubstring(s string, k int) {
//	hashMap := make(map[rune]int)
//	left := 0
//	maxLength := 0
//
//	for right, char := range s {
//		hashMap[char]++
//		for len(hashMap) > k {
//			hashMap[s[left]]--
//			if hashMap[s[left]] == 0 {
//				delete(hashMap, s[left])
//			}
//			left++
//		}
//		maxLength = max(maxLength, right-left+1)
//	}
//	return maxLength
//
//}
