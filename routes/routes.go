package routes

import (
	"gin_web/controller"
	"gin_web/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 单独建立一个模块进行管理
func Setup(mode string) *gin.Engine {
	// 当设置为发布模式时,同步设置gin框架模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // Gin设置成发布模式
	}

	r := gin.New()
	// 关联中间件  日志处理
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册模块
	r.POST("/signup", controller.SignUpHandler)
	// 登录模块
	r.POST("/login", controller.LoginHandler)

	// 登录后才能进行访问的资源
	r.GET("/ping", func(c *gin.Context) {
		// 如果是登录的用户，潘丹请求头中是否有 有效的JWT数据
		// 客户端三种方式携带token： 1.放在请求头 2. 放在请求体 3.放在URL
		// 可以写一个中间件来实现
		isLogin := true
		if isLogin {
			// 如果是登录的用户，返回pong
			c.String(http.StatusOK, "pong")
		} else {
			// 否则就提示登录
			c.String(http.StatusOK, "请登录")
		}

	})

	// 404配置
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "未找到路由",
		})
	})
	return r
}
