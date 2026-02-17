package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation: %s - %s", e.Field, e.Message)
}

type DBError struct {
	Query   string
	Message string
}

func (e *DBError) Error() string {
	return fmt.Sprintf("db error: %s (%s)", e.Message, e.Query)
}

func saveUser(email string) error {
	if email == "" {
		return fmt.Errorf("saveUser: %w", &ValidationError{
			Field: "email", Message: "is required",
		})
	}
	return fmt.Errorf("saveUser: %w", &DBError{
		Query: "INSERT INTO users", Message: "duplicate key",
	})
}

func main() {
	// バリデーションエラー
	err := saveUser("")
	if err != nil {
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("バリデーションエラー: フィールド=%s, メッセージ=%s\n", ve.Field, ve.Message)
		}
	}

	// DBエラー
	err = saveUser("test@example.com")
	if err != nil {
		var de *DBError
		if errors.As(err, &de) {
			fmt.Printf("DBエラー: クエリ=%s, メッセージ=%s\n", de.Query, de.Message)
		}
	}
}
