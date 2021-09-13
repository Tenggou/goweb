package goweb

import (
	"fmt"
	"net/http"
)

func NewEngine() *engine {
	return &engine{
		newRouter(),
	}
}

type engine struct {
	r *router
}

func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(w, req)
	handler, err := e.r.get(ctx.Method, ctx.Path)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		fmt.Fprintf(ctx.Writer, err.Error())
		return
	}

	handler(ctx)
}

func (e *engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *engine) POST(pattern string, handlerFunc HandlerFunc) {
	e.r.add("POST", pattern, handlerFunc)
}

func (e *engine) GET(pattern string, handlerFunc HandlerFunc) {
	e.r.add("GET", pattern, handlerFunc)
}

func (e *engine) DELETE(pattern string, handlerFunc HandlerFunc) {
	e.r.add("DELETE", pattern, handlerFunc)
}

func (e *engine) PUT(pattern string, handlerFunc HandlerFunc) {
	e.r.add("PUT", pattern, handlerFunc)
}
