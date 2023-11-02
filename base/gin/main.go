package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//路由方法 路由规则 处理函数

	// 静态路由
	server.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello,world")
	})

	// 参数路由 路径参数
	server.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hi,"+name)
	})
	// GET /order?id=123查询路由
	server.GET("/order", func(ctx *gin.Context) {
		id := ctx.Query("id")
		ctx.String(http.StatusOK, "订单id:"+id)
	})
	// 通配符匹配
	server.GET("/views/*.html", func(ctx *gin.Context) {
		view := ctx.Param(".html")
		ctx.String(http.StatusOK, "view:"+view) //view:/hello.html
	})
	server.Run(":8080")
}
