package main

import (
	geev2_0 "Web-tutorial/geeV2/geev2.0"
	"net/http"
)
/*
gee v2.0
- 封装 router 通过实现 ServeHTTP 接口接管所有 HTTP
- 调用 router.handle之前构造 context 对象
- context对象包装了两个参数 http.ResponseWriter *http.Request
*/
func main() {
	r := geev2_0.New()
	r.GET("/", func(c *geev2_0.Context) {
		c.HTML(http.StatusOK, "<h1>hello world</h1>")
	})
	r.GET("/hello", func(c *geev2_0.Context) {
		c.String(http.StatusOK, "hello %s, where you are %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *geev2_0.Context) {
		c.JSON(http.StatusOK, geev2_0.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":8900")
}
