package geev2_0

import (
	"encoding/json"
	"fmt"
	"net/http"
)
/*
- 提供 Query 和 PostFOrm 参数的方法
- 快速构造 String/Data/JSON/HTML 响应方法
*/


// H 为 JSON 数据
type H map[string]interface{}

type Context struct {
	// 初始参数 作为 handle
	Writer http.ResponseWriter
	Req    *http.Request

	Path   string
	Method string

	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// 使用 http.Request的 FormValue 查询 Key 对应的第一个 Value （）
// 表单在前 url 在后
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, value ...interface{}) {
	c.SetHeader("Content-type", "text/plain")
	c.Status(code)
	// 类似 format——> %c char
	c.Writer.Write([]byte(fmt.Sprintf(format,value ...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type","application/json")
	c.Status(code)
	encoder:=json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj);err != nil {
		http.Error(c.Writer,err.Error(),500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}
func (c *Context)HTML(code int, html string)  {
	c.SetHeader("Content-Type","application/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
