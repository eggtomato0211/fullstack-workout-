package main

import (
	"errors"
	"fmt"
)

type Task struct {
	Title    string
	Done     bool
	Priority int
}

func NewTask(title string, priority int) (*Task, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}
	if priority < 1 || priority > 5 {
		return nil, fmt.Errorf("priority must be 1-5, got %d", priority)
	}
	return &Task{Title: title, Priority: priority}, nil
}

// ポインタレシーバー: 状態を変更する
func (t *Task) Complete() {
	t.Done = true
}

// 値レシーバー: 読み取り専用
func (t Task) IsHighPriority() bool {
	return t.Priority >= 4
}

func (t Task) String() string {
	check := " "
	if t.Done {
		check = "x"
	}
	return fmt.Sprintf("[%s] %s (優先度:%d)", check, t.Title, t.Priority)
}

func main() {
	task, err := NewTask("Go学習", 3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(task)
	fmt.Println("高優先度:", task.IsHighPriority())

	task.Complete()
	fmt.Println(task)

	// 高優先度タスク
	urgent, _ := NewTask("本番障害対応", 5)
	fmt.Println(urgent)
	fmt.Println("高優先度:", urgent.IsHighPriority())

	// バリデーションエラー
	_, err = NewTask("", 3)
	fmt.Println("Error:", err)

	_, err = NewTask("テスト", 6)
	fmt.Println("Error:", err)
}
