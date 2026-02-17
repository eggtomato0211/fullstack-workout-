package main

import "fmt"

// ValidationError はバリデーションエラーの詳細を持つカスタム型
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

// NotFoundError はリソースが見つからないエラー
type NotFoundError struct {
	Resource string
	ID       int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found: id=%d", e.Resource, e.ID)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "must be non-negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: fmt.Sprintf("must be <= 150, got %d", age)}
	}
	return nil
}

func findProduct(id int) (string, error) {
	products := map[int]string{1: "Go入門書", 2: "キーボード"}
	name, ok := products[id]
	if !ok {
		return "", &NotFoundError{Resource: "product", ID: id}
	}
	return name, nil
}

func main() {
	// 型アサーションでカスタムエラーの詳細にアクセス
	err := validateAge(-5)
	if ve, ok := err.(*ValidationError); ok {
		fmt.Printf("フィールド「%s」のエラー: %s\n", ve.Field, ve.Message)
	}

	_, err = findProduct(99)
	if nfe, ok := err.(*NotFoundError); ok {
		fmt.Printf("%sが見つかりません (ID: %d)\n", nfe.Resource, nfe.ID)
	}
}
