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

// Cache interface defines cache methods.
type Cache interface {
	DeleteLikes(superheroID string) error
	GetChoice(key string) (*model.Choice, error)
	GetChoices(keys []string) (choices []model.Choice, err error)
	GetLikes(superheroID string) ([]string, error)
	GetSuggestions(keys []string) (suggestions []model.Superhero, err error)
	SetSuggestions(suggestions []model.Superhero) error
}

// cache is the Redis client.
type cache struct {
	Redis               redis.Cmdable
	LikesKeyFormat      string
	SuggestionKeyFormat string
}

// New creates a client connection to Redis.
func New(rc redis.Cmdable, likesKeyFormat string, suggestionKeyFormat string) Cache {
	return &cache{
		Redis:               rc,
		LikesKeyFormat:      likesKeyFormat,
		SuggestionKeyFormat: suggestionKeyFormat,
	}
}
