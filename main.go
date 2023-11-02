package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/index", order)
	http.HandleFunc("/", hello)
	server := &http.Server{
		Addr: ":8888",
	}

	fmt.Println("server startup...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}

}

func order(w http.ResponseWriter, r *http.Request) {

}

func hello(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello liwenzhou.com!"))
}
func fbl(i int) int {
	if i <= 1 {
		return i
	} else {
		return fbl(i-1) + fbl(i-2)
	}
}
