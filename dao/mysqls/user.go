package mysqls

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"gin_web/models"
)

const secret = "itadminlx@163.com"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

// 检查指定用户是否存在
func CheckUserExist(username string) (err error) {
	fmt.Println("数据库查询用户是否存在")
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		fmt.Println(err)
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 保存至数据库
func InsertUser(user *models.User) (err error) {
	// 将密码加密
	user.Password = encryptPassword(user.Password)
	// 执行入库语句
	sqlStr := `insert into user(user_id, username, password, email) values(?,?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password, user.Email)
	return
}

// 密码加盐
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// 登录函数 数据库信息验证
func Login(user *models.User) (err error) {
	oPassword := user.Password // 登录密码
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)

	if err == sql.ErrNoRows { // 用户不存在
		return ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
