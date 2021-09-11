package main

import (
	"goweb"
	"net/http"
)

func main() {
	engine := goweb.NewEngine()
	engine.POST("/hello", func(ctx *goweb.Context) {
		ctx.String(http.StatusOK, "this is hello POST function")
	})

	engine.GET("/hello", func(ctx *goweb.Context) {
		ctx.String(http.StatusOK, "this is hello get function")
	})

	_ = engine.Run(":8888")
}
