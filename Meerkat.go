package meerkat

import (
	"net/http"
	"sync"
)

type HTTPErrorHandler func(error, *Context)

type Meerkat struct{
	Server http.Server
	router *HttpRouter
	contextPool sync.Pool
	httpErrorHandler HTTPErrorHandler
	binder Binder
}

func New()	*Meerkat{
	obj := Meerkat{
		router:NewHttpRouter(),
		httpErrorHandler:func(err error,context *Context){ LogInstance().Errorln(err)},
		binder:&DefaultBinder{},
	}
	obj.contextPool.New =func() interface{} {
		return NewContext(&obj)
	}
	return &obj
}

func (meerkat *Meerkat) ServeHTTP(resp http.ResponseWriter,req *http.Request){
	url := req.URL.String()
	methon := req.Method
	handler,len,param := meerkat.router.GetHandler(methon,url)
	if nil != handler{
		context := meerkat.contextPool.Get().(*Context)
		context.Reset(req,resp)
		context.SetRouteParam(url,len,param)
		err :=handler(context)
		meerkat.contextPool.Put(context)
		if err != nil {
			meerkat.HttpErrorHandle(err, context)
		}
	}
}

func (meerkat *Meerkat) HttpErrorHandle(err error,context *Context){
	if meerkat.httpErrorHandler != nil{
		meerkat.httpErrorHandler(err,context)
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

func (meerkat *Meerkat) CONNECT(path string,handler HttpHandler){
	meerkat.router.CONNECT(path,handler)
}

func(meerkat *Meerkat) DELETE(path string,handler HttpHandler){
	meerkat.router.DELETE(path,handler)
}


func (meerkat *Meerkat) HEAD(path string,handler HttpHandler){
	meerkat.router.HEAD(path,handler)
}


func (meerkat *Meerkat) OPTIONS(path string,handler HttpHandler){
	meerkat.router.OPTIONS(path,handler)
}


func (meerkat *Meerkat) PATCH(path string,handler HttpHandler){
	meerkat.router.PATCH(path,handler)
}


func (meerkat *Meerkat) POST(path string,handler HttpHandler){
	meerkat.router.POST(path,handler)
}


func (meerkat *Meerkat) PUT(path string,handler HttpHandler){
	meerkat.router.PUT(path,handler)
}


func (meerkat *Meerkat) TRACE(path string,handler HttpHandler){
	meerkat.router.TRACE(path,handler)
}
