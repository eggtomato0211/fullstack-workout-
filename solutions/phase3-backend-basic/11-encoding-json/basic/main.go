package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	// Marshal: 構造体 → JSON
	user := User{ID: 1, Name: "田中太郎", Email: "tanaka@example.com", Age: 30}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Println("JSON:", string(jsonBytes))

	// MarshalIndent: 整形JSON
	prettyJSON, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("Pretty JSON:")
	fmt.Println(string(prettyJSON))

	// Unmarshal: JSON → 構造体
	jsonStr := `{"id":2,"name":"鈴木花子","email":"suzuki@example.com","age":25}`
	var decoded User
	if err := json.Unmarshal([]byte(jsonStr), &decoded); err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	fmt.Printf("Decoded: %+v\n", decoded)
}
