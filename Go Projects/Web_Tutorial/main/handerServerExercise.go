package main

import "net/http"
//
//func main() {
//
//	mh := myHandler{}
//
//	server := http.Server{
//		Addr: "localhost:1090",
//		Handler: nil,	// use DefaultServeMux
//	}
//	http.Handle("/wo",&mh)
//	http.Handle("/home",http.HandlerFunc(welcome))
//	server.ListenAndServe()
//}

type myHandler struct {
}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello"))
}
func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home!"))
}