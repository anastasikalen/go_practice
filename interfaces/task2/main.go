package main

import (
	"errors"
	"fmt"
)

func main() {
	creditCard := CreditCard{
		CardNumber: 99999999999,
		Balance:    123,
	}
	payPal := PayPal{
		Email:   "test@mail.ru",
		Balance: 1000,
	}
	bankTransfer := BankTransfer{
		AccountNumber: 1111111,
		Balance:       10,
	}
	err := MakePayment(&creditCard, 120)
	if err != nil {
		fmt.Println(err)
	}
	err = MakePayment(&payPal, 1000)
	if err != nil {
		fmt.Println(err)
	}
	err = MakePayment(&bankTransfer, 11)
	if err != nil {
		fmt.Println(err)
	}

}
func MakePayment(pm PaymentMethod, amount float64) error {
	if err := pm.Pay(amount); err != nil {
		return err
	}
	return nil
}

type PaymentMethod interface {
	Pay(amount float64) error
} //КОД ДНЯ 1154

type CreditCard struct {
	CardNumber int
	Balance    float64
}

func (c *CreditCard) Pay(amount float64) error {
	if amount > c.Balance {
		return errors.New("you cannot pay")
	}
	c.Balance = c.Balance - amount
	fmt.Printf("You have %f %d\n", c.Balance, c.CardNumber)
	return nil
}

type PayPal struct {
	Email   string
	Balance float64
}

func (p *PayPal) Pay(amount float64) error {
	if amount > p.Balance {
		return errors.New("you cannot pay")
	}
	p.Balance = p.Balance - amount
	fmt.Printf("You have %f %s\n", p.Balance, p.Email)
	return nil
}

type BankTransfer struct {
	AccountNumber int
	Balance       float64
}

func (b *BankTransfer) Pay(amount float64) error {
	if amount > b.Balance {
		return errors.New("you cannot pay")
	}
	b.Balance = b.Balance - amount
	fmt.Printf("You have %f %d\n", b.Balance, b.AccountNumber)
	return nil
}
