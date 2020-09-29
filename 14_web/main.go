package main

import (
	"fmt"
	"net/http"
)

// Person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

func index(w http.ResponseWriter, r *http.Request) {
	//person1 := Person{firstName: "Vasiliy", lastName: "Terkin", city: "Ekaterinbutg", gender: "m", age: 10}
	fmt.Fprintf(w, "<h1>Hello world</h1>") //+person1)
}

func about(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "<h1>About</h1>")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(1488)
	fmt.Fprintln(w, "my error")

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server starting ...")
	http.ListenAndServe(":3000", nil)
}
