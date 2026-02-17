package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance int
}

func (a *BankAccount) Deposit(amount int) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.Balance += amount
	return nil
}

func (a *BankAccount) Withdraw(amount int) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if amount > a.Balance {
		return fmt.Errorf("insufficient funds: have %d, want %d", a.Balance, amount)
	}
	a.Balance -= amount
	return nil
}

func (a *BankAccount) Transfer(to *BankAccount, amount int) error {
	if err := a.Withdraw(amount); err != nil {
		return fmt.Errorf("transfer failed: %w", err)
	}
	if err := to.Deposit(amount); err != nil {
		// 出金を戻す
		a.Balance += amount
		return fmt.Errorf("transfer failed: %w", err)
	}
	return nil
}

func main() {
	alice := &BankAccount{Owner: "Alice", Balance: 1000}
	bob := &BankAccount{Owner: "Bob", Balance: 500}

	fmt.Printf("初期: Alice=%d, Bob=%d\n", alice.Balance, bob.Balance)

	// 入金
	alice.Deposit(500)
	fmt.Printf("Alice入金後: %d\n", alice.Balance)

	// 出金
	alice.Withdraw(200)
	fmt.Printf("Alice出金後: %d\n", alice.Balance)

	// 残高不足
	err := alice.Withdraw(10000)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 送金
	err = alice.Transfer(bob, 300)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("送金後: Alice=%d, Bob=%d\n", alice.Balance, bob.Balance)
}
