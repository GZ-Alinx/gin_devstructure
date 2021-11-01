package routes

import (
	"gin_web/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 单独建立一个模块进行管理
func Setup() *gin.Engine {
	r := gin.New()
	// 关联中间件  日志处理
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 添加路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return r
}
