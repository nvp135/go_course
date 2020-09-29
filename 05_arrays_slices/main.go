package main

import "fmt"

func main() {
	//Arrays
	var fruitArray [2]string

	//Assign values
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	//Declare and assign
	fruitArray1 := [2]string{"Banana", "Potato"}

	fmt.Println(fruitArray)
	fmt.Println(fruitArray1[1])

	fruitSlice := []string{"Apple", "Banana", "Orange", "Grape"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
