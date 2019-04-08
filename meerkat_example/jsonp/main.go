package main

import (
	"net/http"
	"time"
	"github.com/wlgq2/meerkat"
)
const indexHtml = `<!DOCTYPE html>
<html>
<head><title>Go JSONP Server</title></head>
<body>
<button id="btn">Click to get HTTP header via JSONP</button>
<pre id="result"></pre>
<script>
'use strict';

var btn = document.getElementById("btn");
var result = document.getElementById("result");

function myCallback(acptlang) {
  result.innerHTML = JSON.stringify(acptlang, null, 2);
}

function jsonp() {
  result.innerHTML = "Loading ...";
  var tag = document.createElement("script");
  tag.src = "/jsonp?callback=myCallback";
  document.querySelector("head").appendChild(tag);
}

btn.addEventListener("click", jsonp);
</script>
</body>
</html>`


func main() {
	server := meerkat.New()
	server.GET("/test", func(context *meerkat.Context) error {
		return context.HTML(http.StatusOK, indexHtml)
	})
	// JSONP
	server.GET("/jsonp", func(context *meerkat.Context) error {
		callback := context.QueryParam("callback")
		var content struct {
			Response  string    `json:"response"`
			Timestamp time.Time `json:"timestamp"`
		}
		content.Response = "Sent via JSONP"
		content.Timestamp = time.Now().UTC()
		return context.JSONP(http.StatusOK, callback, &content)
	})
	// Start server
	meerkat.LogInstance().Fatalln(server.Start(":8000"))
}
