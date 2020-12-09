package main

import (
	"fmt"

	"loong.me/gopher/controller/hello"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.Use(cors.Default())

	// 测试服务
	hello.Routes(router)

	// 监听端口
	var port string = "8080"

	listenAddress := fmt.Sprintf(":%s", port)

	router.Run(listenAddress)
}
