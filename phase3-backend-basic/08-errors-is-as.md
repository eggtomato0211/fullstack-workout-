# 08. errors.Is/Asの活用 - エラーラッピング、エラーチェーン

## 🎯 このテーマで学ぶこと

- `fmt.Errorf`の`%w`動詞によるエラーラッピング
- `errors.Is`でエラーチェーンをたどって一致を判定
- `errors.As`でエラーチェーンからカスタムエラーを取得
- エラーに文脈を追加しながら伝搬させるパターン

**なぜ重要か:** 実務では、エラーは複数の層を通過して伝搬します（DB層→リポジトリ層→サービス層→ハンドラー層）。各層でエラーに文脈を追加しつつ、元のエラーの種類を判別できるようにする必要があります。`errors.Is`/`errors.As`はこのための標準的な仕組みです。

## 📖 概念

`fmt.Errorf("...: %w", err)`でエラーをラップすると、元のエラーを内包した新しいエラーが作られます。`errors.Is`はエラーチェーンをたどって指定のエラーと一致するか判定し、`errors.As`はエラーチェーンから特定の型のエラーを取り出します。

**背景と設計意図:** Go 1.13で導入されたエラーラッピングは、「エラーに文脈を追加しつつ、元のエラー情報を保持する」という実務ニーズに応えるものです。`%w`でラップし、`errors.Is`/`errors.As`でアンラップするのが標準パターンです。

**よくある誤解:**

- ❌ 「`%v`と`%w`は同じ」→ `%v`はラップしない（元のエラーを失う）、`%w`はラップする
- ❌ 「`==`でエラーを比較すればいい」→ ラップされたエラーは`==`で一致しない。`errors.Is`を使う
- ❌ 「全てのエラーをラップすべき」→ 文脈が追加されない場合はラップ不要

## 💡 コード例

### 基本: エラーラッピングとerrors.Is

`%w`でエラーをラップし、`errors.Is`でチェーンをたどる基本パターンを学びます。

```go
package main

import (
	"errors"
	"fmt"
)

// センチネルエラー
var ErrNotFound = errors.New("not found")
var ErrPermissionDenied = errors.New("permission denied")

// DB層: 元のエラーを返す
func findUserInDB(id int) (string, error) {
	if id == 99 {
		return "", ErrNotFound
	}
	return "田中太郎", nil
}

// リポジトリ層: エラーに文脈を追加してラップ
func getUserFromRepo(id int) (string, error) {
	name, err := findUserInDB(id)
	if err != nil {
		// %w でラップ: 元のエラー(ErrNotFound)を内包しつつ文脈を追加
		return "", fmt.Errorf("getUserFromRepo(id=%d): %w", id, err)
	}
	return name, nil
}

// サービス層: さらに文脈を追加
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
		// エラーメッセージ全体を表示（各層の文脈が含まれる）
		fmt.Println("エラー:", err)
		// → user service: getUserFromRepo(id=99): not found

		// errors.Is: エラーチェーンをたどって ErrNotFound と一致するか判定
		if errors.Is(err, ErrNotFound) {
			fmt.Println("→ ユーザーが見つかりません（404を返す）")
		}

		// ラップされていても元のエラーを検出できる
		if errors.Is(err, ErrPermissionDenied) {
			fmt.Println("→ 権限がありません（403を返す）")
		} else {
			fmt.Println("→ 権限エラーではない")
		}
	}
}
```

> **💡 次のステップへ:** `errors.Is`でセンチネルエラーを判定する方法を学びました。次は`errors.As`でカスタムエラー型を取り出す方法を学びます。

### 応用: errors.Asでカスタムエラーを取り出す

ラップされたエラーチェーンから、特定の型のカスタムエラーを取り出して詳細情報にアクセスします。

```go
package main

import (
	"errors"
	"fmt"
)

// ValidationError はフィールドのバリデーションエラー
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation: %s - %s", e.Field, e.Message)
}

// DBError はデータベース操作のエラー
type DBError struct {
	Query   string
	Message string
}

func (e *DBError) Error() string {
	return fmt.Sprintf("db error: %s (%s)", e.Message, e.Query)
}

func validateEmail(email string) error {
	if email == "" {
		return &ValidationError{Field: "email", Message: "is required"}
	}
	return nil
}

func saveUser(email string) error {
	if err := validateEmail(email); err != nil {
		// バリデーションエラーをラップ
		return fmt.Errorf("saveUser: %w", err)
	}
	// DB操作のシミュレーション
	return fmt.Errorf("saveUser: %w", &DBError{
		Query:   "INSERT INTO users",
		Message: "duplicate key",
	})
}

func main() {
	// ケース1: バリデーションエラー
	err := saveUser("")
	if err != nil {
		fmt.Println("エラー:", err)

		// errors.As: エラーチェーンから ValidationError を取り出す
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("→ バリデーションエラー: フィールド=%s, メッセージ=%s\n",
				ve.Field, ve.Message)
		}
	}

	// ケース2: DBエラー
	err = saveUser("test@example.com")
	if err != nil {
		fmt.Println("\nエラー:", err)

		var de *DBError
		if errors.As(err, &de) {
			fmt.Printf("→ DBエラー: クエリ=%s, メッセージ=%s\n",
				de.Query, de.Message)
		}
	}
}
```

