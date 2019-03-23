package meerkat

import "strings"

type HttpHandler func(*Context) error

type Router struct{
	tree RadixTree
}

type TreeNodeType struct{
	hanler HttpHandler
	param string
	len int
}

func (router *Router) GetHandler(path string) (HttpHandler,int,string){
	value := router.tree.GetValue(path)
	if value != nil {
		node :=value.(TreeNodeType)
		return node.hanler,node.len,node.param
	}
	return nil,0,""

}
func (router *Router) Insert(key string,value HttpHandler){
	//param 参数
	pos := strings.Index(key,":")
	param :=""
	len :=0
	if pos >=0{
		param = key[pos+1:]
		key = key[:pos+1]
		len = pos
	}
	//路由*
	pos =strings.Index(key,"*")
	if pos >=0{
		key = key[:pos+1]
		param = ""
		len = 0
	}
	node :=TreeNodeType{
		len:len,
		hanler:value,
		param:param,
	};

	router.tree.Insert(key,node)
}
