package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x <= y {
		fmt.Printf("%d is less or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d greater is than %d\n", x, y)
	}

	color := "red"
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Color not is red or blue")
	}

	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color not is red or blue")
	}
}
