package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

// AppError はUnwrapを実装したカスタムエラー
type AppError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %v", e.StatusCode, e.Message, e.Err)
	}
	return fmt.Sprintf("[%d] %s", e.StatusCode, e.Message)
}

func (e *AppError) Unwrap() error { return e.Err }

// DB層
func findOrderInDB(id int) (string, error) {
	if id == 0 {
		return "", ErrNotFound
	}
	return "注文#" + fmt.Sprint(id), nil
}

// リポジトリ層
func getOrder(id int) (string, error) {
	order, err := findOrderInDB(id)
	if err != nil {
		return "", fmt.Errorf("order repository: %w", err)
	}
	return order, nil
}

// サービス層: エラーをAppErrorに変換
func processOrder(id int) error {
	order, err := getOrder(id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return &AppError{StatusCode: 404, Message: "order not found", Err: err}
		}
		return &AppError{StatusCode: 500, Message: "internal error", Err: err}
	}
	fmt.Println("処理完了:", order)
	return nil
}

// ハンドラー層
func handleOrderRequest(id int) {
	err := processOrder(id)
	if err != nil {
		var appErr *AppError
		if errors.As(err, &appErr) {
			fmt.Printf("HTTP %d: %s\n", appErr.StatusCode, appErr.Message)
		}
		if errors.Is(err, ErrNotFound) {
			fmt.Println("→ リソースが見つかりません")
		}
		return
	}
}

func main() {
	handleOrderRequest(1)
	fmt.Println("---")
	handleOrderRequest(0)
}
