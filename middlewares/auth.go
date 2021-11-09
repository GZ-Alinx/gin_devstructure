package middlewares

import (
	"gin_web/controller"
	"gin_web/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// jwt中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		// 判断请求头中是否存在token值
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort() // 跳出函数处理
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, "", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// parrs[1] 获取tokenString， 我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidAuth)
			c.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next() // 接着处理后续的请求 使用c.Get("username")获取请求的用户信息
	}
}
