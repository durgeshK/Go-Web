package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(arg1 http.ResponseWriter, arg2 *http.Request) {
		fmt.Fprintln(arg1, "Foo Controller")
	})
	http.HandleFunc("/dog/", func(arg1 http.ResponseWriter, arg2 *http.Request) {
		tpl, err := template.ParseFiles("dog.gohtml")
		if err != nil {
			http.Error(arg1, "File not found", 404)
		}
		tpl.Execute(arg1, nil)
	})
	http.HandleFunc("/dog.jpg", func(arg1 http.ResponseWriter, arg2 *http.Request) {
		http.ServeFile(arg1, arg2, "dog.jpg")
	})
	http.ListenAndServe(":8080", nil)
}
