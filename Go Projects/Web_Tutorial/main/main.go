package main

import (
	"net/http"
	"text/template"
)

func process(w http.ResponseWriter, r *http.Request)  {
	t,_:=template.ParseFiles("tmpl.html")
	t.Execute(w,"he")
}
func main() {
	server:=http.Server {
	    Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
//func process(w http.ResponseWriter, r *http.Request) {
//	r.ParseMultipartForm(1024)                      // 最大传递字节
//	fileHead := r.MultipartForm.File["uploaded"][0] // 读取指定 body 内容
//	file, err := fileHead.Open()
//	if err == nil {
//		data, err := ioutil.ReadAll(file)
//		if err == nil {
//			fmt.Fprintln(w, string(data))
//		}
//	}
//}
//func writeExample(w http.ResponseWriter, r *http.Request) {
//	str := `<html>
//<head><title>Go Web</title></head>
//<body><h1>Hello World</h1></body>
//</html>`
//	w.Write([]byte(str))
//}
//func writeHeaderExampl(w http.ResponseWriter, r *http.Request) {
//	w.WriteHeader(501)
//	fmt.Fprintln(w, "No such service, try next door")
//}
//func main() {
//	server := http.Server{
//		Addr: "localhost:8080",
//	}
//	http.HandleFunc("/write", writeHeaderExampl)
//	http.HandleFunc("/redirect", headerEXample)
//	http.HandleFunc("/json", jsonExample)
//	server.ListenAndServe()
//}
//func headerEXample(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Location", "http://google.com")
//	w.WriteHeader(302)
//
//}
//
//type Post struct {
//	User    string
//	Threads []string
//}
//
//func jsonExample(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	post := &Post{
//		User : "wlzhou",
//		Threads: []string {"first","second","third"},
//	}
//	json,_:=json.Marshal(post)
//	w.Write(json)
//}
