package meerkat

import "net/http"


type HttpRouter struct{
	routers map[string] *Router
}

func NewHttpRouter() *HttpRouter{
	rst := HttpRouter{}
	rst.routers = map[string] *Router{
		http.MethodGet : &Router{},
		http.MethodConnect : &Router{},
		http.MethodDelete : &Router{},
		http.MethodHead : &Router{},
		http.MethodOptions : &Router{},
		http.MethodPatch : &Router{},
		http.MethodPost : &Router{},
		http.MethodPut : &Router{},
		http.MethodTrace : &Router{},
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

func (router *HttpRouter) CONNECT(path string,handler HttpHandler){
	route := router.routers[http.MethodConnect]
	route.Insert(path,handler)
}

func (router *HttpRouter) DELETE(path string,handler HttpHandler){
	route := router.routers[http.MethodDelete]
	route.Insert(path,handler)
}

func (router *HttpRouter) HEAD(path string,handler HttpHandler){
	route := router.routers[http.MethodHead]
	route.Insert(path,handler)
}


func (router *HttpRouter) OPTIONS(path string,handler HttpHandler){
	route := router.routers[http.MethodOptions]
	route.Insert(path,handler)
}

func (router *HttpRouter) PATCH(path string,handler HttpHandler){
	route := router.routers[http.MethodPatch]
	route.Insert(path,handler)
}


func (router *HttpRouter) POST(path string,handler HttpHandler){
	route := router.routers[http.MethodPost]
	route.Insert(path,handler)
}

func (router *HttpRouter) PUT(path string,handler HttpHandler){
	route := router.routers[http.MethodPut]
	route.Insert(path,handler)
}

func (router *HttpRouter) TRACE(path string,handler HttpHandler){
	route := router.routers[http.MethodTrace]
	route.Insert(path,handler)
}
