package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds         = 7 * 24 * 3600
	VoteScore        float64 = 432
	PostPerAge               = 20
)

var (
	ErrorVoteTimeExpire = errors.New("vote time is already over")
)

// NOTE: user float64  to calculate score
func VoteForPost(userID, postID string, value float64) error {

	// NOTE: check post time
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()

	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrorVoteTimeExpire
	}

	// NOTE: check current user's vote record
	oldValue := client.ZScore(getRedisKey(keyPostVotedZSetPrefix+postID), userID).Val()
	var dir float64
	if value > oldValue {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(oldValue - value)
	_, err := client.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*VoteScore, postID).Result()

	if err != nil {
		return err
	}

	if value == 0 {
		_, err = client.ZRem(getRedisKey(keyPostVotedZSetPrefix + postID)).Result()
	} else {
		_, err = client.ZAdd(getRedisKey(keyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value,
			Member: postID,
		}).Result()
	}

	return err
}
