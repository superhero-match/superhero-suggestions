package es

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/superhero-suggestions/internal/es/model"

	"gopkg.in/olivere/elastic.v7"
)

// GetSuggestions fetches suggestions for the Superhero.
func (es *ES) GetSuggestions(req *model.Request) (shs []*model.Superhero, err error) {
	superheros := make([]*model.Superhero, 0)

	suggestionsQuery := elastic.NewBoolQuery()
	suggestionsQuery = suggestionsQuery.Must(elastic.NewMatchAllQuery())

	maxDistance := strconv.Itoa(req.MaxDistance)

	distanceQuery := elastic.NewGeoDistanceQuery("location")
	distanceQuery = distanceQuery.Lat(req.Lat)
	distanceQuery = distanceQuery.Lon(req.Lon)
	distanceQuery = distanceQuery.Distance(maxDistance + req.DistanceUnit)

	suggestionsQuery = suggestionsQuery.Filter(distanceQuery)

	genderQuery := elastic.NewMatchQuery("gender", req.LookingForGender)
	suggestionsQuery.Must(genderQuery)

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
		From(req.Offset).
		Size(req.Size).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("%+v", searchResult)

	fmt.Println()

	fmt.Println(searchResult.TotalHits())

	for _, hit := range searchResult.Hits.Hits {
		var s model.Superhero

		err := json.Unmarshal(hit.Source, &s)
		if err != nil {
			return nil, err
		}

		fmt.Printf("%+v", s)

		superheros = append(superheros, &s)
	}

	return superheros, nil
}
