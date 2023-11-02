package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	Routerable
	Start(address string) error
}
type sdkHttpServer struct {
	Name    string
	handler Handler
}

// Router 注册路由
func (s *sdkHttpServer) Router(
	method string,
	pattern string,
	handlerFunc func(ctx *Context)) {
	s.handler.Router(method, pattern, handlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {

	http.Handle("/", s.handler)
	err := http.ListenAndServe(address, nil)
	return err
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name:    name,
		handler: NewHandlerBasedOnMap(),
	}
}

// SignUp 注册函数
func SignUp(ctx *Context) {
	//实现了读取HTTP请求Body并解析到结构体的功能,主要步骤是:
	//1. 定义了signUpReq结构体来接收请求参数。
	//2. 使用io.ReadAll读取请求Body内容到[]byte。
	//3. 使用json.Unmarshal把byte数组解析到signUpReq结构体。
	//4. 定义了Response结构体来封装响应。
	//5. 使用json.Marshal对Response编码成JSON。
	//6. 把编码后的byte数组写到响应Writer。
	req := &signUpReq{}

	err := ctx.ReadJSON(req)
	if err != nil {
		err := ctx.BadRequestJson(err)
		if err != nil {
			return
		}
		return
	}
	resp := &Response{
		Code: 0,
		Msg:  "",
		Data: 123,
	}
	err = ctx.WriterJSON(http.StatusOK, resp)
	if err != nil {
		fmt.Println("写入失败，err:%s", err)
	}

}

type signUpReq struct {
	// Tag 声明式api
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
