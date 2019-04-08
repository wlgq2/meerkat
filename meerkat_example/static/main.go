package main

import (
	"github.com/wlgq2/meerkat"
)

func main() {
	server := meerkat.New()
	server.Static("/static", "static/")
	meerkat.LogInstance().Fatalln(server.Start(":8000"))
}

