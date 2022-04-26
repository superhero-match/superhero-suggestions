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

func TestCache_DeleteLikes(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	keys := make([]string, 0)
	keys = append(keys, fmt.Sprintf("likes.for.%s", "123456789"))

	mock := redismock.NewNiceMock(client)
	mock.On("Del", keys).Return(redis.NewIntCmd("", nil))

	mockCache := &cache{
		Redis:               mock,
		LikesKeyFormat:      "likes.for.%s",
		SuggestionKeyFormat: "suggestion.%s",
	}

	err = mockCache.DeleteLikes("123456789")
	assert.NoError(t, err)
}
