package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* 封装HTTP返回 */

func Response(context *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	context.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(context *gin.Context, data gin.H, msg string) {
	Response(context, http.StatusOK, 200, data, msg)
}

func Fail(context *gin.Context, msg string, data gin.H) {
	Response(context, http.StatusOK, 400, data, msg)
}