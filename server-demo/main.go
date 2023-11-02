package main

import "net/http"

func main() {
	server := NewHttpServer("server-server")
	server.Router(http.MethodGet, "/user/signup", SignUp)
	err := server.Start(":6260")
	if err != nil {
		panic(err)
	}
}
