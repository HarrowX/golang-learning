package deliverer

type Deliverer interface {
	Deliver()
	SetDeliveryCh(ch chan int)
	SetStorageCh(ch chan int)
}

func NewDeliverer() Deliverer {
	deliver := Impl{}
	return &deliver
}
