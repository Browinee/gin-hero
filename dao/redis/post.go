package redis

import (
	"master-gin/constants"
	"master-gin/models"
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
