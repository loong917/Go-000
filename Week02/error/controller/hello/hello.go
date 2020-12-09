package hello

import (
	"github.com/gin-gonic/gin"
	"loong.me/gopher/service/hello"
)

// Routes 测试路由
func Routes(route *gin.Engine) {
	router := route.Group("/hello")
	{

		// 测试
		router.GET("/life", hello.Hello)

	}
}
