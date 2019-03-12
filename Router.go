package meerkat

type HttpHandler func(*Context) error

type Router struct{
	RadixTree
}

func (router *Router) GetHandler(path string) HttpHandler{
	value := router.GetValue(path)
	if value != nil {
		return value.(HttpHandler)
	}
	return nil

}