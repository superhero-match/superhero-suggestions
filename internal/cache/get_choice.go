package cache

import (
	"github.com/go-redis/redis"
	"github.com/superhero-suggestions/internal/cache/model"
)

// GetChoice fetches choice(like, dislikes are only in DB) from cache.
func (c *Cache) GetChoice(key string) (*model.Choice, error) {
	res, err := c.Redis.Get(key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	var choice model.Choice

	if err := choice.UnmarshalBinary([]byte(res)); err != nil {
		return nil, err
	}

	return &choice, nil
}
