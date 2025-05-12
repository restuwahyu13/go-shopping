package pinf

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type IRedis interface {
	SetEx(key string, expiration time.Duration, value any) error
	Get(key string) ([]byte, error)
	Del(key string) (int64, error)
	Exists(key string) (int64, error)
	HSet(key string, values ...any) error
	HSetEx(key string, expiration time.Duration, values ...any) error
	HGet(key string, field string) ([]byte, error)
	HGetAll(key string) (map[string]string, error)
	HExists(key string, field string) (bool, error)
	HDel(key string, fields ...string) (int64, error)
	HIncrByFloat(key, field string, incr float64) (float64, error)
	LPush(key string, values ...any) ([]string, error)
	SAdd(key string, members ...any) (int64, error)
	SIsMember(key string, member any) (bool, error)
	SRem(key string, member any) (int64, error)
	SMembers(key string) ([]string, error)
	ZAdd(key string, members ...redis.Z) (int64, error)
	ZRange(key string) ([]string, error)
}
