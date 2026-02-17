package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

// DB層
func findUserInDB(id int) (string, error) {
	if id == 99 {
		return "", ErrNotFound
	}
	return "田中太郎", nil
}

// リポジトリ層: %w でラップ
func getUserFromRepo(id int) (string, error) {
	name, err := findUserInDB(id)
	if err != nil {
		return "", fmt.Errorf("getUserFromRepo(id=%d): %w", id, err)
	}
	return name, nil
}

// サービス層: さらにラップ
func getUserService(id int) (string, error) {
	name, err := getUserFromRepo(id)
	if err != nil {
		return "", fmt.Errorf("user service: %w", err)
	}
	return name, nil
}

func main() {
	_, err := getUserService(99)
	if err != nil {
		fmt.Println("エラー:", err)

		// errors.Is でエラーチェーンをたどる
		if errors.Is(err, ErrNotFound) {
			fmt.Println("→ ユーザーが見つかりません（404を返す）")
		}
	}
}
