package pkg

import (
	"context"
	"time"

	goredis "github.com/redis/go-redis/v9"

	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"
)

type redis struct {
	redis *goredis.Client
	ctx   context.Context
}

func NewRedis(ctx context.Context, url string) (pinf.IRedis, error) {
	parseURL, err := goredis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	rdc := goredis.NewClient(&goredis.Options{
		Addr:            parseURL.Addr,
		Password:        parseURL.Password,
		MaxRetries:      10,
		PoolSize:        20,
		PoolFIFO:        true,
		ReadTimeout:     time.Duration(time.Second * 30),
		WriteTimeout:    time.Duration(time.Second * 30),
		DialTimeout:     time.Duration(time.Second * 60),
		MinRetryBackoff: time.Duration(time.Second * 60),
		MaxRetryBackoff: time.Duration(time.Second * 120),
	})

	Logrus("info", "Redis connection success")
	return &redis{redis: rdc, ctx: ctx}, nil
}

func (h redis) SetEx(key string, expiration time.Duration, value any) error {
	cmd := h.redis.SetEx(h.ctx, key, value, expiration)

	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}

func (h redis) Get(key string) ([]byte, error) {
	cmd := h.redis.Get(h.ctx, key)

	if err := cmd.Err(); err != nil {
		return nil, err
	}

	res := cmd.Val()
	return []byte(res), nil
}

func (h redis) Del(key string) (int64, error) {
	cmd := h.redis.Del(h.ctx, key)

	if err := cmd.Err(); err != nil {
		return 0, err
	}

	return cmd.Val(), nil
}

func (h redis) Exists(key string) (int64, error) {
	cmd := h.redis.Exists(h.ctx, key)

	if err := cmd.Err(); err != nil {
		return 0, err
	}

	return cmd.Val(), nil
}

func (h redis) HSet(key string, values ...any) error {
	cmd := h.redis.HSet(h.ctx, key, values...)

	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}

func (h redis) HSetEx(key string, expiration time.Duration, values ...any) error {
	cmd := h.redis.HSet(h.ctx, key, values)
	h.redis.Expire(h.ctx, key, expiration)

	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}

func (h redis) HGet(key, field string) ([]byte, error) {
	cmd := h.redis.HGet(h.ctx, key, field)

	if err := cmd.Err(); err != nil {
		return nil, err
	}

	res := cmd.Val()
	return []byte(res), nil
}

func (h redis) HGetAll(key string) (map[string]string, error) {
	cmd := h.redis.HGetAll(h.ctx, key)

	if err := cmd.Err(); err != nil {
		return nil, err
	}

	res := cmd.Val()
	return res, nil
}

func (h redis) HExists(key, field string) (bool, error) {
	cmd := h.redis.HExists(h.ctx, key, field)

	if err := cmd.Err(); err != nil {
		return false, err
	}

	res := cmd.Val()
	return res, nil
}

func (h redis) HDel(key string, fields ...string) (int64, error) {
	cmd := h.redis.HDel(h.ctx, key, fields...)

	if err := cmd.Err(); err != nil {
		return -1, err
	}

	res := cmd.Val()
	return res, nil
}

func (h redis) HIncrByFloat(key, field string, incr float64) (float64, error) {
	cmd := h.redis.HIncrByFloat(h.ctx, key, field, incr)

	if err := cmd.Err(); err != nil {
		return -1, err
	}

	res := cmd.Val()
	return res, nil
}

func (h redis) LPush(key string, values ...any) ([]string, error) {
	cmd := h.redis.LPush(h.ctx, key, values)
	if err := cmd.Err(); err != nil {
		return nil, err
	}

	cmdl := h.redis.LRange(h.ctx, key, 0, cmd.Val())
	if err := cmdl.Err(); err != nil {
		return nil, err
	}

	res := cmdl.Val()
	return res, nil
}

func (h redis) SAdd(key string, members ...any) (int64, error) {
	cmd := h.redis.SAdd(h.ctx, key, members...)
	if err := cmd.Err(); err != nil {
		return -1, err
	}

	res := cmd.Val()
	return res, nil
}

func (h redis) SIsMember(key string, member any) (bool, error) {
	cmdl := h.redis.SIsMember(h.ctx, key, member)
	if err := cmdl.Err(); err != nil {
		return false, err
	}

	res := cmdl.Val()
	return res, nil
}

func (h redis) SRem(key string, member any) (int64, error) {
	cmdl := h.redis.SRem(h.ctx, key, member)
	if err := cmdl.Err(); err != nil {
		return -1, err
	}

	res := cmdl.Val()
	return res, nil
}

func (h redis) SMembers(key string) ([]string, error) {
	cmdl := h.redis.SMembers(h.ctx, key)
	if err := cmdl.Err(); err != nil {
		return nil, err
	}

	res := cmdl.Val()
	return res, nil
}

func (h redis) ZAdd(key string, members ...goredis.Z) (int64, error) {
	cmd := h.redis.ZAdd(h.ctx, key)
	if err := cmd.Err(); err != nil {
		return -1, err
	}

	res := cmd.Val()
	return res, nil
}

func (h redis) ZRange(key string) ([]string, error) {
	cmdl := h.redis.ZRange(h.ctx, key, 0, -1)
	if err := cmdl.Err(); err != nil {
		return nil, err
	}

	res := cmdl.Val()
	return res, nil
}
