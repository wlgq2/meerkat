package main

import (
	"github.com/wlgq2/meerkat"
	"net/http"
)

type  Test struct {
Name  string `json:"name" form:"name" query:"name"`
Value string `json:"value" form:"value" query:"value"`
}

func main() {
	server := meerkat.New()
/* commond:
	curl \
	  -X POST \
	  http://127.0.0.1:8001/ \
	  -H 'Content-Type: application/json' \
	  -d '{"name":"test","value":"test"}'

	curl \
	  -X POST \
	  http://127.0.0.1:8001/ \
	  -d 'name=test' \
	  -d 'value=test'

	curl \
	  -X GET \
	  http://127.0.0.1:8001/\?name\=test\&value\=test
*/
	server.GET("/*" ,func(context *meerkat.Context) (err error) {
		test := new(Test)
		if err = context.Bind(test); err != nil {//
			return
		}
		return context.JSON(http.StatusOK, test)
	})
	server.POST("/*" ,func(context *meerkat.Context) (err error) {
		test := new(Test)
		if err = context.Bind(test); err != nil {
			return
		}
		return context.JSON(http.StatusOK, test)
	})
	meerkat.LogInstance().Fatalln(server.Start(":8001"))
}
