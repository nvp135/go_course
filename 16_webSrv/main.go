package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

func generateCookie() *http.Cookie {
	t := time.Now()
	currentDt := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	cookie := http.Cookie{Name: "dt", Value: currentDt}
	return &cookie
}

func index(w http.ResponseWriter, r *http.Request) {
	rURI := r.RequestURI
	fmt.Println(rURI)
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}

	person := Person{firstName: "Vasiliy", lastName: "Terkin", city: "Ekaterinbutg", gender: "m", age: 10}

	result := string(content)

	body := fmt.Sprintf("<h1>Hello %v %v </h1>", person.firstName, person.lastName)

	http.SetCookie(w, generateCookie())

	fmt.Fprintf(w, result, body)
}

func about(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "<h1>About</h1>")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(1488)
	fmt.Fprintln(w, "my error")

}

func main() {
	fmt.Println("Server is starting ...")
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server is working ...")
	http.ListenAndServe(":3000", nil)

}
