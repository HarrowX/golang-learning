package customer

import "fmt"

type Impl struct {
	from chan int
}

func (client *Impl) ApplyDelivery() {
	for val := range client.from {
		fmt.Printf("%10s: %3d\n", "client", val)
	}
}
func (client *Impl) SetReceiveCh(ch chan int) {
	client.from = ch
}
