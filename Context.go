package meerkat

import "net/http"
import "net/url"
import "io"
import "errors"
import "strings"

const (
    CharsetUTF8                  = "charset=UTF-8"
    ContentTypeJSON              = "application/json"
    ContentTypeJSONP             = "application/javascript"
    ContentTypeHTML              = "text/html"
    ContentTypeXML               = "application/xml"
    ContentTypeXML2              = "text/xml"
    ContentTypePlain             = "text/plain"
    ContentTypePOSTForm          = "application/x-www-form-urlencoded"
    ContentTypeMultipartPOSTForm = "multipart/form-data"
)

const DefaultMemory = 32 * 1024 * 1024

type Context struct{
    meerkat *Meerkat
    request  *http.Request
    response *Response
    path     string
    query    url.Values
    routeParams   map[string]string
    //func DefaultBinder(req *http.Request) BinderFunc
}

func NewContext(meerkat *Meerkat) *Context{

    return &Context{
        meerkat:meerkat,
        request:nil,
        response:&Response{nil,false},
        path:"",
    }
}

func (context *Context) Reset(req *http.Request, resp http.ResponseWriter){
    context.response.Reset(resp)
    context.request = req
    context.query = nil
    context.routeParams = make(map[string]string)
}

func (context *Context) SetRouteParam(url string ,len int, param string){
    if param != ""{
        context.routeParams[param] = url[len:]
    }
}

func (context *Context) GetRouteParam(key string) string{
    if value, exist := context.routeParams[key]; exist {
        return value
    } else {
        return ""
    }
}

func (context *Context) QueryParams() url.Values {
    if context.query == nil {
        context.query = context.request.URL.Query()
    }
    return context.query
}

func (context *Context) QueryParam(name string) string {
    if context.query == nil {
        context.query = context.request.URL.Query()
    }
    return context.query.Get(name)
}

func (context *Context) FormParams() (url.Values, error) {
    if strings.HasPrefix(context.request.Header.Get("Content-Type"), ContentTypePOSTForm) {
        if err := context.request.ParseMultipartForm(DefaultMemory); err != nil {
            return nil, err
        }
    } else {
        if err := context.request.ParseForm(); err != nil {
            return nil, err
        }
    }
    return context.request.Form, nil
}

func (context *Context) Bind(obj interface{}) error {
    binder :=  context.meerkat.binder
    if binder ==nil{
        return errors.New("undefine binder.")
    }
    handle := binder.GetBindHanler(context.request)
    if handle != nil {
        return handle(context,obj)
    }
    return errors.New("Can not find bind type.")
}

func (context *Context) String(code int, str string) error{
    return context.response.Bytes(code,[]byte(str))
}

func (context *Context)HTML(code int, html string) error{
    return context.response.HTML(code,[]byte(html))
}

func (context *Context)HTMLBlob(code int, data []byte) error{
    return context.response.HTML(code,data)
}

func (context *Context) JSON(code int, data interface{}) error{
    return context.response.JSONEncode(code,data,"");
}

func (context *Context) JSONBlob(code int, data []byte) error{
    return context.response.JSON(code,data)
}

func (context* Context) JSONP(code int, callback string, data interface{}) error{
    return context.response.JSONPEncode(code,callback,data)
}

func (context *Context) JSONPBlob(code int, callback string, data []byte) error{
    return context.response.JSONP(code,callback,data)
}

func (context *Context) XML(code int, data interface{}) error{
    return context.response.XMLEncode(code,data,"")
}

func (context *Context) XMLBlob(code int, data []byte) error{
    return context.response.XML(code,data)
}

func (context *Context) Blob(code int, contentType string, data []byte) error{
    return context.response.Write(code,contentType,data)
}

func (context *Context) Stream(code int, contentType string, reader io.Reader) error{
    return context.response.Stream(code,contentType,reader)
}

func (context *Context) NoContent(code int) error{
    return context.response.NoContent(code)
}

func (context *Context) Redirect(code int, url string) error{
    return context.response.Redirect(code,url)
}

func (context *Context) File(file string) error{
    return context.response.File(file,context.request)
}


