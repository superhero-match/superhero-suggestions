package cache

import (
	"github.com/superhero-suggestions/internal/cache/model"
)

// GetSuggestions fetches suggestions from cache.
func (c *Cache) GetSuggestions(keys []string) (suggestions []model.Superhero, err error) {
	for _, key := range keys {
		res, err := c.Redis.Get(key).Result()
		if err != nil {
			return nil, err
		}

		var suggestion model.Superhero

		if err := suggestion.UnmarshalBinary([]byte(res)); err != nil {
			return nil, err
		}

		suggestions = append(suggestions, suggestion)
	}

	return suggestions, nil
}
