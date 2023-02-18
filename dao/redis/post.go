package redis

import (
	"master-gin/constants"
	"master-gin/models"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// NOTE: there is no tenary selector in go
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == constants.DefaultOrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	start := (p.Page - 1) * p.PageSize
	end := start + p.PageSize
	// ZRevRange => Desc
	return client.ZRevRange(key, start, end).Result()
}

// NOTE: get vote of each post
func GetPostVoteData(ids []string) (data []int64, err error) {
	// data = make([]int64, 0, len(ids))
	// for _, id := range ids {
	// 	key := getRedisKey(keyPostVotedZSetPrefix + id)
	// 	v := client.ZCount(key, "1", "1").Val()
	// 	data = append(data, v)
	// }

	// NOTE: use pipeline to reduce RTT
	// reference: https://www.liwenzhou.com/posts/Go/go_redis/
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(keyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	zap.L().Debug("redis.GetPostVoteData", zap.Any("cmders", cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
