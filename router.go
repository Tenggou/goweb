package goweb

import (
	"fmt"
	"net/http"
)

/*
保存path到handle function的映射
1. 注册服务，RESTful
2. 服务获取
 */

func PathBuilder(method string, path string) string {
	return method + "-" + path
}

type router struct {
	m map[string]http.HandlerFunc
}

func (r *router) add(method string, path string, fun http.HandlerFunc){
	r.m[PathBuilder(method, path)] = fun
}

func (r *router) get(method string, path string) (http.HandlerFunc, error) {
	p := PathBuilder(method, path)
	if _, ok := r.m[p]; !ok {
		return nil, fmt.Errorf("%v is not existed", p)
	}

	return r.m[p], nil
}