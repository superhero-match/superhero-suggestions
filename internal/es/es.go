package es

import (
	"fmt"

	"github.com/superhero-suggestions/internal/config"
	elastic "gopkg.in/olivere/elastic.v7"
)

// ES holds all the Elasticsearch client relevant data.
type ES struct {
	ES *elastic.Client
}

// NewES creates a client and connects to it.
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
		ES: client,
	}, nil
}