package goweb

import (
	"fmt"
	"strings"
)

/*
保存pattern到handle function的映射
1. 注册服务，RESTful
2. 服务获取
*/

type HandlerFunc func(*Context)

type router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*Node
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*Node), // todo Node起始为空，在get和add里面需要特判
	}
}

func (r *router) add(method string, pattern string, fun HandlerFunc) {
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = newNode()
	}

	r.roots[method].insert(patternParser(pattern), pattern, 0)
	r.handlers[keyBuilder(method, pattern)] = fun
}

func (r *router) get(method string, pattern string) (HandlerFunc, error) {
	if _, ok := r.roots[method]; !ok {
		return nil, fmt.Errorf("method %v is not existed\n", method)
	}

	parts := patternParser(pattern)
	r.roots[method].search(parts, 0)

	p := keyBuilder(method, pattern)
	if _, ok := r.handlers[p]; !ok {
		return nil, fmt.Errorf("%v is not existed", p)
	}

	return r.handlers[p], nil
}

func keyBuilder(method string, pattern string) string {
	return method + "-" + pattern
}

func patternParser(pattern string) []string {
	parts := strings.Split(pattern, "/")

	ps := make([]string, 0)
	for _, part := range parts {
		if part == "" {
			continue
		}

		ps = append(ps, part)
		if part == "*" {
			break
		}
	}

	return ps
}
