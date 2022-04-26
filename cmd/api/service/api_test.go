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
	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	ctrl "github.com/superhero-match/superhero-suggestions/cmd/api/model"
	"github.com/superhero-match/superhero-suggestions/internal/cache"
	"github.com/superhero-match/superhero-suggestions/internal/es"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService_HandleESRequest(t *testing.T) {
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
	keys = append(keys, fmt.Sprintf("choice.%s.%s", "123456789", "987654321"))

	mock := redismock.NewNiceMock(client)
	mock.On("Get", keys[0]).Return(redis.NewStringResult(mockResult, nil))

	mockCache := cache.New(mock, "likes.for.%s", "suggestion.%s")

	mockService := New(mockEs, mockCache, 10, "choice.%s.%s")

	retrievedSuperheroIDs := make([]string, 0)
	retrievedSuperheroIDs = append(retrievedSuperheroIDs, "987654321")

	mockRequest := ctrl.Request{
		ID:                    "987654321",
		LookingForGender:      2,
		Gender:                1,
		LookingForAgeMin:      25,
		LookingForAgeMax:      45,
		MaxDistance:           50,
		DistanceUnit:          "km",
		Lat:                   0.123456789,
		Lon:                   0.123456789,
		RetrievedSuperheroIDs: retrievedSuperheroIDs,
		IsESRequest:           true,
	}

	likeSuperheroIDs := make([]string, 0)
	likeSuperheroIDs = append(likeSuperheroIDs, "11111111111")

	_, _, err = mockService.HandleESRequest(mockRequest, likeSuperheroIDs)
	assert.NoError(t, err)
}
