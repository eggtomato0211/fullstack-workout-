# 11. encoding/jsonと構造体タグ - JSON Marshal/Unmarshal、構造体タグの活用

## 🎯 このテーマで学ぶこと

- `json.Marshal`/`json.Unmarshal`の基本
- 構造体タグによるフィールド名のカスタマイズ
- `omitempty`、`json:"-"`オプションの活用

## 📖 なぜencoding/jsonを理解する必要があるのか

Web APIでは、リクエストの受信（JSON→構造体）とレスポンスの送信（構造体→JSON）が全ての基本です。GoのフィールドはPascalCase（`UserName`）ですが、JSONは通常snake_case（`user_name`）やcamelCase（`userName`）を使います。構造体タグがこのギャップを埋めます。

### こう書かないとどうなるか

```go
// タグなしの構造体
type User struct {
    UserName string
    Email    string
}

// JSON出力: {"UserName":"田中","Email":"..."} ← PascalCaseのまま
// フロントエンドは通常 user_name や userName を期待する → 不一致

// タグありの構造体
type User struct {
    UserName string `json:"user_name"`
    Email    string `json:"email"`
}
// JSON出力: {"user_name":"田中","email":"..."} ← 期待通り
```

また、小文字で始まるフィールド（unexported）はJSON変換されません。これはGoの可視性ルールに従っています。

## 💡 コード例

### 基本: Marshal/Unmarshalと構造体タグ

```go
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
	role  string // 小文字 → JSON変換されない（unexported）
}

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description,omitempty"` // 空文字→省略
	Discount    int    `json:"discount,omitempty"`    // 0→省略
	InternalKey string `json:"-"`                     // JSONに一切含めない
	Stock       *int   `json:"stock,omitempty"`       // nil→省略
}

func intPtr(i int) *int { return &i }

func main() {
	// --- Marshal: 構造体 → JSON ---
	user := User{
		ID: 1, Name: "田中太郎", Email: "tanaka@example.com",
		Age: 30, role: "admin", // roleはJSONに含まれない
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Println("JSON:", string(jsonBytes))
	// → {"id":1,"name":"田中太郎","email":"tanaka@example.com","age":30}

	// MarshalIndent → 整形されたJSON（デバッグ用）
	prettyJSON, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("Pretty JSON:")
	fmt.Println(string(prettyJSON))

	// --- Unmarshal: JSON → 構造体 ---
	jsonStr := `{"id":2,"name":"鈴木花子","email":"suzuki@example.com","age":25}`

	var decoded User
	if err := json.Unmarshal([]byte(jsonStr), &decoded); err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	fmt.Printf("Decoded: %+v\n", decoded)

	// --- omitempty の動作 ---
	fmt.Println("\n--- omitempty ---")
	products := []Product{
		{
			ID: 1, Name: "Go入門書", Price: 3000,
			Description: "Goの基礎を学ぶ本", Discount: 500,
			InternalKey: "INTERNAL_001", Stock: intPtr(10),
		},
		{
			ID: 2, Name: "キーボード", Price: 15000,
			Description: "", Discount: 0,   // omitempty → 省略される
			InternalKey: "INTERNAL_002", Stock: nil, // omitempty → 省略される
		},
	}

	for _, p := range products {
		j, _ := json.MarshalIndent(p, "", "  ")
		fmt.Println(string(j))
		fmt.Println("---")
	}
}
```

### 実践: APIレスポンスの設計パターン

実務でのAPIレスポンス構造体の設計パターンです。ネストした構造体やスライスのJSON変換を学びます。

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// --- APIレスポンスの共通構造 ---
// 成功時はDataにデータ、失敗時はErrorにエラー情報を入れる
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// --- ドメインモデル ---
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Author    Author    `json:"author"`         // ネストした構造体
	Tags      []string  `json:"tags,omitempty"`  // スライス
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	now := time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC)

	// --- 成功レスポンス ---
	article := Article{
		ID: 1, Title: "Go入門", Body: "Goの基礎を学びましょう",
		Author:    Author{ID: 1, Name: "田中太郎"},
		Tags:      []string{"Go", "プログラミング", "入門"},
		CreatedAt: now,
	}
	successResp := APIResponse{Success: true, Data: article}

	jsonBytes, _ := json.MarshalIndent(successResp, "", "  ")
	fmt.Println("=== 成功レスポンス ===")
	fmt.Println(string(jsonBytes))

	// --- エラーレスポンス ---
	errorResp := APIResponse{
		Success: false,
		Error:   &ErrorInfo{Code: "NOT_FOUND", Message: "article not found"},
	}

	jsonBytes, _ = json.MarshalIndent(errorResp, "", "  ")
	fmt.Println("\n=== エラーレスポンス ===")
	fmt.Println(string(jsonBytes))

	// --- JSONリクエストのパース ---
	requestJSON := `{
		"title": "新しい記事",
		"body": "記事の本文です",
		"tags": ["Go", "Tutorial"]
	}`

	var newArticle Article
	if err := json.Unmarshal([]byte(requestJSON), &newArticle); err != nil {
		fmt.Println("Parse error:", err)
		return
	}
	fmt.Printf("\n=== パース結果 ===\n%+v\n", newArticle)
}
```

## 🎯 演習問題

TODOアプリのAPIレスポンスをJSON対応で設計してください。

**要件:**

1. `Todo`構造体: `ID int`, `Title string`, `Done bool`, `DueDate *time.Time`（nilなら期限なし）に構造体タグを付ける
2. `TodoList`構造体: `Todos []Todo`, `Count int`を持つ
3. `ToJSON(v interface{}) (string, error)`: 任意の値をインデント付きJSONに変換
4. `FromJSON(jsonStr string, v interface{}) error`: JSON文字列を任意の構造体にパース
5. DueDateがnilの場合はJSONから省略される（`omitempty`）

**期待される動作:**

- `Todo{Title: "Go学習", Done: false}` → `{"id":0,"title":"Go学習","done":false}`（DueDateは省略）
- JSONの`"done": true`が正しくパースされる

## ✅ 重要ポイント

- [ ] 構造体タグ（`` `json:"field_name"` ``）でJSONフィールド名を制御する
- [ ] `omitempty`でゼロ値のフィールドをJSON出力から省略する
- [ ] `json:"-"`でフィールドをJSON変換から除外する
- [ ] unexported（小文字始まり）フィールドはJSON変換されない

**次のテーマ:** [12. io.Reader/io.Writerの理解](./12-io-reader-writer.md)
