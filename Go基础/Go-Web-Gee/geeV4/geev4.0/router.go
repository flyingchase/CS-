package geev4_0

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node),
	}
}

// 将 pattern 转化为 string 分割为 parts
// 考虑 通配符*的情况
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)

	for _, item := range vs {
		if item != "" {

			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	// make new nodes if key method doesn't exit
	if !ok {
		r.roots[method] = &node{}
	}
	// add New route
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchPaths := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchPaths, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			// :时候丢掉第一个
			if part[0] == ':' {
				params[part[1:]] = searchPaths[index]
			}
			// 通配符*时候保证非第一个
			if part[0] == '*' && len(part) > 1 {
				params[part[1 : ]]=strings.Join(searchPaths[index : ],"/")
				break
			}
		}
		return n ,params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method,c.Path)
	if n != nil {
	    c.Params=params
	    key:=c.Method+"-"+n.pattern
	    r.handlers[key](c)
	}else {
		c.String(http.StatusNotFound,"404 at %s\n",c.Path)
	}
}
