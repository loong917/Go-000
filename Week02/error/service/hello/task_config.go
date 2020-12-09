package hello

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pkgerr "github.com/pkg/errors"
	"loong.me/gopher/repository/hello"
)

// Hello 测试接口
func Hello(ctx *gin.Context) {
	result, err := hello.Hello()
	if err == nil {
		ctx.JSON(http.StatusOK, result)
	}
	if errors.Is(err, hello.ErrNotFound) {
		log.Print(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": pkgerr.Cause(err).Error()})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
