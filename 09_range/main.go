package main

import "fmt"

func main() {
	ids := []int{22, 24, 51, 11, 66, 311}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "Sharon@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
