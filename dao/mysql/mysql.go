package mysql

import (
	"fmt"
	"gin_web/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

// 初始化函数
func Init(cfg *settings.MySQLConfig) (err error) {
	// 从配置文件获取配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.MySQLHost,
		cfg.MySQLPort,
		cfg.Dbname,
		// viper.GetString("mysql.user"),
		// viper.GetString("mysql.password"),
		// viper.GetString("mysql.host"),
		// viper.GetInt("mysql.port"),
		// viper.GetString("mysql.dbname"),
	)

	// 链接数据库
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB Faild, ", zap.Error(err))
		return
	}

	// 数据库配置
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxOpenConns(viper.GetInt("mysql.max_idle_conns"))
	zap.L().Info("mysql Connect Success")
	return
}

// 关闭链接函数
func Close() {
	_ = db.Close()
}
