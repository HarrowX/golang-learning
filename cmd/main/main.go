package main

import (
	"gorutins-learning/internal/customer"
	"gorutins-learning/internal/deliverer"
	"gorutins-learning/internal/supplier"
	"sync"
)

func main() {
	//for waiting until client apply all delivers
	mainGroup := sync.WaitGroup{}
	mainGroup.Add(1)
	defer mainGroup.Wait()

	supplierInstance := supplier.NewSupplier()
	countDelivers := 5

	for i := 0; i < countDelivers; i++ {
		w := deliverer.NewDeliverer()
		supplierInstance.RegisterDelivery(&w)
	}

	customerInstance := customer.NewCustomer()

	supplierInstance.OfferWith(&customerInstance)

	go supplierInstance.Generate(1, 100)

	go supplierInstance.Supply()

	go func() {
		customerInstance.ApplyDelivery()
		mainGroup.Done()
	}()

}
