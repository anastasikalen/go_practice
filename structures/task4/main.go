package main

import "fmt"

func main() {
	ow := BankAccount{
		Owner:   "Alice",
		Balance: 100,
	}

	ow.Deposit(50)
	fmt.Println(ow.Display())
	err := ow.Withdraw(50)
	if err != nil {
		fmt.Println(err)
	}
	ow.Display()
	err = ow.Withdraw(150)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ow.Display())

}

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount //КОД ДНЯ 9426
}
func (b *BankAccount) Withdraw(amount float64) error {
	if amount > b.Balance {
		return fmt.Errorf("not enough balance")
	}
	b.Balance -= amount
	return nil
}
func (b BankAccount) Display() string {
	return fmt.Sprintf("Owner: %s (Balance: %f)", b.Owner, b.Balance)
}
