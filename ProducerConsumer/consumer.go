package main

import (
	"fmt"
)

// Consumer ...
type Consumer struct {
	n     int
	store *Store
}

// Consume ...
func (c *Consumer) Consume() {
	fmt.Printf("Consumer %d Starts Consume\n", c.n)
	for {

		v := c.store.Get()

		fmt.Printf("Consumer %d consume %d\n", c.n, v.Num)
	}
}
