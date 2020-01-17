package cache

import (
	"fmt"

	"github.com/superhero-suggestions/internal/cache/model"
)

// SetSuggestions stores multiple suggestions into Redis cache.
func (c *Cache) SetSuggestions(suggestions []model.Superhero) error {
	for _, suggestion := range suggestions {
		err := c.Redis.Set(
			fmt.Sprintf(c.SuggestionKeyFormat, suggestion.ID),
			suggestion,
			0,
		).Err()
		if err != nil {
			return err
		}
	}

	return nil
}
