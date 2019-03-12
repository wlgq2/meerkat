package meerkat

import "net/http"

func NewContext() *Context{

    return &Context{
        request:nil,
        response:nil,
        path:"",
        handler:nil,
    }
}

type Context struct{
    request  *http.Request
    response *Response
    path     string
    handler  HttpHandler

}

func (obj *Context) Reset(req *http.Request, resp http.ResponseWriter){
    obj.response.Reset(resp)
    obj.request = req
}
