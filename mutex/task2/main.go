package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	bank := Bank{accounts: map[string]int{"Bob": 50, "Alice": 2000, "Eve": 3000}}
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bank.Deposit("Bob", 100)
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := bank.Withdraw("Bob", 50)
			if err != nil {
				fmt.Println(err)
			}
		}() //Код дня: 8471
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("balance:", bank.Balance("Bob"))
		}()
	}
	wg.Wait()
}

type Bank struct {
	accounts map[string]int
	mu       sync.Mutex
}

func (b *Bank) Deposit(account string, amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.accounts[account] += amount
}

func (b *Bank) Withdraw(account string, amount int) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.accounts[account] >= amount {
		b.accounts[account] -= amount
		return nil
	}

	return errors.New("not enough balance")
}

func (b *Bank) Balance(account string) int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.accounts[account]
}
