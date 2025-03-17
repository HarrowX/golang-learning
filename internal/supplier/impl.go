package supplier

import (
	"fmt"
	"gorutins-learning/internal/customer"
	"gorutins-learning/internal/deliverer"
	"sync"
)

type Impl struct {
	storageCh  chan int
	customerCh chan int
	workers    []*deliverer.Deliverer
}

func (brigadier *Impl) Generate(a, b int) {
	for i := a; i <= b; i++ {
		fmt.Printf("%10s: %3d\n", "brig", i)
		brigadier.storageCh <- i
	}
	close(brigadier.storageCh)
}

func (brigadier *Impl) prepareDelivers() {
	for _, w := range brigadier.workers {
		(*w).SetStorageCh(brigadier.storageCh)
		(*w).SetDeliveryCh(brigadier.customerCh)
	}
}

func (brigadier *Impl) Supply() {

	defer close(brigadier.customerCh)

	workersWaitGroup := sync.WaitGroup{}
	defer workersWaitGroup.Wait()
	workersWaitGroup.Add(len(brigadier.workers))

	brigadier.prepareDelivers()

	for _, w := range brigadier.workers {
		go func() {
			defer workersWaitGroup.Done()
			(*w).Deliver()
		}()
	}
}

func (brigadier *Impl) RegisterDelivery(deliver *deliverer.Deliverer) {
	brigadier.workers = append(brigadier.workers, deliver)
}

func (brigadier *Impl) OfferWith(owner *customer.Customer) {
	brigadier.storageCh = make(chan int)
	brigadier.customerCh = make(chan int)
	(*owner).SetReceiveCh(brigadier.customerCh)
}
