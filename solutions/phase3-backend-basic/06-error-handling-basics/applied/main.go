package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func validateUser(u User) error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Age < 0 || u.Age > 150 {
		return fmt.Errorf("invalid age: %d (must be 0-150)", u.Age)
	}
	return nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	// バリデーション
	user := User{Name: "", Email: "test@example.com", Age: 25}
	if err := validateUser(user); err != nil {
		fmt.Println("バリデーションエラー:", err)
	}

	user = User{Name: "田中太郎", Email: "tanaka@example.com", Age: 25}
	if err := validateUser(user); err != nil {
		fmt.Println("バリデーションエラー:", err)
	} else {
		fmt.Println("バリデーション成功")
	}

	// 除算
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("計算エラー:", err)
		return
	}
	fmt.Printf("10 / 3 = %.2f\n", result)

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("計算エラー:", err)
	}
}
