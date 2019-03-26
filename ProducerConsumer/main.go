package main

import "time"

func main() {
	s := new(Store)

	for i := 0; i < 3; i++ {
		p := new(Producer)
		p.store = s
		p.n = i
		go func() {
			p.Produce()
		}()
	}

	for i := 0; i < 2; i++ {
		c := new(Consumer)
		c.store = s
		c.n = i
		go func() {
			c.Consume()
		}()
	}

	time.Sleep(time.Hour * 1)
}
