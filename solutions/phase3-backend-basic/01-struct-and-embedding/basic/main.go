package main

import "fmt"

// User は利用者を表す構造体
type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	// 初期化パターン1: フィールド名を指定（推奨）
	u1 := User{
		Name:  "田中太郎",
		Email: "tanaka@example.com",
		Age:   30,
	}

	// 初期化パターン2: ゼロ値で初期化してからフィールドを設定
	var u2 User
	u2.Name = "鈴木花子"
	u2.Email = "suzuki@example.com"

	// 初期化パターン3: new()でポインタを取得
	u3 := new(User)
	u3.Name = "佐藤一郎"

	fmt.Printf("u1: %+v\n", u1)
	fmt.Printf("u2: %+v\n", u2)
	fmt.Printf("u3: %+v\n", *u3)
}
