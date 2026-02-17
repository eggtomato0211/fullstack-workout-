# 06. エラーハンドリングの基本 - error型、errorsパッケージ

## 🎯 このテーマで学ぶこと

- Goのerror型の仕組み
- `if err != nil`パターンの正しい使い方
- エラーメッセージの書き方の慣習
- エラーの返し方と受け取り方

**なぜ重要か:** Goには例外（try/catch）がありません。エラーは「戻り値」として明示的に扱います。これはGoの最も特徴的な設計であり、全てのGoコードの基盤です。エラーハンドリングを正しく行わないと、サイレントなバグや障害につながります。

## 📖 概念

Goの`error`は`Error() string`メソッドを持つインターフェースです。関数は最後の戻り値としてerrorを返し、呼び出し側は`if err != nil`でチェックします。

**背景と設計意図:** 例外機構は制御フローを複雑にし、どこでエラーが発生するか追跡しにくいという課題があります。Goはエラーを明示的な値として扱うことで、エラー処理を必ず書くことを開発者に促し、堅牢なコードを書きやすくしています。

**実務での活用場面:** API呼び出し、DB操作、ファイル操作、バリデーションなど、失敗する可能性のある全ての操作で使います。

**よくある誤解:**

- ❌ 「エラーは無視してもいい」→ `_`で受け取ってはいけない（特に本番コード）
- ❌ 「エラーメッセージは大文字で始める」→ 小文字で始め、ピリオドで終わらない（Go慣習）
- ❌ 「`if err != nil`は冗長だ」→ 明示的なエラー処理はGoの哲学。慣れると読みやすい

## 💡 コード例

### 基本: error型の基礎

Goのerrorの仕組みと、基本的なエラーチェックパターンを学びます。

```go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// ---- 標準ライブラリのエラーを受け取る ----

	// strconv.Atoi は文字列を整数に変換する
	// 変換できない場合はerrorを返す
	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("変換エラー:", err)
		return
	}
	fmt.Println("変換成功:", num)

	// 変換失敗のケース
	_, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("変換エラー:", err) // strconv.Atoi: parsing "abc": invalid syntax
	}

	// ---- 自分でエラーを作成する ----

	// errors.New で簡単なエラーを作成
	err = errors.New("something went wrong")
	fmt.Println(err) // something went wrong

	// fmt.Errorf で書式付きエラーを作成
	name := "admin"
	err = fmt.Errorf("user not found: %s", name)
	fmt.Println(err) // user not found: admin
}
```

> **💡 次のステップへ:** エラーの基本を学びました。次は関数でエラーを返すパターンを学びます。

### 応用: 関数でエラーを返す

関数の戻り値としてerrorを返すパターンと、エラーメッセージの書き方の慣習を学びます。

```go
package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
}

// validateUser はバリデーション結果をerrorで返す
// エラーメッセージの慣習:
//   - 小文字で始める
//   - ピリオドで終わらない
//   - 文脈を含める（何が失敗したか分かるように）
func validateUser(u User) error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Age < 0 || u.Age > 150 {
		return fmt.Errorf("invalid age: %d (must be 0-150)", u.Age)
	}
	return nil // エラーなし = nil を返す
}

// divide は除算を行い、ゼロ除算の場合にエラーを返す
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	// ---- バリデーションのエラーチェック ----
	user := User{Name: "", Email: "test@example.com", Age: 25}
	if err := validateUser(user); err != nil {
		fmt.Println("バリデーションエラー:", err)
	}

	user = User{Name: "田中太郎", Email: "tanaka@example.com", Age: 25}
	if err := validateUser(user); err != nil {
		fmt.Println("バリデーションエラー:", err)
	} else {
		fmt.Println("バリデーション成功")
	}

	// ---- 除算のエラーチェック ----
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("計算エラー:", err)
		return
	}
	fmt.Printf("10 / 3 = %.2f\n", result)

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("計算エラー:", err) // division by zero
	}
}
```

> **💡 次のステップへ:** 関数でのエラー返却を学びました。次は複数のエラーチェックを組み合わせた実践的なパターンを学びます。

### 実践: エラー処理の実践パターン

実務でよくある「複数のステップでエラーが起こりうる処理」のパターンを学びます。

```go
package main

import (
	"errors"
	"fmt"
	"strings"
)

type Config struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

// parseConfig は設定文字列をパースする
// 各ステップでエラーチェックし、早期リターンする
func parseConfig(raw string) (*Config, error) {
	if raw == "" {
		return nil, errors.New("config string is empty")
	}

	parts := strings.Split(raw, ";")
	if len(parts) != 5 {
		return nil, fmt.Errorf("expected 5 parts, got %d", len(parts))
	}

	config := &Config{
		Host:     parts[0],
		Database: parts[2],
		User:     parts[3],
		Password: parts[4],
	}

	// ポートのパース
	var port int
	_, err := fmt.Sscanf(parts[1], "%d", &port)
	if err != nil {
		return nil, fmt.Errorf("invalid port %q: %v", parts[1], err)
	}
	if port < 1 || port > 65535 {
		return nil, fmt.Errorf("port out of range: %d", port)
	}
	config.Port = port

	// ホストのバリデーション
	if config.Host == "" {
		return nil, errors.New("host is required")
	}

	return config, nil
}

// connectToDatabase は設定を使ってDB接続する（シミュレーション）
func connectToDatabase(config *Config) error {
	if config.User == "" {
		return errors.New("database user is required")
	}
	fmt.Printf("Connected to %s:%d/%s as %s\n",
		config.Host, config.Port, config.Database, config.User)
	return nil
}

func main() {
	// 正常ケース
	configStr := "localhost;5432;mydb;admin;secret"
	config, err := parseConfig(configStr)
	if err != nil {
		fmt.Println("設定エラー:", err)
		return
	}

	if err := connectToDatabase(config); err != nil {
		fmt.Println("接続エラー:", err)
		return
	}

	fmt.Println("セットアップ完了!")

	// エラーケース
	fmt.Println("\n--- エラーケース ---")
	badConfigs := []string{
		"",                                // 空文字列
		"localhost;abc;mydb;admin;secret",  // 無効なポート
		"localhost;99999;mydb;admin;secret", // ポート範囲外
	}

	for _, c := range badConfigs {
		_, err := parseConfig(c)
		if err != nil {
			fmt.Println("設定エラー:", err)
		}
	}
}
```

## 🎯 演習問題

ユーザー登録処理をエラーハンドリング付きで実装してください。

**要件:**

1. `RegisterRequest`構造体: `Username`, `Email`, `Password string`を持つ
2. `ValidateRequest(req RegisterRequest) error`: 以下をチェック
   - Usernameが空でないこと
   - Usernameが3文字以上であること
   - Emailに`@`が含まれること
   - Passwordが8文字以上であること
3. `RegisterUser(req RegisterRequest) (*User, error)`: バリデーション→ユーザー作成。各ステップでエラーチェック
4. エラーメッセージはGoの慣習に従う（小文字開始、ピリオドなし）

**期待される動作:**

- `RegisterUser({Username: "ab", ...})` → "username must be at least 3 characters"
- `RegisterUser({Email: "invalid", ...})` → "email must contain @"
- 正常な入力 → User構造体を返す

## ✅ 重要ポイント

- [ ] エラーは戻り値として明示的に扱う（例外ではない）
- [ ] `if err != nil`で必ずチェックする
- [ ] エラーメッセージは小文字で始め、ピリオドで終わらない
- [ ] エラーが起きたら早期リターンする（ネストを浅く保つ）

**次のテーマ:** [07. カスタムエラーの実装](./07-custom-errors.md)
