# 08. errors.Is/Asの活用 - エラーラッピング、エラーチェーン

## 🎯 このテーマで学ぶこと

- `fmt.Errorf`の`%w`動詞によるエラーラッピング
- `errors.Is`でエラーチェーンをたどって一致を判定
- `errors.As`でエラーチェーンからカスタムエラーを取得

## 📖 なぜerrors.Is/Asを理解する必要があるのか

実務のアプリケーションでは、エラーは複数の層を通過して伝搬します（DB層→リポジトリ層→サービス層→ハンドラー層）。各層でエラーに文脈を追加しつつ、元のエラーの種類も判別したい。`errors.Is`/`errors.As`はこの「ラップされたエラーの中身を調べる」仕組みです。

### こう書かないとどうなるか

```go
// %v でフォーマット → 元のエラー情報が失われる
return fmt.Errorf("user service: %v", err)
// ↓ この後 errors.Is(err, ErrNotFound) が false になる！

// %w でラップ → 元のエラーを保持したまま文脈を追加
return fmt.Errorf("user service: %w", err)
// ↓ errors.Is(err, ErrNotFound) が true になる

// == で比較 → ラップされたエラーは一致しない
if err == ErrNotFound { ... } // ラップされていると false

// errors.Is → チェーンをたどって一致を判定
if errors.Is(err, ErrNotFound) { ... } // ラップされていても true
```

`%v`と`%w`の1文字の違いが、エラーの追跡可能性を決定します。

## 💡 コード例

### 基本: エラーラッピングとerrors.Is/As

```go
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")
var ErrPermissionDenied = errors.New("permission denied")

// ValidationError はカスタムエラー型
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation: %s - %s", e.Field, e.Message)
}

// --- 各層でエラーに文脈を追加しながら伝搬 ---

// DB層: 元のエラーを返す
func findUserInDB(id int) (string, error) {
	if id == 99 {
		return "", ErrNotFound
	}
	return "田中太郎", nil
}

// リポジトリ層: %w で文脈を追加してラップ
func getUserFromRepo(id int) (string, error) {
	name, err := findUserInDB(id)
	if err != nil {
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

func validateEmail(email string) error {
	if email == "" {
		return &ValidationError{Field: "email", Message: "is required"}
	}
	return nil
}

func createUser(email string) error {
	if err := validateEmail(email); err != nil {
		return fmt.Errorf("createUser: %w", err) // カスタムエラーもラップ可能
	}
	return nil
}

func main() {
	// --- errors.Is: センチネルエラーの判定 ---
	_, err := getUserService(99)
	if err != nil {
		fmt.Println("エラー:", err)
		// → user service: getUserFromRepo(id=99): not found

		// errors.Is: ラップされていても元のErrNotFoundを検出できる
		if errors.Is(err, ErrNotFound) {
			fmt.Println("→ 404を返す")
		}
		if errors.Is(err, ErrPermissionDenied) {
			fmt.Println("→ 403を返す")
		} else {
			fmt.Println("→ 権限エラーではない")
		}
	}

	// --- errors.As: カスタムエラー型の取り出し ---
	err = createUser("")
	if err != nil {
		fmt.Println("\nエラー:", err)

		// errors.As: ラップされたValidationErrorを取り出す
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("→ バリデーションエラー: フィールド=%s, メッセージ=%s\n",
				ve.Field, ve.Message)
		}
	}
}
```

### 実践: Unwrapとレイヤー構造でのエラー伝搬

カスタムエラー型に`Unwrap()`メソッドを実装すると、`errors.Is`/`errors.As`がチェーンをたどれるようになります。

```go
package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

// AppError はアプリケーション共通のエラー型
// Unwrap()を実装して、errors.Is/Asがチェーンをたどれるようにする
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

// Unwrap で元のエラーを返す
// → errors.Is(appErr, ErrNotFound) が機能するようになる
func (e *AppError) Unwrap() error {
	return e.Err
}

// --- DB層 → リポジトリ層 → サービス層 → ハンドラー層 ---

func findOrderInDB(id int) (string, error) {
	if id == 0 {
		return "", ErrNotFound
	}
	return "注文#" + fmt.Sprint(id), nil
}

func getOrder(id int) (string, error) {
	order, err := findOrderInDB(id)
	if err != nil {
		return "", fmt.Errorf("order repository: %w", err)
	}
	return order, nil
}

func processOrder(id int) error {
	order, err := getOrder(id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return &AppError{StatusCode: 404, Message: "order not found", Err: err}
		}
		return &AppError{StatusCode: 500, Message: "internal error", Err: err}
	}
	fmt.Println("処理完了:", order)
	return nil
}

func handleOrderRequest(id int) {
	err := processOrder(id)
	if err != nil {
		// errors.As でAppErrorを取り出してステータスコードを使う
		var appErr *AppError
		if errors.As(err, &appErr) {
			fmt.Printf("HTTP %d: %s\n", appErr.StatusCode, appErr.Message)
		}

		// Unwrapのおかげで、AppErrorの中のErrNotFoundも検出できる
		if errors.Is(err, ErrNotFound) {
			fmt.Println("→ リソースが見つかりません")
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

## ✅ 重要ポイント

- [ ] `%w`でラップ、`%v`はラップしない（元のエラーを保持するかの違い）
- [ ] `errors.Is`でセンチネルエラーとの一致をチェーン全体で判定
- [ ] `errors.As`でカスタムエラー型をチェーンから取り出す
- [ ] `Unwrap() error`メソッドを実装すると`errors.Is`/`As`がチェーンをたどれる

**次のテーマ:** [09. defer文の活用](./09-defer.md)
