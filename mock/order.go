package mock

import (
	cart "github.com/marcusolsson/coverage-cravings"
)

type OrderRepository struct {
	FindByIDFunc func(id string) (*cart.Order, error)
	SaveFunc     func(o *cart.Order) error
}

func (r *OrderRepository) Save(o *cart.Order) error {
	return r.SaveFunc(o)
}

func (r *OrderRepository) FindByID(id string) (*cart.Order, error) {
	return r.FindByIDFunc(id)
}
