package es

import (
	"fmt"

	"github.com/superhero-suggestions/internal/config"

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
