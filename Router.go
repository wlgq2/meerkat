package meerkat

type HttpHandler func(string)

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