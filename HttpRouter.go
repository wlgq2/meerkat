package meerkat

import "net/http"


type HttpRouter struct{
	routers map[string] *Router
}

func NewHttpRouter() *HttpRouter{
	rst := HttpRouter{}
	rst.routers = map[string] *Router{
		http.MethodGet : &Router{},
	}
	return &rst
}

func (router *HttpRouter) GetHandler(method string, path string) (HttpHandler,int,string){
	route := router.routers[method]
	return route.GetHandler(path)
}

func (router *HttpRouter) GET(path string,handler HttpHandler){
	route := router.routers[http.MethodGet]
	route.Insert(path,handler)
}