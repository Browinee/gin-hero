package redis

import (
	"master-gin/constants"
	"master-gin/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size
	// ZRevRange => Desc
	return client.ZRevRange(key, start, end).Result()
}
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// NOTE: there is no tenary selector in go
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == constants.DefaultOrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return getIDsFormKey(key, p.Page, p.PageSize)
}

// NOTE: get vote of each post
func GetPostVoteData(ids []string) (data []int64, err error) {
	// data = make([]int64, 0, len(ids))
	// for _, id := range ids {
	// 	key := getRedisKey(KeyPostVotedZSetPrefix + id)
	// 	v := client.ZCount(key, "1", "1").Val()
	// 	data = append(data, v)
	// }

	// NOTE: use pipeline to reduce RTT
	// reference: https://www.liwenzhou.com/posts/Go/go_redis/
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
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

func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == constants.DefaultOrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	communityKey := getRedisKey(KeyCommunityPostSetPrefix + strconv.Itoa((int(p.CommunityID))))
	cachedKey := orderKey + strconv.Itoa(int(p.CommunityID))
	if client.Exists(orderKey).Val() < 1 {
		pipeline := client.Pipeline()
		pipeline.ZInterStore(cachedKey, redis.ZStore{
			Aggregate: "MAX",
		}, communityKey, orderKey)
		pipeline.Expire(cachedKey, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return getIDsFormKey(cachedKey, p.Page, p.PageSize)
}
