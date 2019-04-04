package privMux

import (
	"io"
	"net/http"
)

type PrivHandler struct{}
func (*PrivHandler)ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	if fn, ok := g_mux[req.URL.String()]; ok {
		fn(rsp, req)
		return
	}
	io.WriteString(rsp, "Bad Request")
}

var g_mux map[string] func(http.ResponseWriter, *http.Request)
func InitUrlMux() {
	g_mux = make(map[string] func(http.ResponseWriter, *http.Request)) 
}

func SetUrlHandler(url string, fn func(http.ResponseWriter, *http.Request)) {
	g_mux[url] = fn
}