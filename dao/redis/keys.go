package redis

const (
	KeyPrefix                 = "ginhero:"
	KeyPostTimeZSet           = "post:time"
	KeyPostScoreZSet          = "post:score"
	KeyPostVotedZSetPrefix    = "post:voted:"
	KeyCommunityPostSetPrefix = "community:"
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
