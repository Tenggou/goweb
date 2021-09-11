package goweb

import (
	"fmt"
	"net/http"
)

func NewEngine() *engine{
	return &engine{
		router{
			m: make(map[string]http.HandlerFunc),
		},
	}
}

type engine struct {
	r router
}

func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path
	handler, err := e.r.get(method, path)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	handler(w, req)
}

func (e *engine)POST(path string, handlerFunc http.HandlerFunc)  {
	e.r.add("POST", path, handlerFunc)
}

func (e *engine)GET(path string, handlerFunc http.HandlerFunc)  {
	e.r.add("GET", path, handlerFunc)
}

func (e *engine)DELETE(path string, handlerFunc http.HandlerFunc)  {
	e.r.add("DELETE", path, handlerFunc)
}

func (e *engine)PUT(path string, handlerFunc http.HandlerFunc)  {
	e.r.add("PUT", path, handlerFunc)
}


func (e *engine)Run(addr string) error {
	return http.ListenAndServe(addr, e)
}