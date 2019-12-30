package service

import (
	"go.uber.org/zap"

	"github.com/superhero-suggestions/internal/cache"
	"github.com/superhero-suggestions/internal/config"
	"github.com/superhero-suggestions/internal/es"
)

// Service holds all the different services that are used when handling request.
type Service struct {
	ES         *es.ES
	Cache      *cache.Cache
	PageSize   int
	Logger     *zap.Logger
	TimeFormat string
}

// NewService creates value of type Service.
func NewService(cfg *config.Config) (*Service, error) {
	e, err := es.NewES(cfg)
	if err != nil {
		return nil, err
	}

	c, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	defer logger.Sync()

	return &Service{
		ES:         e,
		Cache:      c,
		PageSize:   cfg.App.PageSize,
		Logger:     logger,
		TimeFormat: cfg.App.TimeFormat,
	}, nil
}
