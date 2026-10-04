package repository

import (
	"fmt"

	"github.com/Vivek7964/microservice-demo/payments/config"
	"github.com/Vivek7964/microservice-demo/payments/model"
)

type Repository interface {
	GetPaymentIntent(cartID string) (*model.PaymentIntent, error)
}

func NewRepository(config config.DatabaseConfiguration) (Repository, error) {
	if config.Type == "inmemory" {
		return newInMemoryRepository(), nil
	}

	return nil, fmt.Errorf("Unknown repository type: %s", config.Type)
}