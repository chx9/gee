package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]HandleFunc
	roots    map[string]*node
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandleFunc),
		roots:    make(map[string]*node),
	}
}
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

func (r *router) AddRouter(method string, pattern string, handler HandleFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {

	key := c.Method + "-" + c.Path
	log.Printf("%s\n", key)
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND:%s \n", c.Path)
	}
}
