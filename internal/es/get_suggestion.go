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

	"github.com/superhero-suggestions/internal/es/model"

	"gopkg.in/olivere/elastic.v7"
)

// GetSuggestion retrieves single Superhero.
// This method is going to be used to fetch a single Superhero when displaying
// suggestion profile. This method will be only called if that suggestion is not going to be
// found in Cache.
func (es *ES) GetSuggestion(superheroID string) (superhero *model.Superhero, err error) {
	fmt.Println(superheroID)
	fmt.Println(es.Index)

	q := elastic.NewTermQuery("superheroID", superheroID)

	fmt.Println()
	fmt.Printf("%+v", q)
	fmt.Println()

	searchResult, err := es.Client.Search().
		Index(es.Index).
		Query(q).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	fmt.Printf("SearchResult: %+v", searchResult)

	fmt.Println()

	fmt.Println(searchResult.TotalHits())

	if searchResult.TotalHits() > 0 {
		for _, hit := range searchResult.Hits.Hits {
			fmt.Printf("Hit: %+v", hit)

			err := json.Unmarshal(hit.Source, &superhero)
			if err != nil {
				return nil, err
			}

			fmt.Println()
			fmt.Printf("Superhero Unmarshalled: %+v", &superhero)
			fmt.Println()
		}
	}

	return superhero, nil
}