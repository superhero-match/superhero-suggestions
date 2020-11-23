/*
  Copyright (C) 2019 - 2021 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/superhero-match/superhero-suggestions/internal/cache/model"
)

// GetSuggestions fetches suggestions from cache.
func (c *Cache) GetSuggestions(keys []string) (suggestions []model.Superhero, err error) {
	for _, key := range keys {
		res, err := c.Redis.Get(fmt.Sprintf(c.SuggestionKeyFormat, key)).Result()
		if err != nil && err != redis.Nil {
			return nil, err
		}

		if len(res) == 0 {
			continue
		}

		var suggestion model.Superhero

		if err := suggestion.UnmarshalBinary([]byte(res)); err != nil {
			return nil, err
		}

		suggestions = append(suggestions, suggestion)
	}

	return suggestions, nil
}
