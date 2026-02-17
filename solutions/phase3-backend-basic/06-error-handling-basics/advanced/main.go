package main

import (
	"fmt"
	"strings"
)

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

type User struct {
	ID       int
	Username string
	Email    string
}

var nextID = 1

func ValidateRequest(req RegisterRequest) error {
	if req.Username == "" {
		return fmt.Errorf("username is required")
	}
	if len(req.Username) < 3 {
		return fmt.Errorf("username must be at least 3 characters")
	}
	if !strings.Contains(req.Email, "@") {
		return fmt.Errorf("email must contain @")
	}
	if len(req.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}
	return nil
}

func RegisterUser(req RegisterRequest) (*User, error) {
	if err := ValidateRequest(req); err != nil {
		return nil, fmt.Errorf("validation failed: %v", err)
	}

	user := &User{
		ID:       nextID,
		Username: req.Username,
		Email:    req.Email,
	}
	nextID++
	return user, nil
}

func main() {
	// 正常ケース
	user, err := RegisterUser(RegisterRequest{
		Username: "tanaka",
		Email:    "tanaka@example.com",
		Password: "password123",
	})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("登録成功: %+v\n", user)
	}

	// エラーケース
	testCases := []RegisterRequest{
		{Username: "ab", Email: "test@example.com", Password: "password123"},
		{Username: "tanaka", Email: "invalid", Password: "password123"},
		{Username: "tanaka", Email: "test@example.com", Password: "short"},
	}

	for _, tc := range testCases {
		_, err := RegisterUser(tc)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
