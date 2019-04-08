package main

import (
	"fmt"
	"github.com/wlgq2/meerkat"
	"net/http"
)

type TestData struct {
	Num int  `json:"num" xml:"num"`
	Str string `json:"str" xml:"str"`
}


func main() {
	server := meerkat.New()
	server.GET("/string", func(context *meerkat.Context) error{
		return  context.String(http.StatusOK,"hello world!")
	})
	server.GET("/json", func(context *meerkat.Context) error{
		resp := TestData{1024,"test..."}
		return  context.JSON(http.StatusOK,resp)
	})
	server.GET("/jsonblob", func(context *meerkat.Context) error{
		resp := `{ "int":"123" , "name":"test" }`
		return  context.JSONBlob(http.StatusOK,[]byte(resp))
	})
	server.GET("/jsonp", func(context *meerkat.Context) error{
		resp := TestData{1024,"test..."}
		return  context.JSONP(http.StatusOK,"callback",resp)
	})
	server.GET("/html", func(context *meerkat.Context) error{
		html :=`<!DOCTYPE html>
		<html><head><title>This is a title</title></head><body>
		<p>Hello world!</p></body></html>`
		return  context.HTML(http.StatusOK,html)
	})
	server.GET("/xml", func(context *meerkat.Context) error{
		resp := TestData{1024,"test..."}
		return  context.XML(http.StatusOK,resp)
	})
	server.GET("/xmlblob", func(context *meerkat.Context) error{
		resp := `<test> hello world! </test>`
		return  context.XMLBlob(http.StatusOK,[]byte(resp))
	})
	server.GET("/null", func(context *meerkat.Context) error{
		return  context.NoContent(http.StatusOK)
	})
	server.GET("/redirect", func(context *meerkat.Context) error{
		return  context.Redirect(http.StatusMovedPermanently,"string")
	})
	server.GET("/file", func(context *meerkat.Context) error{
		return  context.File("test.png")
	})
	//中间件
	rst := ""
	server.GET("/middleware", func (context *meerkat.Context) error {
		return context.String(http.StatusOK, rst)
	},
	func( handler meerkat.HttpHandler) meerkat.HttpHandler{
		rst ="middleware test."
		return handler
	})
	err := server.Start(":8002")
	fmt.Println(err)
}