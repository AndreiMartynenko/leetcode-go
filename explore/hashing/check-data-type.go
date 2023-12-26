package main

import "fmt"

func checkDataType() {

	var count int = 42
	var message string = "go find type"
	var isCheck bool = true
	var amount float32 = 10.2
	var newSlice = []string{}
	newNewSlice := []int{1, 2, 3}

	fmt.Printf("variable count=%v is of type %T \n", count, count)
	fmt.Printf("variable message='%v' is of type %T \n", message, message)
	fmt.Printf("variable isCheck='%v' is of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount=%v is of type %T \n", amount, amount)
	fmt.Printf("variable newNewSlice =%v is of type %T \n", newNewSlice, newSlice)
}
