package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (ctx *Context) ReadJSON(req interface{}) error {
	body, err := io.ReadAll(ctx.R.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		return err
	}
	return nil
}

func (ctx *Context) WriterJSON(code int, resp interface{}) error {
	ctx.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = ctx.W.Write(respJson)
	return err
}
func (ctx *Context) OkJson(resp interface{}) error {
	return ctx.WriterJSON(http.StatusOK, resp)
}
func (ctx *Context) SystemErrorJson(resp interface{}) error {
	return ctx.WriterJSON(http.StatusInternalServerError, resp)
}
func (ctx *Context) BadRequestJson(resp interface{}) error {
	return ctx.WriterJSON(http.StatusBadRequest, resp)
}
func NewContext(writer http.ResponseWriter,
	request *http.Request) (ctx *Context) {
	ctx = &Context{
		W: writer,
		R: request,
	}
	return ctx
}
