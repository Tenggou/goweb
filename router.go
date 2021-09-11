package goweb

import (
	"fmt"
)

/*
保存path到handle function的映射
1. 注册服务，RESTful
2. 服务获取
*/

type HandlerFunc func(*Context)

func keyBuilder(method string, path string) string {
	return method + "-" + path
}

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{make(map[string]HandlerFunc)}
}

func (r *router) add(method string, path string, fun HandlerFunc) {
	r.handlers[keyBuilder(method, path)] = fun
}

func (r *router) get(method string, path string) (HandlerFunc, error) {
	p := keyBuilder(method, path)
	if _, ok := r.handlers[p]; !ok {
		return nil, fmt.Errorf("%v is not existed", p)
	}

	return r.handlers[p], nil
}
