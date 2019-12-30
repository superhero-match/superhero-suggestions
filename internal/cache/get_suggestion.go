package cache

import (
	"github.com/go-redis/redis"
	"github.com/superhero-suggestions/internal/cache/model"
)

// GetSuggestions fetches suggestions from cache.
func (c *Cache) GetSuggestion(key string) (*model.Superhero, error) {
	res, err := c.Redis.Get(key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	var suggestion model.Superhero

	if err := suggestion.UnmarshalBinary([]byte(res)); err != nil {
		return nil, err
	}

	return &suggestion, nil
}
