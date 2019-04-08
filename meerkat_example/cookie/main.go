package main

import (
	"net/http"
	"github.com/wlgq2/meerkat"
)

func main() {
	server := meerkat.New()
	server.GET("/set", func(context *meerkat.Context) error {
		cookie := new(http.Cookie)
		cookie.Name = "test"
		cookie.Value = "1234"
		context.SetCookie(cookie)
		return context.String(http.StatusOK, "write a cookie.")
	})
	// JSONP
	server.GET("/get", func(context *meerkat.Context) error {
		cookie, err := context.Cookie("test")
		if err != nil {
			return err
		}
		return context.String(http.StatusOK, "cookie{ name:"+cookie.Name+" value:"+cookie.Value+"}")
	})
	// Start server
	meerkat.LogInstance().Fatalln(server.Start(":8000"))
}

