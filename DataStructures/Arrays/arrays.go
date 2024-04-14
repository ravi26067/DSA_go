package main

import (
	"fmt"
	"reflect"
)

// This allows the function to work with the array dynamically, regardless of its type
func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	//This line uses the reflect package to get a reflect.Value object representing the value of the arrayType parameter. This allows the function to work with the array dynamically, regardless of its type
	if arr.Kind() != reflect.Array {
		panic("Invalid datatype")
	}

	for i := 0; i < arr.Len(); i++ {
		//This line compares the element at index i in the array (converted to an interface{} type using Interface()) with the item parameter.
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func main() {
	sports := [...]string{"football", "basketball", "hockey"}

	fmt.Println("Looping through the array")
	for i := 0; i < len(sports); i++ {
		fmt.Println(sports[i])
	}

	fmt.Println("The sport football exist? (true or false)")
	//Searching for a particular item
	fmt.Println(itemExists(sports, "football"))
	fmt.Println("Dummy test")
}
