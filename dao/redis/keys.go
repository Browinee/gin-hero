package redis

const (
	KeyPrefix              = "ginhero:"
	KeyPostTimeZSet        = "post:time"
	KeyPostScoreZSet       = "post:score"
	keyPostVotedZSetPrefix = "post:voted:"
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
