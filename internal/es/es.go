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
	"fmt"

	"github.com/superhero-match/superhero-suggestions/internal/config"

	"gopkg.in/olivere/elastic.v7"
)

// ES holds all the Elasticsearch client relevant data.
type ES struct {
	Client    *elastic.Client
	Host      string
	Port      string
	Cluster   string
	Index     string
	BatchSize int
}

// NewES creates a client connection to Elasticsearch.
func NewES(cfg *config.Config) (es *ES, err error) {
	client, err := elastic.NewClient(
		elastic.SetURL(
			fmt.Sprintf(
				"http://%s:%s",
				cfg.ES.Host,
				cfg.ES.Port,
			),
		),
	)
	if err != nil {
		return nil, err
	}

	return &ES{
		Client:    client,
		Host:      cfg.ES.Host,
		Port:      cfg.ES.Port,
		Cluster:   cfg.ES.Cluster,
		Index:     cfg.ES.Index,
		BatchSize: cfg.ES.BatchSize,
	}, nil
}
