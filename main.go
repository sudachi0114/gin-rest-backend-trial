package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("[path]", r.URL.Path)
	fmt.Println("[scheme]", r.URL.Scheme)
	fmt.Println("[url_long]", r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("  [key]", k)
		fmt.Println("  [val]", v)
	}

	fmt.Fprintf(w, "Hello world")
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
