package config

import (
	"github.com/gin-gonic/gin"
	"loong.me/gopher/service"
)

// Routes 测试路由
func Routes(route *gin.Engine) {
	router := route.Group("/")
	{

		// 测试
		router.GET("/hello", service.Hello)

	}
}
