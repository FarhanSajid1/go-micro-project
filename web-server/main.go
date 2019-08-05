package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

var templ *template.Template
var err error

// database attributes
var APIHOST = getEnv("APIHOST", "localhost")
var url = fmt.Sprintf("http://%s:8080/rpc", APIHOST)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func init() {
	templ = template.Must(template.ParseGlob("templates/*"))
}

type Post struct {
	Service string  `json:"service"`
	Method  string  `json:"method"`
	Request Request `json:"request"`
}
type Request struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Company  string `json:"company"`
	Password string `json:"password"`
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Printf("could not parse form %v", err)
		}
		name := r.FormValue("name")
		company := r.FormValue("company")
		password := r.FormValue("password")

		PostUser := &Post{
			Service: "go.micro.srv.user",
			Method:  "UserService.Create",
			Request: Request{
				Name:     name,
				Company:  company,
				Password: password,
			},
		}
		jsonValue, _ := json.Marshal(PostUser)
		log.Printf("url %v", url)
		response, err := http.Post(url,
			"application/json",
			bytes.NewBuffer(jsonValue),
		)
		// Reading from the binary respose into a string
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		resp := buf.String()
		templ.ExecuteTemplate(w, "index.html", resp)
	} else {
		templ.ExecuteTemplate(w, "index.html", nil)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Printf("could not parse form %v", err)
		}
		email := r.FormValue("email")
		password := r.FormValue("password")

		PostUser := &Post{
			Service: "go.micro.srv.user",
			Method:  "UserService.Auth",
			Request: Request{
				Email:    email,
				Password: password,
			},
		}
		jsonValue, _ := json.Marshal(PostUser)
		log.Print(url)
		response, err := http.Post(url,
			"application/json",
			bytes.NewBuffer(jsonValue),
		)
		// Reading from the binary respose into a string
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		resp := buf.String()
		templ.ExecuteTemplate(w, "login.html", resp)
	} else {
		templ.ExecuteTemplate(w, "login.html", nil)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login/", login)
	http.ListenAndServe(":8990", nil)
}
