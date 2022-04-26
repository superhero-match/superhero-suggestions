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

package service

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"

	"github.com/superhero-match/superhero-suggestions/cmd/api/model"
	"github.com/superhero-match/superhero-suggestions/internal/cache"
	chm "github.com/superhero-match/superhero-suggestions/internal/cache/model"
)

func TestService_GetCachedSuggestions(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	mockSuggestionResult := `
	 {
	   "id": "987654321",
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
	keys = append(keys, fmt.Sprintf("suggestion.%s", "987654321"))

	mock := redismock.NewNiceMock(client)
	mock.On("Get", keys[0]).Return(redis.NewStringResult(mockSuggestionResult, nil)).Once()

	mockChoicesResult := `
	 {
	   "id": "987654321",
	   "choice": 1,
	   "superheroID": "987654321",
	   "chosenSuperheroID": "123456789",
	   "createdAt": "2022-04-26T12:00:00"
	 }
	`
	choiceKeys := make([]string, 0)
	choiceKeys = append(choiceKeys, fmt.Sprintf("choice.%s.%s", "987654321", "123456789"))

	mock.On("Get", choiceKeys[0]).Return(redis.NewStringResult(mockChoicesResult, nil)).Once()

	mockCache := cache.New(mock, "likes.for.%s", "suggestion.%s")

	mockService := New(nil, mockCache, 10, "choice.%s.%s")

	mockSuperheroIDs := make([]string, 0)
	mockSuperheroIDs = append(mockSuperheroIDs, "987654321")

	mockRequest := model.Request{
		ID:                    "123456789",
		LookingForGender:      2,
		Gender:                1,
		LookingForAgeMin:      25,
		LookingForAgeMax:      45,
		MaxDistance:           50,
		DistanceUnit:          "km",
		Lat:                   0.123456789,
		Lon:                   0.123456789,
		SuperheroIDs:          mockSuperheroIDs,
		RetrievedSuperheroIDs: make([]string, 0),
		IsESRequest:           false,
	}

	_, err = mockService.GetCachedSuggestions(mockRequest)
	assert.NoError(t, err)
}

func TestService_CacheSuggestions(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	mockProfilePictures := make([]model.ProfilePicture, 0)
	mockProfilePictures = append(mockProfilePictures, model.ProfilePicture{
		ID:                1,
		SuperheroID:       "123456789",
		ProfilePictureURL: "https://www.test.com/1",
		Position:          1,
	})
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

	mockCacheProfilePictures := make([]chm.ProfilePicture, 0)
	mockCacheProfilePictures = append(mockCacheProfilePictures, chm.ProfilePicture{
		ID:                1,
		SuperheroID:       "123456789",
		ProfilePictureURL: "https://www.test.com/1",
		Position:          1,
	})

	cachedSuggestion := chm.Superhero{
		ID:                "123456789",
		SuperheroName:     "Unit Tester 1",
		MainProfilePicURL: "https://www.test.com",
		ProfilePictures:   mockCacheProfilePictures,
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

	mock := redismock.NewNiceMock(client)
	mock.On("Set", key, cachedSuggestion, exp).Return(redis.NewStatusResult("", nil))

	mockCache := cache.New(mock, "likes.for.%s", "suggestion.%s")

	mockService := New(nil, mockCache, 10, "choice.%s.%s")

	err = mockService.CacheSuggestions(mockSuggestions)
	assert.NoError(t, err)
}

func TestService_GetCachedChoices(t *testing.T) {
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
	keys := make([]string, 0)
	keys = append(keys, fmt.Sprintf("choice.%s.%s", "987654321", "123456789"))

	mock := redismock.NewNiceMock(client)
	mock.On("Get", keys[0]).Return(redis.NewStringResult(mockResult, nil))

	mockCache := cache.New(mock, "likes.for.%s", "suggestion.%s")

	mockService := New(nil, mockCache, 10, "choice.%s.%s")

	_, err = mockService.GetCachedChoices(keys)
	assert.NoError(t, err)
}

func TestService_GetCachedChoice(t *testing.T) {
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

	mockCache := cache.New(mock, "likes.for.%s", "suggestion.%s")

	mockService := New(nil, mockCache, 10, "choice.%s.%s")

	_, err = mockService.GetCachedChoice(key)
	assert.NoError(t, err)
}

func TestService_GetLikes(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	mockResult := make([]string, 0)
	mockResult = append(mockResult, "987654321")

	key := fmt.Sprintf("likes.for.%s", "123456789")

	mock := redismock.NewNiceMock(client)
	mock.On("SMEMBERS", key).Return(redis.NewStringSliceCmd(mockResult, nil))

	mockCache := cache.New(mock, "likes.for.%s", "suggestion.%s")

	mockService := New(nil, mockCache, 10, "choice.%s.%s")

	_, err = mockService.GetLikes("123456789")
	assert.NoError(t, err)
}

func TestService_DeleteLikes(t *testing.T) {
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

	mockCache := cache.New(mock, "likes.for.%s", "suggestion.%s")

	mockService := New(nil, mockCache, 10, "choice.%s.%s")

	err = mockService.DeleteLikes("123456789")
	assert.NoError(t, err)
}
