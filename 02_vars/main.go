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
	var i, j, k int
	fmt.Printf("%b %b %b\n", i, j, k)
	var b, f, s = true, 43.4, "vasya"
	fmt.Printf("%v %f %s\n", b, f, s)

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
	fmt.Printf("%b\n", age)
	fmt.Printf("%v\n", age)

	p := new(int) //p type of *int, pointer to an unnamed int var
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	str := "hello world!"
	fmt.Println(len(s))
	fmt.Println(str[0], str[10])

	fmt.Println("goodbye" + str[5:])

	s1 := "first"
	s2 := s1
	s1 += ", second"

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println("\a")

	fmt.Println(HasPrefix("teststring", "te"))
	fmt.Println(HasSuffix("teststring", "ing"))

}

//HasPrefix returns true if s contains prefix pref
func HasPrefix(s, pref string) bool {
	return len(s) >= len(pref) && s[:len(pref)] == pref
}

//HasSuffix returns true if s contains suffix pref
func HasSuffix(s, suf string) bool {
	return len(s) >= len(suf) && s[len(s)-len(suf):] == suf
}

//Contains returns true if s contains substr
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s, substr) {
			return true
		}
	}
	return false
}
