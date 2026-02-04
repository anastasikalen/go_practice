package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	bank := Bank{accounts: map[string]int{"Bob": 50, "Alice": 2000, "Eve": 3000}}
	var wg sync.WaitGroup

	accounts := []string{"Bob", "Alice", "Eve"}
	op := randomOperations(accounts, 20)

	for _, op := range op { //Код дня: 8471
		wg.Add(1)
		go func(op Operation) {
			defer wg.Done()
			if op.Type == "deposit" {
				bank.Deposit(op.Account, op.Amount)
				fmt.Printf("Deposited %d to %s\n", op.Amount, op.Account)
			}
			if op.Type == "withdrawal" {
				err := bank.Withdraw(op.Account, op.Amount)
				if err != nil {
					fmt.Printf("Failed to withdraw %d from %s: %s\n", op.Amount, op.Account, err)
				} else {
					fmt.Printf("Withdrew %d from %s\n", op.Amount, op.Account)
				}
			}
		}(op)
	}
	wg.Wait()

	fmt.Println("\n Final balances:")
	for _, acc := range accounts {
		fmt.Printf("%s: %d\n", acc, bank.Balance(acc))
	}
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

type Operation struct {
	Account string
	Amount  int
	Type    string
}

func randomOperations(accounts []string, n int) []Operation {
	rand.Seed(time.Now().UnixNano())
	ops := make([]Operation, n)
	for i := 0; i < n; i++ {
		account := accounts[rand.Intn(len(accounts))]
		amount := rand.Intn(500) + 1
		opType := "deposit"
		if rand.Intn(2) == 0 {
			opType = "withdraw"
		}
		ops[i] = Operation{Account: account, Amount: amount, Type: opType}
	}
	return ops
}
