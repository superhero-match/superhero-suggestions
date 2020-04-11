/*
  Copyright (C) 2019 - 2020 MWSOFT
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
package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/superhero-match/superhero-suggestions/internal/cache"
	"github.com/superhero-match/superhero-suggestions/internal/config"
	"github.com/superhero-match/superhero-suggestions/internal/es/model"
	"strconv"
)

const (
	male   = int(1)
	female = int(2)
	both   = int(3)
)

// GetSuggestions fetches suggestions for the Superhero.
func (es *ES) GetSuggestions(req *model.Request) (superheros []model.Superhero, err error) {
	suggestionsQuery := elastic.NewBoolQuery()
	suggestionsQuery = suggestionsQuery.Must(elastic.NewMatchAllQuery())

	maxDistance := strconv.Itoa(req.MaxDistance)

	distanceQuery := elastic.NewGeoDistanceQuery("location")
	distanceQuery = distanceQuery.Lat(req.Lat)
	distanceQuery = distanceQuery.Lon(req.Lon)
	distanceQuery = distanceQuery.Distance(maxDistance + req.DistanceUnit)

	suggestionsQuery = suggestionsQuery.Filter(distanceQuery)

	userIDQuery := elastic.NewMatchQuery("superhero_id", req.ID)
	suggestionsQuery.MustNot(userIDQuery)

	if req.LookingForGender == both {
		maleGenderQuery := elastic.NewMatchQuery("gender", male)
		femaleGenderQuery := elastic.NewMatchQuery("gender", female)

		suggestionsQuery.Should(maleGenderQuery)
		suggestionsQuery.Should(femaleGenderQuery)
	} else {
		genderQuery := elastic.NewMatchQuery("gender", req.LookingForGender)
		suggestionsQuery.Must(genderQuery)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	ch, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	// Fetch superheroes ids who liked this user.
	likeSuperheroIDs, err := ch.GetLikes(req.ID)
	if err != nil {
		return nil, err
	}

	// Configure the query to include in result all users who liked this user.
	// The goal behind this is to make users match as quick as possible.
	// No need to be on the app for ages, just match as quick as possible and
	// go have fun.
	if likeSuperheroIDs != nil && len(likeSuperheroIDs) > 0 {
		idsToBeIncluded := make([]interface{}, len(likeSuperheroIDs))
		for index, value := range likeSuperheroIDs {
			idsToBeIncluded[index] = value
		}

		includeSuperherosQuery := elastic.NewTermsQuery(
			"superhero_id",
			idsToBeIncluded...,
		)

		suggestionsQuery.Must(includeSuperherosQuery)
	}

	// Delete the likes as they were already included in the Elasticsearch query.
	// No need to be fetching the same users over and over again.
	err = ch.DeleteLikes(req.ID)
	if err != nil {
		return nil, err
	}

	if len(req.RetrievedSuperheroIDs) > 0 {
		idsToBeExcluded := make([]interface{}, len(req.RetrievedSuperheroIDs))
		for index, value := range req.RetrievedSuperheroIDs {
			idsToBeExcluded[index] = value
		}

		excludeSuperherosQuery := elastic.NewTermsQuery(
			"superhero_id",
			idsToBeExcluded...,
		)

		suggestionsQuery.MustNot(excludeSuperherosQuery)
	}

	ageRangeQuery := elastic.NewBoolQuery().
		Filter(
			elastic.NewRangeQuery("age").
				From(req.LookingForAgeMin).
				To(req.LookingForAgeMax),
		)
	suggestionsQuery.Must(ageRangeQuery)

	src, err := suggestionsQuery.Source()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	got := string(data)

	fmt.Println(got)

	searchResult, err := es.Client.Search().
		Index(es.Index).
		Query(suggestionsQuery).
		Pretty(true).
		Size(es.BatchSize).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println()

	fmt.Printf("%+v", searchResult)

	fmt.Println()

	fmt.Println("searchResult.TotalHits()")
	fmt.Println(searchResult.TotalHits())

	for _, hit := range searchResult.Hits.Hits {
		var s model.Superhero

		err := json.Unmarshal(hit.Source, &s)
		if err != nil {
			return nil, err
		}

		fmt.Printf("%+v", s)

		superheros = append(superheros, s)
	}

	return superheros, nil
}
