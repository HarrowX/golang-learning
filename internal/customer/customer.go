package customer

type Customer interface {
	ApplyDelivery()
	SetReceiveCh(ch chan int)
}

func NewCustomer() Customer {
	return &Impl{}
}
