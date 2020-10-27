package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8181", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/")
	fmt.Fprint(w, "Index Page")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println("about")
	fmt.Fprint(w, "About Page")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("contact")
	fmt.Fprint(w, "Contact Page")
}
