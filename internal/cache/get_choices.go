/*
  Copyright (C) 2019 - 2022 MWSOFT
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
	"github.com/go-redis/redis"

	"github.com/superhero-match/superhero-suggestions/internal/cache/model"
)

// GetChoices fetches choices(likes, dislikes are only in DB) from cache.
func (c *cache) GetChoices(keys []string) (choices []model.Choice, err error) {
	for _, key := range keys {
		res, err := c.Redis.Get(key).Result()
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
