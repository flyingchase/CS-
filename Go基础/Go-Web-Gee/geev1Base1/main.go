/*
 * @Author: wlzhou
 * @Date:   2021-09-04 18:20:04
 * @Last Modified time: 2021-09-04 18:25:31
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/hello",helloHandler)

	log.Fatal(http.ListenAndServe("localhost:8080",nil))
	
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)	
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k,v:=range r.Header {
		fmt.Fprintf(w,"Header[%q] = %q\n",k,v)
	}
}
