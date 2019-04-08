package main

import (
	"github.com/wlgq2/meerkat"
	"net/http"
)

func main() {
	server := meerkat.New()
	server.GET("/hello",func(context *meerkat.Context) error{
		return context.String(http.StatusOK,"hello world~")
	})
	server.GET("/test*",func(context *meerkat.Context) error{
		return context.String(http.StatusOK,context.GetRouteParam("*"))
	})
	server.GET("/set:param",func(context *meerkat.Context) error{
		return context.String(http.StatusOK,context.GetRouteParam("param"))
	})
	meerkat.LogInstance().Fatalln(server.Start(":8000"))
}

