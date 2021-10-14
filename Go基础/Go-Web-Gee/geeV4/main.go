package main

import (
	geev4_0 "Web-tutorial/geeV4/geev4.0"
	"net/http"
)

func main() {
	r := geev4_0.New()
	r.GET("index", func(c *geev4_0.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *geev4_0.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee1</h1>")
		})

		v1.GET("/hello", func(c *geev4_0.Context) {

			c.String(http.StatusOK, "hello %s, where you are at %s\n", c.Query("name"), c.Path)

		})

	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *geev4_0.Context) {
			c.String(http.StatusOK, "hello %s, you are at %s\n", c.Param("name"), c.Path)

		})

		v2.POST("/login", func(c *geev4_0.Context) {

			c.JSON(http.StatusOK, geev4_0.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	r.Run(":9999")
}
