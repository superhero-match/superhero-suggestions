package cache

import (
	"github.com/superhero-suggestions/internal/cache/model"
)

// SetSuggestions stores multiple suggestions into Redis cache.
func (c *Cache) SetSuggestions (suggestions []model.Superhero) error {
	var pairs []interface{}

	for _, suggestion := range suggestions {
		pairs = append(pairs, suggestion.ID, suggestion)
	}

	return c.Redis.MSet(pairs...).Err()
}
