package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func index(arg1 http.ResponseWriter, arg2 *http.Request) {
	fmt.Fprintln(arg1, "Welcome to my page, type the url to go to differnt routes.")
}

func dog(arg1 http.ResponseWriter, arg2 *http.Request) {
	fmt.Fprintln(arg1, "Hi, this is Sheru, the DOG")
}

func aboutMe(arg1 http.ResponseWriter, arg2 *http.Request) {
	tpl.ExecuteTemplate(arg1, "tpl.gohtml", "Durgesh Kumar")
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(aboutMe))
	//http.HandleFunc("/", index)
	//http.HandleFunc("/dog/", dog)
	//http.HandleFunc("/me/", aboutMe)
	http.ListenAndServe(":8080", nil)
}
