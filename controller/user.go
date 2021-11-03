package controller

import (
	"errors"
	"fmt"
	"gin_web/dao/mysqls"
	"gin_web/models"

	"gin_web/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 注册函数
func SignUpHandler(c *gin.Context) {
	// 获取参数和参数校验 前后端分离(JSON数据)
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("\nSignup with invalid param,err: ", zap.Error(err))

		// 判断错误是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok { // 如果不是validator类型，则返回错误
			ResponseError(c, CodeInvalidParam)
			return
		}
		// 返回错误并关联错误日志
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	fmt.Println(p)

	// 业务处理
	if err := service.SignUp(p); err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		if errors.Is(err, mysqls.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		// 返回响应
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, CodeSuccess)
}

// 登录函数
func LoginHandler(c *gin.Context) {
	// 获取请求参数、参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 错误处理
		zap.L().Error("Login with invalid param", zap.Error(err))

		// 判单错误类型是否时validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		// 关联MSG的响应
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 业务逻辑处理
	token, err := service.Login(p)
	if err != nil {
		// 输入日志信息
		zap.L().Error("service.login failed,err:", zap.String("username:", p.Username), zap.Error(err))

		// 判断错误类型 断言
		if errors.Is(err, mysqls.ErrorUserNotExist) {
			ResponseError(c, CodeuserNotExist)
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 返回响应
	ResponseSuccess(c, token)
}
