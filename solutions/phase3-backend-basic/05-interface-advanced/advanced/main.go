package main

import "fmt"

// PaymentMethod は支払い方法のインターフェース
type PaymentMethod interface {
	Pay(amount int) error
	Name() string
}

type CreditCard struct {
	Number string
	Limit  int
}

func (c CreditCard) Pay(amount int) error {
	if amount > c.Limit {
		return fmt.Errorf("credit limit exceeded: limit=%d, amount=%d", c.Limit, amount)
	}
	return nil
}

func (c CreditCard) Name() string { return "クレジットカード" }

type BankTransfer struct {
	AccountNumber string
	Balance       int
}

func (b BankTransfer) Pay(amount int) error {
	if amount > b.Balance {
		return fmt.Errorf("insufficient balance: balance=%d, amount=%d", b.Balance, amount)
	}
	return nil
}

func (b BankTransfer) Name() string { return "銀行振込" }

// ProcessPayment は型スイッチで支払い方法に応じたメッセージを返す
func ProcessPayment(method PaymentMethod, amount int) string {
	err := method.Pay(amount)
	if err != nil {
		return fmt.Sprintf("支払い失敗 (%s): %v", method.Name(), err)
	}

	switch v := method.(type) {
	case CreditCard:
		return fmt.Sprintf("カード(%s)で%d円支払い完了", v.Number[len(v.Number)-4:], amount)
	case BankTransfer:
		return fmt.Sprintf("口座(%s)から%d円振込完了", v.AccountNumber, amount)
	default:
		return fmt.Sprintf("%sで%d円支払い完了", method.Name(), amount)
	}
}

func main() {
	card := CreditCard{Number: "4111111111111111", Limit: 10000}
	bank := BankTransfer{AccountNumber: "123-456-789", Balance: 50000}

	fmt.Println(ProcessPayment(card, 5000))  // 成功
	fmt.Println(ProcessPayment(card, 15000)) // 限度額超過
	fmt.Println(ProcessPayment(bank, 30000)) // 成功
	fmt.Println(ProcessPayment(bank, 60000)) // 残高不足
}
