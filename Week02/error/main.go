package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	config "loong.me/gopher/controller"
)

func main() {

	router := gin.Default()

	router.Use(cors.Default())

	// 测试服务
	config.Routes(router)

	// 监听端口
	var port string = "8080"

	listenAddress := fmt.Sprintf(":%s", port)

	router.Run(listenAddress)
}
