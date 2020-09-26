package main

import "fmt"

func main() {
	// var name string = "Kolyan"
	// name := "Kolyan"
	// email := "kolya@kolya.ru"

	name, email := "Kolya", "kolya@kolya.ru"
	var age int32 = 30
	size := 1.3
	var isCool = true

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
	fmt.Printf("%b\n", age)
	fmt.Printf("%v\n", age)
}
