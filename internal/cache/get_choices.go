package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/superhero-suggestions/internal/cache/model"
)

// GetChoices fetches choices(likes, dislikes are only in DB) from cache.
func (c *Cache) GetChoices(keys []string) (choices []model.Choice, err error) {
	for _, key := range keys {
		res, err := c.Redis.Get(key).Result()
		fmt.Println("GetChoices err: ")
		fmt.Println(res)
		fmt.Println(err)
		if err != nil && err != redis.Nil {
			return nil, err
		}

		if len(res) == 0 {
			continue
		}

		var choice model.Choice

		if err := choice.UnmarshalBinary([]byte(res)); err != nil {
			return nil, err
		}

		choices = append(choices, choice)
	}

	return choices, nil
}