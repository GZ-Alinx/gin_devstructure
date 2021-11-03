package service

import (
	"fmt"
	"gin_web/dao/mysqls"
	"gin_web/models"
	"gin_web/pkg/jwt"
	"gin_web/pkg/snowflake"
)

// 注册逻辑
func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	err = mysqls.CheckUserExist(p.Username)
	if err != nil {
		return err
	}

	// 生成UID 雪花算法，Twitter开源由64位整数组成分布式ID，性能较高，并且在单机上递增
	userID := snowflake.GenID()

	// 构造用户数据实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
		Email:    p.Email,
	}
	fmt.Println("构造数据: ", user)
	// 保存至数据库
	fmt.Println("存储用户数据")
	return mysqls.InsertUser(user)
}

// 登录逻辑处理
func Login(p *models.ParamLogin) (token string, err error) {
	// 直接登录
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递是指针，拿到user_id
	if err := mysqls.Login(user); err != nil {
		fmt.Println("登录失败")
		return "", err
	}

	// 生成jwt token
	fmt.Println("生成token")
	return jwt.GenToken(user.UserID, user.Username)
}
