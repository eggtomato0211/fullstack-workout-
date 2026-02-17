package main

import (
	"errors"
	"fmt"
)

// User はコンストラクタパターンで作成される構造体
type User struct {
	name  string // unexported: 外部から直接変更不可
	email string
	age   int
}

// NewUser はUserのコンストラクタ関数
func NewUser(name, email string, age int) (*User, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if age < 0 || age > 150 {
		return nil, errors.New("age must be between 0 and 150")
	}
	return &User{name: name, email: email, age: age}, nil
}

// Getter メソッド
func (u *User) Name() string  { return u.name }
func (u *User) Email() string { return u.email }
func (u *User) Age() int      { return u.age }

func main() {
	// 正常ケース
	user, err := NewUser("田中太郎", "tanaka@example.com", 30)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Name: %s, Email: %s, Age: %d\n", user.Name(), user.Email(), user.Age())

	// バリデーションエラーのケース
	_, err = NewUser("", "test@example.com", 25)
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = NewUser("テスト", "test@example.com", -1)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
