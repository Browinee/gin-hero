package redis

import (
	"fmt"
	"master-gin/settings"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.Port")),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	_, err = client.Ping().Result()
	if err != nil {
		zap.L().Error("connect redis failed, err %v/n", zap.Error(err))
		return
	}
	return err
}

func Close() {
	_ = client.Close()
}
