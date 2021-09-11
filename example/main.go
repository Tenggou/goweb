package main

import (
	"goweb"
	"net/http"
)

func main(){
	engine := goweb.NewEngine()
	engine.POST("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("this is hello router"))
	})

	_ = engine.Run(":8888")
}