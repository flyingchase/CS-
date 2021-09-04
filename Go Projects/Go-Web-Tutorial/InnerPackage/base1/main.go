package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k,v:=range r.Header {
		fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "URL.path= ", r.URL.Path)
}