> **💡 次のステップへ:** `errors.As`でカスタムエラーの詳細を取得する方法を学びました。次は実務でのエラー伝搬パターンを学びます。

### 実践: レイヤー構造でのエラー伝搬

実務のアプリケーションで、各層でエラーに文脈を追加しながら伝搬させるパターンを学びます。

```go
package main

import (
	"errors"
	"fmt"
)

// ---- エラー定義 ----

var ErrNotFound = errors.New("not found")
var ErrConflict = errors.New("conflict")

type AppError struct {
	StatusCode int
	Message    string
	Err        error // 元のエラーを保持
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %v", e.StatusCode, e.Message, e.Err)
	}
	return fmt.Sprintf("[%d] %s", e.StatusCode, e.Message)
}

// Unwrap で元のエラーを返す → errors.Is/As がチェーンをたどれる
func (e *AppError) Unwrap() error {
	return e.Err
}

// ---- DB層 ----

func findOrderInDB(id int) (string, error) {
	if id == 0 {
		return "", ErrNotFound
	}
	return "注文#" + fmt.Sprint(id), nil
}

// ---- リポジトリ層 ----

func getOrder(id int) (string, error) {
	order, err := findOrderInDB(id)
	if err != nil {
		return "", fmt.Errorf("order repository: %w", err)
	}
	return order, nil
}

// ---- サービス層 ----

func processOrder(id int) error {
	order, err := getOrder(id)
	if err != nil {
		// errors.Is で元のエラーを判定し、適切な AppError に変換
		if errors.Is(err, ErrNotFound) {
			return &AppError{
				StatusCode: 404,
				Message:    "order not found",
				Err:        err,
			}
		}
		return &AppError{
			StatusCode: 500,
			Message:    "internal error",
			Err:        err,
		}
	}
	fmt.Println("処理完了:", order)
	return nil
}

// ---- ハンドラー層 ----

func handleOrderRequest(id int) {
	err := processOrder(id)
	if err != nil {
		// AppError として取り出してステータスコードを使う
		var appErr *AppError
		if errors.As(err, &appErr) {
			fmt.Printf("HTTP %d: %s\n", appErr.StatusCode, appErr.Message)

			// さらに元のエラーもチェック可能（Unwrapチェーン）
			if errors.Is(err, ErrNotFound) {
				fmt.Println("→ リソースが見つかりません")
			}
		} else {
			fmt.Println("HTTP 500: unexpected error")
		}
		return
	}
}

func main() {
	handleOrderRequest(1) // 正常
	fmt.Println("---")
	handleOrderRequest(0) // 404
}
```

## 🎯 演習問題

データ取得パイプラインでのエラーラッピングを実装してください。

**要件:**

1. `ErrInvalidInput`と`ErrTimeout`のセンチネルエラーを定義
2. `fetchData(url string) (string, error)`: URLが空なら`ErrInvalidInput`をラップして返す。"slow"を含むなら`ErrTimeout`をラップして返す
3. `processData(url string) (string, error)`: `fetchData`を呼び、エラーに文脈を追加してラップ
4. `handleRequest(url string)`: `errors.Is`で判定し、`ErrInvalidInput`→"400 Bad Request"、`ErrTimeout`→"504 Gateway Timeout"を出力

**期待される動作:**

- `handleRequest("")` → "400 Bad Request"と表示
- `handleRequest("https://slow.example.com")` → "504 Gateway Timeout"と表示
- エラーメッセージに各層の文脈が含まれている

## ✅ 重要ポイント

- [ ] `%w`でラップ、`%v`はラップしない（元のエラーを保持するかの違い）
- [ ] `errors.Is`でセンチネルエラーとの一致をチェーン全体で判定
- [ ] `errors.As`でカスタムエラー型をチェーンから取り出す
- [ ] `Unwrap() error`メソッドを実装すると`errors.Is`/`As`がチェーンをたどれる

**次のテーマ:** [09. defer文の活用](./09-defer.md)
