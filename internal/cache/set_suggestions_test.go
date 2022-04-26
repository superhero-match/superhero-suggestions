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
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"

	"github.com/superhero-match/superhero-suggestions/internal/cache/model"
)

func TestCache_SetSuggestions(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	mockProfilePictures := make([]model.ProfilePicture, 0)
	mockSuggestions := make([]model.Superhero, 0)
	mockSuggestion := model.Superhero{
		ID:                "123456789",
		SuperheroName:     "Unit Tester 1",
		MainProfilePicURL: "https://www.test.com",
		ProfilePictures:   mockProfilePictures,
		Gender:            2,
		Age:               30,
		Lat:               0.123456789,
		Lon:               0.123456789,
		Birthday:          "1985-04-26T12:00:00",
		Country:           "Test Country",
		City:              "Test City",
		SuperPower:        "Unit Testing",
		AccountType:       "FREE",
		CreatedAt:         "2022-04-26T12:00:00",
	}
	mockSuggestions = append(mockSuggestions, mockSuggestion)

	key := fmt.Sprintf("suggestion.%s", "123456789")

	exp := time.Duration(0)

	mock := redismock.NewNiceMock(client)
	mock.On("Set", key, mockSuggestion, exp).Return(redis.NewStatusResult("", nil))

	mockCache := &cache{
		Redis:               mock,
		LikesKeyFormat:      "likes.for.%s",
		SuggestionKeyFormat: "suggestion.%s",
	}

	err = mockCache.SetSuggestions(mockSuggestions)
	assert.NoError(t, err)
}
