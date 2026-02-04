package main

import (
	"fmt"
	"sync"
)

func main() {
	c := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(c *int) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			*c++
		}(&c)
	}
	wg.Wait() //Код дня: 8471

	fmt.Println(c)

}
