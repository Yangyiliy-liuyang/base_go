package main

import "net/http"

// Routerable 可路由的接口
type Routerable interface {
	Router(method string,
		pattern string,
		handleFun func(ctx *Context))
}

// Handler 一方面负责具体的路由，一方面路由下去
type Handler interface {
	http.Handler
	Routerable
}
type HandlerBasedOnMap struct {
	//key method + url
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) Router(
	method string,
	pattern string,
	handlerFunc func(ctx *Context)) {

	key := h.key(method, pattern)
	h.handlers[key] = handlerFunc
}
func (h HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	key := h.key(request.Method, request.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		ctx := NewContext(writer, request)
		handler(ctx)
	} else {
		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write([]byte("not any router math"))
	}
}

func (h *HandlerBasedOnMap) key(method string,
	pattern string) string {
	return method + "#" + pattern
}

var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context), 4),
	}
}
