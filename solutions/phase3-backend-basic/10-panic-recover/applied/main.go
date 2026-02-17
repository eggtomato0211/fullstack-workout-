package main

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL string
	Port        int
}

// Must で始まる関数はpanic可能性を示す慣習
func MustLoadConfig() *Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		panic("DATABASE_URL environment variable is required")
	}
	return &Config{DatabaseURL: dbURL, Port: 8080}
}

// panicが不適切な場面: 通常のerrorで返す
func findUser(id int) (string, error) {
	users := map[int]string{1: "田中太郎"}
	name, ok := users[id]
	if !ok {
		return "", fmt.Errorf("user not found: id=%d", id)
	}
	return name, nil
}

func main() {
	// panicが不適切な場面: errorで処理
	_, err := findUser(99)
	if err != nil {
		fmt.Println("通常のエラー:", err)
	}

	// panicが適切な場面: 起動時の必須設定チェック
	fmt.Println("設定を読み込みます...")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("起動エラー:", r)
		}
	}()

	_ = MustLoadConfig()
}
