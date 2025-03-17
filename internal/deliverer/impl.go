package deliverer

import "fmt"

type Impl struct {
	from   chan int
	target chan int
}

func (deliver *Impl) Deliver() {
	for val := range deliver.from {
		fmt.Printf("%10s: %3d\n", "deliver", val)
		deliver.target <- val
	}
}

func (deliver *Impl) SetDeliveryCh(ch chan int) {
	deliver.target = ch
}
func (deliver *Impl) SetStorageCh(ch chan int) {
	deliver.from = ch
}
