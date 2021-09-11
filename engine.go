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

func (e *engine) POST(path string, handlerFunc HandlerFunc) {
	e.r.add("POST", path, handlerFunc)
}

func (e *engine) GET(path string, handlerFunc HandlerFunc) {
	e.r.add("GET", path, handlerFunc)
}

func (e *engine) DELETE(path string, handlerFunc HandlerFunc) {
	e.r.add("DELETE", path, handlerFunc)
}

func (e *engine) PUT(path string, handlerFunc HandlerFunc) {
	e.r.add("PUT", path, handlerFunc)
}
