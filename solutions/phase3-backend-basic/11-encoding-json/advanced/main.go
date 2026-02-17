package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Todo struct {
	ID      int        `json:"id"`
	Title   string     `json:"title"`
	Done    bool       `json:"done"`
	DueDate *time.Time `json:"due_date,omitempty"`
}

type TodoList struct {
	Todos []Todo `json:"todos"`
	Count int    `json:"count"`
}

func ToJSON(v interface{}) (string, error) {
	jsonBytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal error: %w", err)
	}
	return string(jsonBytes), nil
}

func FromJSON(jsonStr string, v interface{}) error {
	if err := json.Unmarshal([]byte(jsonStr), v); err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}
	return nil
}

func main() {
	due := time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC)

	list := TodoList{
		Todos: []Todo{
			{ID: 1, Title: "Go学習", Done: false, DueDate: &due},
			{ID: 2, Title: "テスト作成", Done: true},  // DueDateなし → 省略
			{ID: 3, Title: "デプロイ", Done: false},   // DueDateなし → 省略
		},
		Count: 3,
	}

	// 構造体 → JSON
	jsonStr, err := ToJSON(list)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("=== TodoList JSON ===")
	fmt.Println(jsonStr)

	// JSON → 構造体
	fmt.Println("\n=== パース結果 ===")
	var parsed TodoList
	if err := FromJSON(jsonStr, &parsed); err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, todo := range parsed.Todos {
		status := "未完了"
		if todo.Done {
			status = "完了"
		}
		dueStr := "なし"
		if todo.DueDate != nil {
			dueStr = todo.DueDate.Format("2006-01-02")
		}
		fmt.Printf("[%s] %s (期限: %s)\n", status, todo.Title, dueStr)
	}
}
