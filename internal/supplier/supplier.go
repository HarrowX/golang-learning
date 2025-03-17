package supplier

import (
	"gorutins-learning/internal/customer"
	"gorutins-learning/internal/deliverer"
)

type Supplier interface {
	Generate(a, b int)
	Supply()
	OfferWith(client *customer.Customer)
	RegisterDelivery(deliver *deliverer.Deliverer)
}

func NewSupplier() Supplier {
	brigadier := Impl{nil, nil, make([]*deliverer.Deliverer, 0)}
	return &brigadier
}
