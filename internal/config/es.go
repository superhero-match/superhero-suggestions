package config

// ES holds the configuration values for the Elasticsearch client.
type ES struct {
	Host     string `env:"ES_HOST" default:"localhost"`
	Port     int    `env:"ES_PORT" default:"9200"`
	Cluster string `env:"ES_CLUSTER" default:"superhero_match"`
	Index string `env:"ES_INDEX" default:"superheros_index"`
	Type string `env:"ES_DOCUMENT_TYPE" default:"superhero"`
}