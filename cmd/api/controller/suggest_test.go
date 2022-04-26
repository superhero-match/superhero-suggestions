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

package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/superhero-match/superhero-suggestions/cmd/api/model"
	"github.com/superhero-match/superhero-suggestions/cmd/api/service"
	"github.com/superhero-match/superhero-suggestions/internal/cache"
	chm "github.com/superhero-match/superhero-suggestions/internal/cache/model"
	"github.com/superhero-match/superhero-suggestions/internal/es"
)

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func TestController_Suggest(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer testServer.Close()

	response := `{
							  "took" : 982,
							  "timed_out" : false,
							  "_shards" : {
								"total" : 5,
								"successful" : 5,
								"skipped" : 0,
								"failed" : 0
							  },
							  "hits" : {
								"total" : {
								  "value" : 10000,
								  "relation" : "gte"
								},
								"max_score" : 1.0,
								"hits" : [
								  {
									"_index" : "superhero",
									"_type" : "_doc",
									"_id" : "2ds34f6w-43f5-2344-dsf4-kf9ekw9fke9w",
									"_score" : 1.0,
									"_source" : {
									  "superhero_id" : "123456789",
									  "email" : "test@test.com",
                                      "name" : "Test Tester 1",
                                      "superhero_name": "Unit Tester 1",
                                      "main_profile_pic_url": "https://www.test.com/1",
                                      "profile_pics": [{
                                        "id": 1,
                                        "superhero_id": "123456789",
                                        "url": "https://www.test.com/2",
                                        "position": 1 
                                      }],
                                      "gender": 1,
                                      "looking_for_gender": 2,
                                      "age": 36,
                                      "looking_for_age_min": 25,
                                      "looking_for_age_max": 45,
                                      "looking_for_distance_max": 50,
                                      "distance_unit": "km",
                                      "location": {
                                        "lat": 0.123456789,
                                        "lon": 0.123456789
                                      },
									  "birthday" : "1985-04-26T12:00:00",
									  "country" : "Test Country",
									  "city" : "Test City",
									  "superpower" : "Unit Testing",
									  "account_type" : "FREE",
									  "created_at" : "2022-04-26T12:00:00"
									}
								  }
								]
							}
					}`

	mockClient, err := es.MockElasticSearchClient(testServer.URL, response)
	assert.NoError(t, err)

	mockEs := es.New(mockClient, "superhero", 50)

	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	mockGetLikesResult := make([]string, 0)
	mockGetLikesResult = append(mockGetLikesResult, "987654321")

	key := fmt.Sprintf("likes.for.%s", "123456789")

	mock := redismock.NewNiceMock(client)
	mock.On("SMEMBERS", key).Return(redis.NewStringSliceCmd(mockGetLikesResult, nil)).Once()

	keys := make([]string, 0)
	keys = append(keys, fmt.Sprintf("likes.for.%s", "987654321"))

	mock.On("Del", keys).Return(redis.NewIntCmd("", nil)).Once()

	mockCache := cache.New(mock, "likes.for.%s", "suggestion.%s")

	mockService := service.New(mockEs, mockCache, 10, "choice.%s.%s")

	retrievedSuperheroIDs := make([]string, 0)
	retrievedSuperheroIDs = append(retrievedSuperheroIDs, "987654321")

	likeSuperheroIDs := make([]string, 0)
	likeSuperheroIDs = append(likeSuperheroIDs, "11111111111")

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	mockController := New(mockService, logger, "2006-01-02T15:04:05")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	MockJsonPost(
		ctx,
		map[string]interface{}{
			"id":                    "987654321",
			"lookingForGender":      2,
			"gender":                1,
			"lookingForAgeMin":      25,
			"lookingForAgeMax":      45,
			"maxDistance":           50,
			"distanceUnit":          "km",
			"lat":                   0.123456789,
			"lon":                   0.123456789,
			"superheroIds":          make([]string, 0),
			"retrievedSuperheroIds": make([]string, 0),
			"isEsRequest":           true,
		},
	)

	mockController.Suggest(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestController_SuggestGetCachedSuggestions(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer testServer.Close()

	response := `{
							  "took" : 982,
							  "timed_out" : false,
							  "_shards" : {
								"total" : 5,
								"successful" : 5,
								"skipped" : 0,
								"failed" : 0
							  },
							  "hits" : {
								"total" : {
								  "value" : 10000,
								  "relation" : "gte"
								},
								"max_score" : 1.0,
								"hits" : [
								  {
									"_index" : "superhero",
									"_type" : "_doc",
									"_id" : "2ds34f6w-43f5-2344-dsf4-kf9ekw9fke9w",
									"_score" : 1.0,
									"_source" : {
									  "superhero_id" : "123456789",
									  "email" : "test@test.com",
                                      "name" : "Test Tester 1",
                                      "superhero_name": "Unit Tester 1",
                                      "main_profile_pic_url": "https://www.test.com/1",
                                      "profile_pics": [{
                                        "id": 1,
                                        "superhero_id": "123456789",
                                        "url": "https://www.test.com/2",
                                        "position": 1 
                                      }],
                                      "gender": 1,
                                      "looking_for_gender": 2,
                                      "age": 36,
                                      "looking_for_age_min": 25,
                                      "looking_for_age_max": 45,
                                      "looking_for_distance_max": 50,
                                      "distance_unit": "km",
                                      "location": {
                                        "lat": 0.123456789,
                                        "lon": 0.123456789
                                      },
									  "birthday" : "1985-04-26T12:00:00",
									  "country" : "Test Country",
									  "city" : "Test City",
									  "superpower" : "Unit Testing",
									  "account_type" : "FREE",
									  "created_at" : "2022-04-26T12:00:00"
									}
								  }
								]
							}
					}`

	mockClient, err := es.MockElasticSearchClient(testServer.URL, response)
	assert.NoError(t, err)

	mockEs := es.New(mockClient, "superhero", 50)

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

	mockService := service.New(mockEs, mockCache, 10, "choice.%s.%s")

	retrievedSuperheroIDs := make([]string, 0)
	retrievedSuperheroIDs = append(retrievedSuperheroIDs, "987654321")

	likeSuperheroIDs := make([]string, 0)
	likeSuperheroIDs = append(likeSuperheroIDs, "11111111111")

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	mockController := New(mockService, logger, "2006-01-02T15:04:05")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	MockJsonPost(
		ctx,
		map[string]interface{}{
			"id":                    "987654321",
			"lookingForGender":      2,
			"gender":                1,
			"lookingForAgeMin":      25,
			"lookingForAgeMax":      45,
			"maxDistance":           50,
			"distanceUnit":          "km",
			"lat":                   0.123456789,
			"lon":                   0.123456789,
			"superheroIds":          make([]string, 0),
			"retrievedSuperheroIds": make([]string, 0),
			"isEsRequest":           false,
		},
	)

	mockController.Suggest(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)
}
