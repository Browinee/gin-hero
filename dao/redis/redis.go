package redis

import (
	"fmt"
	"master-gin/settings"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.Port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	_, err = rdb.Ping().Result()

	return err
}

func Close() {
	_ = rdb.Close()
}
