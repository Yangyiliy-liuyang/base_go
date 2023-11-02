package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := *gin.Engine
	server.GET(http.StatusOK, func(ctx *gin.Context) {

	})
}
