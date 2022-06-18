package middleware

import (
	"Gin/common"
	"Gin/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/* 认证中间件 */

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取authorization header
		tokenString := context.GetHeader("Authorization")

		// validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}

		// 通过验证后获取claim中的UserId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)


		// 用户
		if user.ID == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}

		// 如果用户存在 将user信息写入上下文
		context.Set("user", user)

		context.Next()
	}
}