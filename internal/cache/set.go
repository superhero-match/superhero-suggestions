package cache

import (
	"github.com/superhero-suggestions/internal/cache/model"
	"time"
)

// SetSuggestions stores multiple suggestions into Redis cache.
func (c *Cache) SetSuggestions(suggestions []model.Superhero) error {
	for _, suggestion := range suggestions {
		if err := c.Redis.Set(suggestion.ID, suggestion, time.Hour).Err(); err != nil {
			return err
		}
	}

	return nil
}
