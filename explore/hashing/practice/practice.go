package main

import "fmt"

//hashMap := make(map[int]int)

//Iterating over a map

func main() {

	hashMap := map[int]int{
		1: 2,
		5: 3,
		7: 2,
	}

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

	for key, value := range hashMap {
		fmt.Println(key, value)
	}
}
