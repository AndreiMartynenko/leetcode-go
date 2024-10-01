package main

import "fmt"

func searchElement(collection []int, target int) bool {
	hashTable := make(map[int]bool)
	for _, element := range collection {
		if hashTable[element] {
			return true // Element found
		} else {
			hashTable[element] = true
		}
	}
	return false // Element not found
}

func main() {
	myCollection := []int{1, 2, 3, 4, 5, 2, 6}
	targetElement := 4
	result := searchElement(myCollection, targetElement)
	fmt.Printf("Element %d is %s\n", targetElement, map[bool]string{true: "found", false: "not found"}[result])
}

// Testing for adding comments 
