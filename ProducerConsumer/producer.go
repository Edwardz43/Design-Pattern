package main

import (
	"fmt"
	"math/rand"
)

// Producer ...
type Producer struct {
	n     int
	store *Store
}

// Produce ...
func (p *Producer) Produce() {
	fmt.Printf("Producer %d Starts Produce\n", p.n)
	for {
		r := rand.Intn(3000)
		// time.Sleep(time.Duration(r) * time.Millisecond)

		e := new(Product)

		e.Num = r

		p.store.Add(*e)

		fmt.Printf("Producer %d produce %d\n", p.n, e.Num)
	}
}
