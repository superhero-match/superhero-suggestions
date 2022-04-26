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

func TestCache_GetChoice(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	mockResult := `
	  {
        "id": "test-choice-id",
        "choice": 1,
        "superheroID": "987654321",
        "chosenSuperheroID": "123456789",
        "createdAt": "2022-04-26T12:00:00"
      }
	`

	key := fmt.Sprintf("choice.%s.%s", "987654321", "123456789")

	mock := redismock.NewNiceMock(client)
	mock.On("Get", key).Return(redis.NewStringResult(mockResult, nil))

	mockCache := &cache{
		Redis:               mock,
		LikesKeyFormat:      "likes.for.%s",
		SuggestionKeyFormat: "suggestion.%s",
	}

	_, err = mockCache.GetChoice(key)
	assert.NoError(t, err)
}
