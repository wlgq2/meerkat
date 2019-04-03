package middleware

import (
	//"io"
	//"net"
	//"net/http"
	"net/url"
	"net/http/httputil"
	"sync/atomic"
	"github.com/wlgq2/meerkat"
)

type  ProxyBalancer interface {
	Next() *url.URL
}

type  DefaultProxyBalancer struct {
	index uint32
	urls [](*url.URL)
}

func (balancer *DefaultProxyBalancer) Next() *url.URL{
	size := len(balancer.urls)-1
	if size < 0 {
		return nil
	}
	atomic.AddUint32(&balancer.index, 1)
	if int(balancer.index)>(len(balancer.urls)-1) {
		balancer.index = 0
	}
	return balancer.urls[balancer.index]
}

func NewDefaultProxyBalancer(urls ...(*url.URL)) *DefaultProxyBalancer{
	rst := DefaultProxyBalancer{index :0}
	rst.urls = append(rst.urls,urls...)
	return &rst
}

func Proxy(balancer ProxyBalancer) meerkat.MiddlewareFunc {
	return func(handle meerkat.HttpHandler) meerkat.HttpHandler {
		return func(context *meerkat.Context) error {
			target := balancer.Next()
			if target!=nil {
				req := context.Request()
				resp := context.Response().Writer()
				proxy := httputil.NewSingleHostReverseProxy(target)
				proxy.ServeHTTP(resp,req)
			}
			return nil
		}
	}
}
