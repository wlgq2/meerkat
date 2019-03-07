package meerkat

import (
	"net/http"
)

type Meerkat struct{
	Server http.Server

	router *HttpRouter
}

func New()	*Meerkat{
	obj := Meerkat{}
	obj.router = NewHttpRouter()
	return &obj
}
func (meerkat *Meerkat) ServeHTTP(resp http.ResponseWriter,req *http.Request){
	url := req.URL.String()
	methon := req.Method
	handler := meerkat.router.GetHandler(methon,url)
	if nil != handler{
		handler("test")
	}
	//resp.Write(test)
}

func (meerkat *Meerkat) Start(addr string){
	meerkat.Server.Addr = addr
	meerkat.Server.Handler = meerkat
	meerkat.Server.ListenAndServe()
}

func (meerkat *Meerkat) GET(path string,handler HttpHandler){
	meerkat.router.GET(path,handler)
}

