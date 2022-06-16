package gee

import (
	"net/http"
)

type HandleFunc func(*Context)

type Engine struct {
	// router map[string]HandleFunc
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) AddRouter(method string, pattern string, handler HandleFunc) {
	engine.router.AddRouter(method, pattern, handler)
}
func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.AddRouter("GET", pattern, handler)
}
func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.AddRouter("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	e.router.handle(c)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
