package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read val from address

	fmt.Println(*b)
	fmt.Println(*&a)

	fmt.Println(f())
	//change val with pointer
	*b = 10
	fmt.Println(a)

	incrVar := 11
	fmt.Println(incr(&incrVar))
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
