package meerkat

import (
	"net/http"
	"sync"
	"errors"
	"io"
)

type HTTPErrorHandler func(error, *Context)

type RenderInterface interface {
	Render(io.Writer, string, interface{}, *Context) error
}


type Meerkat struct{
	Server http.Server
	router *HttpRouter
	contextPool sync.Pool
	httpErrorHandler HTTPErrorHandler
	binder Binder
	render RenderInterface
}

func New()	*Meerkat{
	obj := Meerkat{
		router:NewHttpRouter(),
		httpErrorHandler:func(err error,context *Context){ LogInstance().Errorln(err)},
		binder:&DefaultBinder{},
		render:nil,
	}
	obj.contextPool.New =func() interface{} {
		return NewContext(&obj)
	}
	return &obj
}

func (meerkat *Meerkat) ServeHTTP(resp http.ResponseWriter,req *http.Request){
	path := req.URL.Path
	methon := req.Method
	handler,len,param := meerkat.router.GetHandler(methon,path)
	if nil != handler{
		context := meerkat.contextPool.Get().(*Context)
		context.Reset(req,resp)
		context.SetRouteParam(path,len,param)
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


func (meerkat *Meerkat) Static(path string,root string) {
	if root == "" {
		root = "."
	}

	 meerkat.GET(path+"/*", func(context *Context) error {
		file := context.GetRouteParam("*")
		if file == "" {
			return errors.New("not static file.")
		}
		return context.File(root+file)
	})

}

func (meerkat *Meerkat) SetRender(obj RenderInterface) {
	meerkat.render = obj
}

func (meerkat *Meerkat) Render(writer io.Writer, name string, data interface{},context *Context) error{
	if meerkat.render != nil{
		return meerkat.render.Render(writer,name,data,context)
	}
	return  errors.New("undifine render func.")
}