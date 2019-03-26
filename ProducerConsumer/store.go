package main

import (
	"time"
)

// Product ...
type Product struct {
	Num int
}

// Store ...
type Store struct {
	products []Product
}

// Add ...
func (s *Store) Add(p Product) {
	if len(s.products) >= 2 {

		time.Sleep(time.Second * 1)

	}

	s.products = append(s.products, p)
}

// Get ...
func (s *Store) Get() *Product {
	for {
		if len(s.products) <= 0 {
			time.Sleep(time.Millisecond * 100)
			continue
		}

		v := &s.products[0]

		s.products = append(s.products[1:])

		return v
	}
}
