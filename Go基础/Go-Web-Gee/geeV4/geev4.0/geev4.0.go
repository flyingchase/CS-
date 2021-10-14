package geev4_0

import (
	"log"
	"net/http"
)

type HandlerFunc func(c *Context)

// Engine implenent ServeHTTP
type (
	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
	RouterGroup struct {
		prefix      string
		middlewares []HandlerFunc
		parent      *RouterGroup
		engine      *Engine
	}
)

func (e *Engine)ServeHTTP(w http.ResponseWriter, r * http.Request)  {
	c:=newContext(w,r)
	e.router.handle(c)
}
func New() *Engine {
	engine:=&Engine {
	    router: newRouter(),
	}
	engine.RouterGroup=&RouterGroup {
	    engine : engine,
	}
	engine.groups=[]*RouterGroup {engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine:=group.engine
	newGroup := & RouterGroup{
		prefix : prefix,
		parent: group,
		engine : engine,
	}
	engine.groups=append(engine.groups, newGroup)
	return newGroup
	
	
}
func (group *RouterGroup) addRoute(method string,comp string,handler HandlerFunc)  {
	pattern :=group.prefix+comp
	log.Printf("Route %4s  -  %s",method,pattern)
	group.engine.router.addRoute(method,pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc)  {
	group.addRoute("GET",pattern, handler)

}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST",pattern, handler)
}

func (e *Engine) Run(addr string)(err error)  {
	return http.ListenAndServe(addr, e)
}




































