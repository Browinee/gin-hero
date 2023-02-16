package redis

import (
	"errors"
	"time"
)

const (
	oneWeekInSeconds         = 7 * 24 * 3600
	VoteScore        float64 = 432
	PostPerAge               = 20
)

var (
	ErrorVoteTimeExpire = errors.New("Vote time is already over.")
)

// NOTE: user float64  to calculate score
func VoteForPost(userID, postID string, value float64) error {
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()

	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrorVoteTimeExpire
	}

	// NOTE: check  current user's vote record
	oldValue := client.ZScore(getRedisKey(keyPostVotedZSetPrefix+postID), userID).Val()

	Math.abs()

}
