/*
 * @Author: wlzhou
 * @Date:   2021-09-03 11:27:40
 * @Last Modified time: 2021-09-03 11:32:22
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)
func sayHelloName (w http.ResponseWriter,r *http.Request) {
	r.ParseForm()


	fmt.Println(r.Form)
	fmt.Println("path:",r.URL.Path)

	fmt.Println("scheme: ",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k,v:=range r.Form {
		fmt.Println("key: ",k,"value: ",v)
	}

	fmt.Fprintf(w,"hello ")
}

func main() {
	http.HandleFunc("/",sayHelloName)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil {
		log.Fatal("listenAndServer: ",err)
	} 
		
	
}