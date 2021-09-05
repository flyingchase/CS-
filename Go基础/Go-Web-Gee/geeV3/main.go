package main

import (
	geev3_0 "Web-tutorial/geeV3/geev3.0"
	"net/http"
)

func main() {
	r:=geev3_0.New()
	r.GET("/", func(c *geev3_0.Context) {
		c.HTML(http.StatusOK, "<h1>I am Fine!\n</h1>")
	})
	r.GET("/hello", func(c *geev3_0.Context) {
		c.String(http.StatusOK,"hello %s, the path is :%s\n",c.Query("name"),c.Path)
	})

	r.GET("/hello/:name", func(c *geev3_0.Context) {
		c.String(http.StatusOK,"hello %s,you are at %s\n",c.Param("name"),c.Path)
	})

	r.GET("/assets/*filepath", func(c *geev3_0.Context) {
		c.JSON(http.StatusOK,geev3_0.H{"filepath":c.Param("filepath")})
	})

	r.Run(":9999")
}
