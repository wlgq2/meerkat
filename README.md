# meerkat
meerkat is a web framework bese golang.

## Feature

* Customized effective http's router base on radix tree.
* Request data's binding from json, parameter etc.
* Support response with string, json, jsonp, html, xml, redirect, file, etc.
* Static files support.
* Template support.
* Extensible middleware.
* Custom log interface.

## Quick start
```go
package main

import (
	"github.com/wlgq2/meerkat"
	"net/http"
)

func main() {
	server := meerkat.New()
	server.GET("/hello",func(context *meerkat.Context) error{
		return context.String(http.StatusOK,"hello world.")
	})
	server.GET("/test*",func(context *meerkat.Context) error{
		return context.String(http.StatusOK,"test.")
	})
	server.GET("/set:param",func(context *meerkat.Context) error{
		return context.String(http.StatusOK,context.GetRouteParam("param"))
	})
	meerkat.LogInstance().Fatalln(server.Start(":8000"))
}

```

## More Examples
[meerkat's examples](https://github.com/wlgq2/meerkat_example)
