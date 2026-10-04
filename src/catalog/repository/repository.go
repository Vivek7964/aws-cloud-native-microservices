package repository

import (
	"fmt"

	"github.com/Vivek7964/microservice-demo/catalog/config"
	"github.com/Vivek7964/microservice-demo/catalog/model"
	"github.com/prometheus/client_golang/prometheus"
)

type Repository interface {
	List(tags []string, order string, pageNum, pageSize int) ([]model.Product, error)
	Count(tags []string) (int, error)
	Get(id string) (*model.Product, error)
	Tags() ([]model.Tag, error)
	Collector() prometheus.Collector
	ReaderCollector() prometheus.Collector
}

func NewRepository(config config.DatabaseConfiguration) (Repository, error) {
	if config.Type == "json" {
		return newJSONRepository()
	}

	return nil, fmt.Errorf("Unknown database type: %s", config.Type)
}

func contains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}

	return false
}