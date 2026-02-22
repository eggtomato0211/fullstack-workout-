# 06. エラーハンドリングの基本 - error型、errorsパッケージ

## 🎯 このテーマで学ぶこと

- Goのerror型の仕組みと`if err != nil`パターン
- エラーメッセージの書き方の慣習
- 複数ステップでのエラー処理と早期リターン

## 📖 なぜGoのエラーハンドリングを理解する必要があるのか

Goには例外（try/catch）がありません。エラーは「戻り値」として明示的に扱います。これはGoの最も特徴的な設計であり、全てのGoコードの基盤です。

### こう書かないとどうなるか

```go
// エラーを無視すると...
result, _ := doSomething()  // エラーを _ で捨てている
// → doSomethingが失敗してもresultはゼロ値のまま処理が続く
// → 後の処理で「なぜかデータがない」という分かりにくいバグになる
```

Goでは`if err != nil`を「冗長だ」と感じるかもしれませんが、これがGoの哲学です。エラーを必ず明示的に処理することで、「何が失敗したか」「どこで失敗したか」がコードから一目で分かります。

### エラーメッセージの慣習

Goのエラーメッセージには決まった書き方があります：

```go
// 正しい慣習
return errors.New("user not found")          // 小文字で始める
return fmt.Errorf("invalid age: %d", age)    // ピリオドで終わらない

// 避けるべき書き方
return errors.New("User not found.")  // 大文字開始、ピリオド終わり
return errors.New("Error: failed")    // "Error:"は冗長
```

なぜか？ エラーメッセージは `fmt.Errorf("parse config: %v", err)` のようにラップされて連結されるため、各メッセージが文の断片として読めるべきだからです。

## 💡 コード例

### 基本: error型とエラーを返す関数

エラーの作成方法と、関数でエラーを返すパターンを学びます。

```go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

type User struct {
	Name  string
	Email string
	Age   int
}

// バリデーション結果をerrorで返す
func validateUser(u User) error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Age < 0 || u.Age > 150 {
		// fmt.Errorfで文脈（具体的な値）を含める
		return fmt.Errorf("invalid age: %d (must be 0-150)", u.Age)
	}
	return nil // エラーなし = nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	// 標準ライブラリのエラーを受け取る
	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("変換エラー:", err)
		return
	}
	fmt.Println("変換成功:", num)

	// バリデーション
	user := User{Name: "", Email: "test@example.com", Age: 25}
	if err := validateUser(user); err != nil {
		fmt.Println("バリデーションエラー:", err) // name is required
	}

	// 除算
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

### 実践: 複数ステップのエラー処理

実務では「設定のパース → バリデーション → DB接続」のように、複数のステップが連続する処理がよくあります。各ステップでエラーチェックし、早期リターンすることでネストを浅く保ちます。

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

// 各ステップでエラーチェックし、早期リターン
// → ネストが深くならず、正常系の流れが読みやすい
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

	var port int
	_, err := fmt.Sscanf(parts[1], "%d", &port)
	if err != nil {
		return nil, fmt.Errorf("invalid port %q: %v", parts[1], err)
	}
	if port < 1 || port > 65535 {
		return nil, fmt.Errorf("port out of range: %d", port)
	}
	config.Port = port

	if config.Host == "" {
		return nil, errors.New("host is required")
	}

	return config, nil
}

func connectToDatabase(config *Config) error {
	if config.User == "" {
		return errors.New("database user is required")
	}
	fmt.Printf("Connected to %s:%d/%s as %s\n",
		config.Host, config.Port, config.Database, config.User)
	return nil
}

func main() {
	// 正常ケース：パース → 接続の流れ
	config, err := parseConfig("localhost;5432;mydb;admin;secret")
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
		"",                               // 空文字列
		"localhost;abc;mydb;admin;secret", // 無効なポート
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

1. `RegisterRequest`構造体: `Username`, `Email`, `Password string`
2. `ValidateRequest(req RegisterRequest) error`: Username空/3文字未満、Emailに@なし、Password8文字未満でエラー
3. `RegisterUser(req RegisterRequest) (*User, error)`: バリデーション → ユーザー作成
4. エラーメッセージはGoの慣習に従う（小文字開始、ピリオドなし）

**期待される動作:**

- `RegisterUser({Username: "ab", ...})` → "username must be at least 3 characters"
- `RegisterUser({Email: "invalid", ...})` → "email must contain @"

## ✅ 重要ポイント

- [ ] エラーは戻り値として明示的に扱う（例外ではない）
- [ ] `if err != nil`で必ずチェックし、`_`で無視しない
- [ ] エラーメッセージは小文字で始め、ピリオドで終わらない
- [ ] 早期リターンでネストを浅く保つ

**次のテーマ:** [07. カスタムエラーの実装](./07-custom-errors.md)
