package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

const (
	oneWeekInSeconds         = 7 * 24 * 3600
	VoteScore        float64 = 432
	PostPerAge               = 20
)

var (
	ErrorVoteTimeExpire = errors.New("vote time is already over")
	ErrorVoteRepeated   = errors.New("dont' vote the same thig again")
)

// NOTE: user float64  to calculate score
func VoteForPost(userID, postID string, value float64) error {

	// NOTE: check post time
	// ex: zscore ginheropost:time postID
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()

	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		zap.L().Error(" service.VoteForPost time expire", zap.Error(ErrorVoteTimeExpire))
		return ErrorVoteTimeExpire
	}

	// NOTE: check current user's vote record
	oldValue := client.ZScore(getRedisKey(keyPostVotedZSetPrefix+postID), userID).Val()

	if oldValue == value {
		zap.L().Error("vote the same options again")
		return ErrorVoteRepeated
	}
	var dir float64
	if value > oldValue {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(oldValue - value)

	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*VoteScore, postID)

	if value == 0 {
		pipeline.ZRem(getRedisKey(keyPostVotedZSetPrefix+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(keyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value,
			Member: postID,
		})
	}
	_, err := pipeline.Exec()
	return err
}

func CreatePost(postID int64) error {
	// NOTE: transaction
	pipeline := client.TxPipeline()
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	_, err := pipeline.Exec()
	return err
}
