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

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestCache_GetSuggestions(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	mockResult := `
	  {
        "id": "test-superhero-id",
        "superheroName": "Unit Tester 1",
        "mainProfilePicUrl": "https://test.superhero.com",
        "profilePictures": [],
        "gender": 2,
        "age": 30,
        "lat": 0.123456789,
        "lon": 0.123456789,
        "birthday": "1985-04-26T12:00:00",
        "country": "Test Country",
        "city": "Test City",
        "superpower": "Unit Testing",
        "accountType": "FREE",
        "createdAt": "2022-04-26T12:00:00"
      }
	`
	keys := make([]string, 0)
	keys = append(keys, fmt.Sprintf("suggestion.%s", "123456789"))

	mockKeys := make([]string, 0)
	mockKeys = append(mockKeys, "123456789")

	mock := redismock.NewNiceMock(client)
	mock.On("Get", keys[0]).Return(redis.NewStringResult(mockResult, nil))

	mockCache := &cache{
		Redis:               mock,
		LikesKeyFormat:      "likes.for.%s",
		SuggestionKeyFormat: "suggestion.%s",
	}

	_, err = mockCache.GetSuggestions(mockKeys)
	assert.NoError(t, err)
}
