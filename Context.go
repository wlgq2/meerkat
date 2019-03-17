package meerkat

import "net/http"
import "io"

type Context struct{
    request  *http.Request
    response *Response
    path     string
    //handler  HttpHandler

}

func NewContext() *Context{

    return &Context{
        request:nil,
        response:&Response{nil,false},
        path:"",
        //handler:nil,
    }
}



func (obj *Context) Reset(req *http.Request, resp http.ResponseWriter){
    obj.response.Reset(resp)
    obj.request = req
}


func (obj *Context) String(code int, str string) error{
    return obj.response.Bytes(code,[]byte(str))
}

func (obj *Context)HTML(code int, html string) error{
    return obj.response.HTML(code,[]byte(html))
}

func (obj *Context)HTMLBlob(code int, data []byte) error{
    return obj.response.HTML(code,data)
}

func (obj *Context) JSON(code int, data interface{}) error{
    return obj.response.JSONEncode(code,data,"");
}

func (obj *Context) JSONBlob(code int, data []byte) error{
    return obj.response.JSON(code,data)
}

func (obj* Context) JSONP(code int, callback string, data interface{}) error{
    return obj.response.JSONPEncode(code,callback,data)
}

func (obj *Context) JSONPBlob(code int, callback string, data []byte) error{
    return obj.response.JSONP(code,callback,data)
}

func (obj *Context) XML(code int, data interface{}) error{
    return obj.response.XMLEncode(code,data,"")
}

func (obj *Context) XMLBlob(code int, data []byte) error{
    return obj.response.XML(code,data)
}

func (obj *Context) Blob(code int, contentType string, data []byte) error{
    return obj.response.Write(code,contentType,data)
}

func (obj *Context) Stream(code int, contentType string, reader io.Reader) error{
    return obj.response.Stream(code,contentType,reader)
}

func (obj *Context) NoContent(code int) error{
    return obj.response.NoContent(code)
}

func (obj *Context) Redirect(code int, url string) error{
    return obj.response.Redirect(code,url)
}

func (obj *Context) File(file string) error{
    return obj.response.File(file,obj.request)
}
