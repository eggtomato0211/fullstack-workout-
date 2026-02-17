package main

import "fmt"

// Notifier は通知を送るインターフェース
type Notifier interface {
	Notify(message string) error
}

type EmailNotifier struct {
	To string
}

func (n EmailNotifier) Notify(message string) error {
	fmt.Printf("メール送信: %s - %s\n", n.To, message)
	return nil
}

type SlackNotifier struct {
	Channel string
}

func (n SlackNotifier) Notify(message string) error {
	fmt.Printf("Slack送信: %s - %s\n", n.Channel, message)
	return nil
}

// SendAll は全Notifierにメッセージを送信する
func SendAll(notifiers []Notifier, message string) []error {
	var errs []error
	for _, n := range notifiers {
		if err := n.Notify(message); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func main() {
	notifiers := []Notifier{
		EmailNotifier{To: "user@example.com"},
		EmailNotifier{To: "admin@example.com"},
		SlackNotifier{Channel: "#general"},
		SlackNotifier{Channel: "#alerts"},
	}

	errs := SendAll(notifiers, "サーバーデプロイ完了")
	if len(errs) > 0 {
		fmt.Println("エラーあり:", errs)
	} else {
		fmt.Println("全通知成功")
	}
}
