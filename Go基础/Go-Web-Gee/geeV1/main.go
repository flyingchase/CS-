package main

import (
	geev1_02 "Web-tutorial/geeV1/geev1.0"
	"fmt"
	"net/http"
)
/*
v1.0 功能
- 实现路由表映射
	engine 内有 Router 表（map[string]Handlerfunc）
		- key 是请求方法+addr 地址 v 是自定义的路由方法 Handler
- 提供用户注册的静态路由方法 Handler
- 包装启动函数 ListenAndServe 作为 Run



*/
func main() {
	r := geev1_02.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
	})

	r.GET("/hello",func(w http.ResponseWriter, r *http.Request){
		for k,v:=range r.Header {
		    fmt.Fprintf(w, "Header[%q] = %q\n",k,v)
		}
	})

	r.Run("localhost:8900")
}
