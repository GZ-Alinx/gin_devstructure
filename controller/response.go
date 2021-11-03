package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	根据错误码返回信息
{
	"code": 1001 	// 错误代码
	"msg": xx 		// 提示信息
	"data":"{}"		// 数据
}

*/

// 返回响应函数
type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// 错误返回值
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// 成功返回值
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

// 成功返回值
func ResponseErrorWithMsg(c *gin.Context, code ResCode, Msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  Msg,
		Data: nil,
	})
}
