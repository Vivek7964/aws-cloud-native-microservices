package repository

import (
	"github.com/Vivek7964/microservice-demo/payments/model"
)

type inMemoryRepository struct {
	paymentIntents map[string]*model.PaymentIntent
}

func newInMemoryRepository() Repository {
	return &inMemoryRepository{
		paymentIntents: make(map[string]*model.PaymentIntent),
	}
}

func (r *inMemoryRepository) GetPaymentIntent(cartID string) (*model.PaymentIntent, error) {
	if paymentIntent, ok := r.paymentIntents[cartID]; ok {
		return paymentIntent, nil
	}

	return &model.PaymentIntent{}, nil
}