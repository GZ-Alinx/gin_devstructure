package redis

import (
	"fmt"
	"gin_web/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

// // 初始化函数
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.RedisHost,
			cfg.RedisPort,
			// viper.GetString("redis.host"),
			// viper.GetInt("redis.port"),
		),
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.Pool_size,

		// Password: viper.GetString("redis.password"),
		// DB:       viper.GetInt("redis.db"),
		// PoolSize: viper.GetInt("redis.pool_size"),
	})
	_, err = rdb.Ping().Result()
	return err
}

// 关闭函数
func Close() {
	rdb.Close()
}
