package service

import (
	"master-gin/dao/redis"
	"master-gin/models"
	"strconv"

	"go.uber.org/zap"
)

// NOTE: Usecase reference: https://github.com/mao888/bluebell/blob/main/bluebell_backend/logic/vote.go
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Info("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postId", p.PostID),
		zap.Int8("Direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
