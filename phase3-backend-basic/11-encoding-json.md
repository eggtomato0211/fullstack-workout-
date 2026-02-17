# 11. encoding/jsonと構造体タグ - JSON Marshal/Unmarshal、構造体タグの活用

## 🎯 このテーマで学ぶこと

- `json.Marshal`/`json.Unmarshal`の基本
- 構造体タグによるフィールド名のカスタマイズ
- `omitempty`オプションの活用
- ネストした構造体やスライスのJSON変換

**なぜ重要か:** Web APIでは、リクエストの受信（JSON→構造体）とレスポンスの送信（構造体→JSON）が基本です。構造体タグでJSONのフィールド名をコントロールする方法は、Go WebAPI開発の最も基礎的なスキルです。

## 📖 概念

`encoding/json`パッケージはGoの構造体とJSONの相互変換を行います。構造体タグ（`` `json:"..."` ``）でフィールド名の変換ルールを指定します。GoはcamelCase、JSONはsnake_caseが多いため、タグで変換を制御します。

**背景と設計意図:** Goの構造体フィールドはPascalCase（大文字始まり）で書きますが、JSONのキーはsnake_caseやcamelCaseが一般的です。構造体タグはこのギャップを埋めるための仕組みです。

**よくある誤解:**

- ❌ 「小文字のフィールドもJSON変換される」→ unexported（小文字始まり）フィールドはJSON変換されない
- ❌ 「`omitempty`はゼロ値を省略する」→ 正確にはゼロ値のフィールドをJSON出力から省略する
- ❌ 「構造体タグなしでもOK」→ フィールド名がそのままJSONキーになる（PascalCase）

## 💡 コード例

### 基本: Marshal/Unmarshalの基礎

構造体とJSONの相互変換の基本を学びます。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`          // JSON: "id"
	Name  string `json:"name"`        // JSON: "name"
	Email string `json:"email"`       // JSON: "email"
	Age   int    `json:"age"`         // JSON: "age"
	role  string // 小文字フィールド → JSON変換されない（unexported）
}

func main() {
	// ---- Marshal: 構造体 → JSON ----
	user := User{
		ID:    1,
		Name:  "田中太郎",
		Email: "tanaka@example.com",
		Age:   30,
		role:  "admin", // unexported → JSONに含まれない
	}

	// json.Marshal → []byte を返す
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Println("JSON:", string(jsonBytes))
	// → {"id":1,"name":"田中太郎","email":"tanaka@example.com","age":30}

	// json.MarshalIndent → 整形されたJSON
	prettyJSON, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("Pretty JSON:")
	fmt.Println(string(prettyJSON))

	// ---- Unmarshal: JSON → 構造体 ----
	jsonStr := `{"id":2,"name":"鈴木花子","email":"suzuki@example.com","age":25}`

	var decoded User
	if err := json.Unmarshal([]byte(jsonStr), &decoded); err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	fmt.Printf("Decoded: %+v\n", decoded)
	// → {ID:2 Name:鈴木花子 Email:suzuki@example.com Age:25 role:}
}
```

> **💡 次のステップへ:** 基本的なMarshal/Unmarshalを学びました。次は構造体タグのオプション（`omitempty`など）を学びます。

### 応用: 構造体タグのオプション

`omitempty`や`-`オプションを使って、JSONの出力を制御する方法を学びます。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       int     `json:"price"`
	Description string  `json:"description,omitempty"` // 空文字列の場合はJSONから省略
	Discount    int     `json:"discount,omitempty"`    // 0の場合はJSONから省略
	InternalKey string  `json:"-"`                     // JSONに含めない
	Stock       *int    `json:"stock,omitempty"`       // nilの場合はJSONから省略
}

func intPtr(i int) *int { return &i }

func main() {
	// omitempty の動作確認
	products := []Product{
		{
			ID:          1,
			Name:        "Go入門書",
			Price:       3000,
			Description: "Goの基礎を学ぶ本",
			Discount:    500,
			InternalKey: "INTERNAL_001",
			Stock:       intPtr(10),
		},
		{
			ID:          2,
			Name:        "キーボード",
			Price:       15000,
			Description: "", // omitempty → JSONから省略される
			Discount:    0,  // omitempty → JSONから省略される
			InternalKey: "INTERNAL_002",
			Stock:       nil, // omitempty → JSONから省略される
		},
	}

	for _, p := range products {
		jsonBytes, _ := json.MarshalIndent(p, "", "  ")
		fmt.Println(string(jsonBytes))
		fmt.Println("---")
	}

	// ---- JSONに未知のフィールドがある場合 ----
	// 構造体に定義されていないフィールドは無視される
	jsonStr := `{"id":3,"name":"マウス","price":5000,"color":"black","weight":100}`

	var product Product
	if err := json.Unmarshal([]byte(jsonStr), &product); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Decoded (未知フィールドは無視): %+v\n", product)
}
```

> **💡 次のステップへ:** 構造体タグのオプションを学びました。次はネストした構造体やAPIレスポンスの実務パターンを学びます。

### 実践: APIレスポンスの設計パターン

実務でのAPIレスポンス構造体の設計パターンを学びます。

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// ---- APIレスポンスの共通構造 ----

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ---- ドメインモデル ----

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Author    Author    `json:"author"`              // ネストした構造体
	Tags      []string  `json:"tags,omitempty"`       // スライス
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ---- ページネーション ----

type PaginatedResponse struct {
	Items      []Article `json:"items"`
	TotalCount int       `json:"total_count"`
	Page       int       `json:"page"`
	PerPage    int       `json:"per_page"`
	HasMore    bool      `json:"has_more"`
}

func main() {
	now := time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC)

	// ---- 成功レスポンス ----
	article := Article{
		ID:    1,
		Title: "Go入門",
		Body:  "Goの基礎を学びましょう",
		Author: Author{
			ID:   1,
			Name: "田中太郎",
		},
		Tags:      []string{"Go", "プログラミング", "入門"},
		CreatedAt: now,
		UpdatedAt: now,
	}

	successResp := APIResponse{
		Success: true,
		Data:    article,
	}

	jsonBytes, _ := json.MarshalIndent(successResp, "", "  ")
	fmt.Println("=== 成功レスポンス ===")
	fmt.Println(string(jsonBytes))

	// ---- エラーレスポンス ----
	errorResp := APIResponse{
		Success: false,
		Error: &ErrorInfo{
			Code:    "NOT_FOUND",
			Message: "article not found",
		},
	}

	jsonBytes, _ = json.MarshalIndent(errorResp, "", "  ")
	fmt.Println("\n=== エラーレスポンス ===")
	fmt.Println(string(jsonBytes))

	// ---- JSONリクエストのパース ----
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
