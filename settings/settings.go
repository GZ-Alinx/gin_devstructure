package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify" // 监控文件变化库
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

//  使用结构体保存配置项
var Conf = new(AppConfig)

type (

	// 配置项目
	AppConfig struct {
		Name       string `mapstructure:"name"`
		Mode       string `mapstructure:"mode"`
		AppPort    int    `mapstructure:"port"`
		ListenHost string `mapstructure:"listenhost"`
		StartTime  string `mapstructure:"start_time"`
		MachineID  int64  `mapstructure:"machine_id"`

		*LogsConfig  `mapstructure:"logs"`
		*MySQLConfig `mapstructure:"mysql"`
		*RedisConfig `mapstructure:"redis"`
	}

	// 日志配置
	LogsConfig struct {
		Level       string `mapstructure:"level"`
		Logfilename string `mapstructure:"logfilename"`
		Max_size    int    `mapstructure:"Max_size"`
		Max_age     int    `mapstructure:"Max_age"`
		Max_backups int    `mapstructure:"Max_backups"`
	}
	// Mysql配置
	MySQLConfig struct {
		MySQLHost      string `mapstructure:"mysqlhost"`
		MySQLPort      int    `mapstructure:"mysqlport"`
		MySQLUser      string `mapstructure:"user"`
		MySQLPassword  string `mapstructure:"mysqlpassword"`
		MySQLDbname    string `mapstructure:"dbname"`
		Max_open_conns int    `mapstructure:"max_open_conns"`
		Max_idle_conns int    `mapstructure:"max_idle_conns"`
	}
	// Redis配置
	RedisConfig struct {
		RedisHost string `mapstructure:"redishost"`
		RedisPort int    `mapstructure:"redisport"`
		Password  string `mapstructure:"redispassword"`
		Db        int    `mapstructure:"db"`
		Pool_size int    `mapstructure:"pool_size"`
	}
)

// 使用viper管理配置
func Init(filePath string) (err error) {

	viper.SetConfigName("config") // 读取文件名称（无需后缀名）
	viper.SetConfigFile(filePath)
	viper.SetConfigType("yaml")    // 指定文件类型
	viper.AddConfigPath("./conf/") // 配置文件查找路径
	viper.AddConfigPath(".")

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		zap.L().Error("viper.ReadInConfig Faild:", zap.Error(err))
		return err
	}

	// 序列化配置到结构体中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Println("viper Unmarshal Failed,", err)
		return
	}

	// 配置自动加载，热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已更新 ...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Println("viper Unmarshal Failed, err", err)
		}
	})
	return nil
}
