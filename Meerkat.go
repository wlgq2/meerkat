package meerkat

import (
	"net/http"
	"sync"
)

type Meerkat struct{
	Server http.Server
	router *HttpRouter
	contextPool sync.Pool
}

func New()	*Meerkat{
	obj := Meerkat{}
	obj.router = NewHttpRouter()
	obj.contextPool.New =func() interface{} {
		return NewContext()
	}
	return &obj
}
func (meerkat *Meerkat) ServeHTTP(resp http.ResponseWriter,req *http.Request){
	url := req.URL.String()
	methon := req.Method
	handler := meerkat.router.GetHandler(methon,url)
	if nil != handler{
		context := meerkat.contextPool.Get().(*Context)
		context.Reset(req,resp)
		handler(context)
		meerkat.contextPool.Put(context)
	}
}

func (meerkat *Meerkat) Start(addr string) error {
	meerkat.Server.Addr = addr
	meerkat.Server.Handler = meerkat
	return meerkat.Server.ListenAndServe()
}

func (meerkat *Meerkat) GET(path string,handler HttpHandler){
	meerkat.router.GET(path,handler)
}

