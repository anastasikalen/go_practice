package main

import (
	"errors"
	"fmt"
)

func main() {
	person1 := &Account{"Alice", 1000}
	person2 := &Account{"Bob", 2000}

	err := person1.Transfer(person2, 100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(person1)
	fmt.Println(person2)

} //Код дня: 1074

type Account struct {
	Owner   string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}
func (a *Account) Withdraw(amount float64) error {
	if a.Balance < amount {
		return errors.New("not enough balance")

	}
	a.Balance -= amount
	return nil
}
func (a *Account) Transfer(to *Account, amount float64) error {
	if a.Balance < amount {
		return errors.New("not enough balance")
	}
	err := a.Withdraw(amount)
	if err != nil {
		return err
	}
	to.Deposit(amount)
	return nil
}
