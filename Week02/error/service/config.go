package service

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	pkgerr "github.com/pkg/errors"
	"loong.me/gopher/internal/errcode"
	"loong.me/gopher/repository"
)

// Hello 测试接口
func Hello(ctx *gin.Context) {
	result, err := repository.Hello()
	if err == nil {
		ctx.JSON(http.StatusOK, result)
	}
	if errors.Is(err, errcode.ErrNotFound) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": pkgerr.Cause(err).Error()})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
