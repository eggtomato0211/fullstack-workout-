# 07. カスタムエラーの実装 - errors.New、fmt.Errorf、カスタムエラー型

## 🎯 このテーマで学ぶこと

- `errors.New`と`fmt.Errorf`の使い分け
- カスタムエラー型の定義（Error()メソッドの実装）
- センチネルエラー（パッケージレベルのエラー変数）
- エラーに文脈情報を持たせる方法

**なぜ重要か:** 基本的なエラーメッセージだけでは、エラーの種類に応じた分岐処理ができません。カスタムエラーを使うことで、「このエラーはNotFoundなのかValidationエラーなのか」をプログラムで判別でき、適切なHTTPステータスコードやユーザーメッセージを返せるようになります。

## 📖 概念

Goのerrorインターフェースは`Error() string`を実装するだけで満たせます。これを利用して、エラーの種類や詳細情報を持つカスタムエラー型を作成できます。また、パッケージレベルでエラー変数を定義する「センチネルエラー」パターンもよく使われます。

**よくある誤解:**

- ❌ 「エラーはstringだけで十分」→ 型で分岐したい場面ではカスタムエラーが必要
- ❌ 「全てのエラーにカスタム型を作る」→ 分岐が不要な場面では`errors.New`で十分
- ❌ 「センチネルエラーを大量に定義する」→ 必要最小限にとどめる

## 💡 コード例

### 基本: センチネルエラー

パッケージレベルでエラー変数を定義し、エラーの種類を判別できるようにします。

```go
package main

import (
	"errors"
	"fmt"
)

// センチネルエラー: パッケージレベルで定義する定数的なエラー
// 命名規則: Err + 名前（例: ErrNotFound, ErrUnauthorized）
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
)

type User struct {
	ID   int
	Name string
	Role string
}

var users = map[int]*User{
	1: {ID: 1, Name: "田中太郎", Role: "admin"},
	2: {ID: 2, Name: "鈴木花子", Role: "user"},
}

func findUser(id int) (*User, error) {
	user, ok := users[id]
	if !ok {
		return nil, ErrNotFound // センチネルエラーを返す
	}
	return user, nil
}

func checkPermission(user *User, action string) error {
	if user.Role != "admin" && action == "delete" {
		return ErrForbidden
	}
	return nil
}

func main() {
	// センチネルエラーとの比較で分岐処理
	user, err := findUser(99)
	if err != nil {
		if err == ErrNotFound {
			fmt.Println("ユーザーが見つかりません")
		} else {
			fmt.Println("予期しないエラー:", err)
		}
		return
	}
	fmt.Println("Found:", user.Name)
}
```

> **💡 次のステップへ:** センチネルエラーで基本的なエラー分岐を学びました。次はエラーに詳細情報を持たせるカスタムエラー型を学びます。

### 応用: カスタムエラー型

エラーに詳細情報（ステータスコード、フィールド名など）を持たせるカスタム型を定義します。

```go
package main

import "fmt"

// ValidationError はバリデーションエラーの詳細を持つカスタムエラー型
type ValidationError struct {
	Field   string // エラーが発生したフィールド名
	Message string // エラーメッセージ
}

// Error() を実装して error インターフェースを満たす
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

// NotFoundError はリソースが見つからないエラー
type NotFoundError struct {
	Resource string // リソース種別（user, product 等）
	ID       int    // 検索に使ったID
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found: id=%d", e.Resource, e.ID)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Message: "must be non-negative",
		}
	}
	if age > 150 {
		return &ValidationError{
			Field:   "age",
			Message: fmt.Sprintf("must be <= 150, got %d", age),
		}
	}
	return nil
}

func findProduct(id int) (string, error) {
	products := map[int]string{1: "Go入門書", 2: "キーボード"}
	name, ok := products[id]
	if !ok {
		return "", &NotFoundError{Resource: "product", ID: id}
	}
	return name, nil
}

func main() {
	// ValidationError の型アサーションでフィールド名を取得
	err := validateAge(-5)
	if err != nil {
		// 型アサーションでカスタムエラーの詳細にアクセス
		if ve, ok := err.(*ValidationError); ok {
			fmt.Printf("フィールド「%s」のエラー: %s\n", ve.Field, ve.Message)
		}
	}

	// NotFoundError の型アサーションでリソース情報を取得
	_, err = findProduct(99)
	if err != nil {
		if nfe, ok := err.(*NotFoundError); ok {
			fmt.Printf("%sが見つかりません (ID: %d)\n", nfe.Resource, nfe.ID)
		}
	}
}
```

