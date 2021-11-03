package models

// 定义请求结构体  添加参数校验Tag
// binding validator模块对应的Tag

// 注册数据结构体
type ParamSignUp struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Re_password string `json:"re_password" binding:"required,eqfield=Password"`
	Email       string `json:"email" binding:"required"`
	// required 必须要有值 否则 Controllers中的ShouldBindJSON会报错
}

// 登录数据结构体
type ParamLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
