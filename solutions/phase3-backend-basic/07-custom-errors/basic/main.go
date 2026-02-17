package main

import (
	"errors"
	"fmt"
)

// センチネルエラー
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
)

var users = map[int]string{
	1: "田中太郎",
	2: "鈴木花子",
}

func findUser(id int) (string, error) {
	name, ok := users[id]
	if !ok {
		return "", ErrNotFound
	}
	return name, nil
}

func main() {
	// センチネルエラーとの比較で分岐
	name, err := findUser(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found:", name)
	}

	_, err = findUser(99)
	if err == ErrNotFound {
		fmt.Println("ユーザーが見つかりません")
	} else if err == ErrUnauthorized {
		fmt.Println("権限がありません")
	}
}
