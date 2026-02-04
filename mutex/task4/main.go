package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	texts := []string{
		"go is awesome",
		"concurrency is powerful",
		"go routines and channels",
		"mutexes prevent data races",
	}
	var mu sync.Mutex
	var wg sync.WaitGroup
	m := make(map[string]int)

	for _, text := range texts {
		t := strings.TrimSpace(text)
		for _, word := range strings.Split(t, " ") {
			wg.Add(1)

			w := word
			go func(w string) {
				mu.Lock() //Код дня: 8471
				m[w]++
				mu.Unlock()
				wg.Done()
			}(w)
		}

	}
	wg.Wait()

	fmt.Println(m)
}