> **💡 次のステップへ:** カスタムエラー型の定義と型アサーションでの分岐を学びました。次は実務でのHTTPステータスコードとの連携パターンを学びます。

### 実践: HTTPレスポンスとカスタムエラーの連携

APIサーバーでカスタムエラーをHTTPステータスコードに変換するパターンを学びます。

```go
package main

import "fmt"

// AppError はアプリケーション共通のエラー型
type AppError struct {
	Code    int    // HTTPステータスコード
	Message string // ユーザー向けメッセージ
	Detail  string // 開発者向け詳細情報
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%d] %s: %s", e.Code, e.Message, e.Detail)
}

// エラー生成のヘルパー関数
func NewBadRequestError(detail string) *AppError {
	return &AppError{Code: 400, Message: "Bad Request", Detail: detail}
}

func NewNotFoundError(detail string) *AppError {
	return &AppError{Code: 404, Message: "Not Found", Detail: detail}
}

func NewInternalError(detail string) *AppError {
	return &AppError{Code: 500, Message: "Internal Server Error", Detail: detail}
}

// ---- ビジネスロジック ----

type Product struct {
	ID    int
	Name  string
	Price int
}

var products = map[int]*Product{
	1: {ID: 1, Name: "Go入門書", Price: 3000},
}

func getProduct(id int) (*Product, error) {
	if id <= 0 {
		return nil, NewBadRequestError(fmt.Sprintf("invalid product id: %d", id))
	}
	product, ok := products[id]
	if !ok {
		return nil, NewNotFoundError(fmt.Sprintf("product id=%d", id))
	}
	return product, nil
}

// handleRequest はHTTPハンドラーをシミュレート
func handleRequest(productID int) {
	product, err := getProduct(productID)
	if err != nil {
		// AppErrorならステータスコードに応じた処理
		if appErr, ok := err.(*AppError); ok {
			fmt.Printf("HTTP %d: %s (%s)\n", appErr.Code, appErr.Message, appErr.Detail)
		} else {
			// 予期しないエラーは500で返す
			fmt.Printf("HTTP 500: Internal Server Error (%v)\n", err)
		}
		return
	}
	fmt.Printf("HTTP 200: %+v\n", product)
}

func main() {
	handleRequest(1)  // HTTP 200: 正常
	handleRequest(99) // HTTP 404: Not Found
	handleRequest(-1) // HTTP 400: Bad Request
}
```

## 🎯 演習問題

ファイル操作を想定したカスタムエラーシステムを設計してください。

**要件:**

1. `FileError`カスタムエラー型: `Op string`（操作名）, `Path string`（ファイルパス）, `Message string`を持つ
2. `ReadFile(path string) (string, error)`: パスが空ならFileError、存在しないパスならFileErrorを返す
3. `WriteFile(path, content string) error`: パスが空、contentが空ならそれぞれFileErrorを返す
4. `handleFileError(err error)`: 型アサーションでFileErrorの詳細を表示、それ以外は一般エラーとして表示

**期待される動作:**

- `ReadFile("")` → FileError{Op: "read", Path: "", Message: "path is empty"}
- `handleFileError(err)` → "ファイル操作エラー [read] : path is empty"

## ✅ 重要ポイント

- [ ] センチネルエラー（`var ErrXxx = errors.New(...)`）で既知のエラーを定義
- [ ] カスタムエラー型でエラーに詳細情報を持たせる
- [ ] `Error() string`を実装すればerrorインターフェースを満たす
- [ ] 型アサーション（`err.(*CustomError)`）でエラーの種類を判別する

**次のテーマ:** [08. errors.Is/Asの活用](./08-errors-is-as.md)
