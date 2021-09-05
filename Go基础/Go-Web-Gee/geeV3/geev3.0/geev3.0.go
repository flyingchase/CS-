package geev3_0

import (
	"log"
	"net/http"
)

// define HandlerFunc from w r to context
type HandlerFunc func(c *Context)

// Engine实现 ServeHTTP 接口
type Engine struct {
	router *router
}

// engine 实现了 ServeHTTP即可认为Handler
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	// 调用router 字段的 handle 方法设置 handler
	engine.router.handle(c)
}

func New() *Engine {
	return &Engine{router: newRouter()}
}
// 使用 router 的 addRoute方法
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRoute(method,pattern,handler)

}
// 调用 addRoute方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
