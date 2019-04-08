package main

import (
	"github.com/wlgq2/meerkat"
	"net/http"
	"net/url"
	"github.com/wlgq2/meerkat/middleware"
)

func runServer(port string){
	server := meerkat.New()

	server.GET("/*",func(context *meerkat.Context) error{
		return context.String(http.StatusOK,"server "+port)
	})
	server.Start(port)
}

func main() {
	//运行服务 balance server
	go runServer(":8001")
	go runServer(":8002")
	go runServer(":8003")

	//代理服务 proxy server
	server := meerkat.New()
	url1, _ := url.Parse("http://localhost:8001")
	url2, _ := url.Parse("http://localhost:8002")
	url3, _ := url.Parse("http://localhost:8003")

	server.Use(middleware.Proxy(
		middleware.NewDefaultProxyBalancer(url1,url2,url3)))
	meerkat.LogInstance().Fatalln(server.Start(":8000"))
}

