package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(arg1 http.ResponseWriter, arg2 *http.Request) {
		fmt.Fprintln(arg1, "Welcome to my page, type the url to go to differnt routes.")
	})
	http.HandleFunc("/dog/", func(arg1 http.ResponseWriter, arg2 *http.Request) {
		fmt.Fprintln(arg1, "Hi, this is Sheru, the DOG")
	})
	http.HandleFunc("/me/", func(arg1 http.ResponseWriter, arg2 *http.Request) {
		fmt.Fprintln(arg1, "Hi, this is Durgesh here.")
	})
	http.ListenAndServe(":8080", nil)
}
