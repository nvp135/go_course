package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Appointment struct
type Appointment struct {
	Date    string `json:appointmentDate`
	StartMn string `json:appointmentStartMn`
	ID      string `json:appointmentId`
	UserId  string `json:appointmentUserId`
}

func main() {
	fmt.Println("Server is starting ...")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/script.js", jsHandler)
	http.HandleFunc("/style.css", cssHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/cookies", cookiesHandler)
	http.HandleFunc("/cookies/script.js", jsHandler)
	http.HandleFunc("/cookies/style.css", cssHandler)
	http.HandleFunc("/register", registryHandler)
	http.HandleFunc("/register/style.css", cssHandler)
	http.HandleFunc("/testcookies", testcookiesHandler)

	fmt.Println("Server is working ...")

	http.ListenAndServe(":80", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println(err)
	}

	// for _, cookie := range r.Cookies() {
	// 	cookie.Name
	// }

	dtGenerate, cookieErr := r.Cookie("dtGenerate")

	if cookieErr != nil {
		fmt.Printf("Cookie not found: '%v' and will be generated.", cookieErr)
		dtGenerate = generateCookie()
		http.SetCookie(w, dtGenerate)
	}

	person := Person{firstName: "Vasiliy", lastName: "Terkin", city: "Ekaterinbutg", gender: "m", age: 10}

	result := string(content)

	body := fmt.Sprintf("<h1>Hello %v %v </h1></br> Cookie generated: %v", person.firstName, person.lastName, dtGenerate.Value)

	fmt.Fprintf(w, result, body)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	fmt.Fprintln(w, "my error")
}

func cookiesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	result, _ := ioutil.ReadFile("./cookies/cookies.html")
	fmt.Fprintln(w, string(result))
}

func testcookiesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	cookie, err := r.Cookie("appointment")
	if err != nil {
		fmt.Println("Cookie wasn't found")
		return
	}

	data, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	var appointment Appointment

	err = json.Unmarshal(data, &appointment)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("%v, %v, %v, %v\r\n", appointment.Date, appointment.StartMn, appointment.ID, appointment.UserId)

	result, _ := ioutil.ReadFile("./cookies/testcookies.html")
	fmt.Fprintln(w, string(result))
}

func registryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	html, _ := ioutil.ReadFile("./register/register.html")
	resultHtml := string(html)
	if r.Method == "GET" {
		resultHtml = fmt.Sprintf(resultHtml,
			`<input name="name" type="text" value="" />`,
			`<input name="address" type="text" value="" />`,
			`<br><input type="submit" value="Подтвердить" />`)
		fmt.Fprintln(w, resultHtml)
		return
	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		address := r.FormValue("address")
		resultHtml = fmt.Sprintf(resultHtml,
			`<label>`+name+`</label>`,
			`<label>`+address+`</label>`,
			``)
		fmt.Fprintln(w, resultHtml)
	}
}

func generateCookie() *http.Cookie {
	t := time.Now()
	currentDt := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	cookie := http.Cookie{Name: "dtGenerate", Value: currentDt, MaxAge: 60}
	return &cookie
}

func jsHandler(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI[1:]
	jsFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "text/javascript")
	fmt.Fprintln(w, string(jsFile))
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI[1:]
	cssFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "text/css")
	fmt.Fprintln(w, string(cssFile))
}
