package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func main() {
	x := "helloasdasdasdsadsad!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
}

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error os.Getwd: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}
