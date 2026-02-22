# 07. カスタムエラーの実装 - errors.New、fmt.Errorf、カスタムエラー型

## 🎯 このテーマで学ぶこと

- カスタムエラー型の定義（Error()メソッドの実装）
- センチネルエラー（パッケージレベルのエラー変数）
- エラーに文脈情報を持たせてHTTPステータスコードと連携する方法

## 📖 なぜカスタムエラーを理解する必要があるのか

基本的な`errors.New`や`fmt.Errorf`だけでは、エラーの**種類**を判別できません。実務のAPIサーバーでは「このエラーはNotFoundなのかValidationエラーなのか」によって返すHTTPステータスコードやメッセージが変わります。

### こう書かないとどうなるか

```go
// 文字列だけのエラーでは種類の判別ができない
err := errors.New("user not found")

// こうするしかない → 脆い（メッセージを変えたら壊れる）
if err.Error() == "user not found" {
    w.WriteHeader(404)
}

// カスタムエラー型なら → 型で安全に判別できる
if _, ok := err.(*NotFoundError); ok {
    w.WriteHeader(404) // 型が合えば確実にNotFound
}
```

### センチネルエラー vs カスタムエラー型

- **センチネルエラー**（`var ErrNotFound = errors.New("not found")`）: 単純な種類の判別に。詳細情報が不要な場面
- **カスタムエラー型**（`type ValidationError struct{...}`）: エラーに追加情報（フィールド名、IDなど）を持たせたい場面

## 💡 コード例

### 基本: センチネルエラーとカスタムエラー型

```go
package main

import (
	"errors"
	"fmt"
)

// --- センチネルエラー ---
// 命名規則: Err + 名前
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
)

// --- カスタムエラー型 ---
// エラーに詳細情報を持たせる

// ValidationError はバリデーションエラーの詳細を持つ
type ValidationError struct {
	Field   string // どのフィールドがエラーか
	Message string // 何が問題か
}

// Error() を実装 → error インターフェースを満たす
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", e.Field, e.Message)
}

// NotFoundError はリソースが見つからないエラー
type NotFoundError struct {
	Resource string // user, product など
	ID       int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found: id=%d", e.Resource, e.ID)
}

// --- 使用例 ---

func validateAge(age int) error {
	if age < 0 || age > 150 {
		return &ValidationError{
			Field:   "age",
			Message: fmt.Sprintf("must be 0-150, got %d", age),
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
	// 型アサーションでカスタムエラーの詳細にアクセス
	err := validateAge(-5)
	if err != nil {
		if ve, ok := err.(*ValidationError); ok {
			fmt.Printf("フィールド「%s」: %s\n", ve.Field, ve.Message)
		}
	}

	_, err = findProduct(99)
	if err != nil {
		if nfe, ok := err.(*NotFoundError); ok {
			fmt.Printf("%sが見つかりません (ID: %d)\n", nfe.Resource, nfe.ID)
		}
	}
}
```

### 実践: HTTPレスポンスとカスタムエラーの連携

APIサーバーでカスタムエラーをHTTPステータスコードに変換するパターンです。実務で最もよく使うカスタムエラーの使い方です。

```go
package main

import "fmt"

// AppError はアプリケーション共通のエラー型
// なぜCode/Message/Detailを分けるか：
// - Code: HTTPステータスコード（プログラムが使う）
// - Message: ユーザー向けメッセージ（クライアントに返す）
// - Detail: 開発者向け詳細（ログに出す、クライアントには返さない）
type AppError struct {
	Code    int
	Message string
	Detail  string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%d] %s: %s", e.Code, e.Message, e.Detail)
}

// ヘルパー関数でエラー生成を統一
func NewBadRequestError(detail string) *AppError {
	return &AppError{Code: 400, Message: "Bad Request", Detail: detail}
}

func NewNotFoundError(detail string) *AppError {
	return &AppError{Code: 404, Message: "Not Found", Detail: detail}
}

// --- ビジネスロジック ---

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

// HTTPハンドラーをシミュレート
func handleRequest(productID int) {
	product, err := getProduct(productID)
	if err != nil {
		if appErr, ok := err.(*AppError); ok {
			fmt.Printf("HTTP %d: %s (%s)\n", appErr.Code, appErr.Message, appErr.Detail)
		} else {
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
