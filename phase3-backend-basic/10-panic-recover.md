# 10. panic/recoverの理解 - 例外的状況の処理、ミドルウェアでの活用

## 🎯 このテーマで学ぶこと

- panicの発生とrecoverによる捕捉
- panicを使うべき場面と避けるべき場面
- deferと組み合わせたrecoverパターン（HTTPミドルウェア）

## 📖 なぜpanic/recoverを理解する必要があるのか

Goの通常のエラーハンドリングは`error`型を使いますが、**プログラムの続行が不可能な致命的エラー**にはpanicが使われます。また、HTTPサーバーでpanicをrecoverし500エラーとして返すのは実務の標準パターンです。

### こう書かないとどうなるか

```go
// panicを使いすぎると → try/catchのように乱用してしまう
func findUser(id int) string {
    user, ok := users[id]
    if !ok {
        panic("user not found") // ← これは通常のエラー。panicは不適切
    }
    return user
}

// 正しくは error で返す
func findUser(id int) (string, error) {
    user, ok := users[id]
    if !ok {
        return "", fmt.Errorf("user not found: id=%d", id)
    }
    return user, nil
}
```

### panicの使い分け

- **panicが適切**: 設定の不備（必須環境変数がない）、不変条件の違反（プログラマのミス）
- **panicが不適切**: ユーザー入力のエラー、ファイルが見つからない、ネットワークエラーなど通常起こりうる事象
- **命名慣習**: `Must`で始まる関数はpanic可能性を示す（`regexp.MustCompile`など）

## 💡 コード例

### 基本: panicとrecoverの動作

```go
package main

import (
	"fmt"
	"os"
)

// recoverはdeferの中でのみ動作する
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			// panicから回復し、エラーとして返す
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()

	return a / b, nil // ゼロ除算は runtime panic を引き起こす
}

// --- panicが適切な場面: Must パターン ---

type Config struct {
	DatabaseURL string
	Port        int
}

// Must で始まる関数はpanic可能性を示す慣習
// アプリ起動時の設定読み込みなど、失敗したら続行不可の場面で使う
func MustLoadConfig() *Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		panic("DATABASE_URL environment variable is required")
	}
	return &Config{DatabaseURL: dbURL, Port: 8080}
}

// --- panicが不適切な場面: 通常のエラーはerrorで返す ---
func findUser(id int) (string, error) {
	users := map[int]string{1: "田中太郎"}
	name, ok := users[id]
	if !ok {
		return "", fmt.Errorf("user not found: id=%d", id) // panicしない
	}
	return name, nil
}

func main() {
	// 正常ケース
	result, err := safeDivide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 3 =", result)
	}

	// panicが発生するケース → recoverしてエラーとして返す
	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // recovered from panic: ...
	}

	// recoverしたのでプログラムは続行できる
	fmt.Println("プログラムは続行中")

	// 通常のエラーはerrorで処理
	_, err = findUser(99)
	if err != nil {
		fmt.Println("通常のエラー:", err)
	}
}
```

### 実践: HTTPミドルウェアでのrecover

実際のWebサーバーでpanic recoveryミドルウェアを実装するパターンです。リクエスト処理中のpanicをキャッチし、サーバー全体がクラッシュしないようにします。

```go
package main

import "fmt"

// --- HTTPをシミュレートする簡易的な型 ---

type Request struct {
	Path   string
	Method string
}

type Response struct {
	StatusCode int
	Body       string
}

type Handler func(req *Request) *Response

// --- recoveryミドルウェア ---
// なぜ必要か: 1つのリクエストのpanicでサーバー全体が落ちるのを防ぐ
func withRecovery(next Handler) Handler {
	return func(req *Request) (resp *Response) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("[PANIC RECOVERED] %s %s: %v\n", req.Method, req.Path, r)
				resp = &Response{
					StatusCode: 500,
					Body:       "Internal Server Error",
				}
			}
		}()
		return next(req)
	}
}

// --- ロギングミドルウェア ---
func withLogging(next Handler) Handler {
	return func(req *Request) *Response {
		fmt.Printf("[LOG] %s %s\n", req.Method, req.Path)
		resp := next(req)
		fmt.Printf("[LOG] %s %s → %d\n", req.Method, req.Path, resp.StatusCode)
		return resp
	}
}

// --- ハンドラー ---

func handleHello(req *Request) *Response {
	return &Response{StatusCode: 200, Body: "Hello, World!"}
}

func handlePanic(req *Request) *Response {
	panic("something unexpected happened!")
}

func main() {
	// ミドルウェアを適用: recovery → logging → handler
	helloHandler := withRecovery(withLogging(handleHello))
	panicHandler := withRecovery(withLogging(handlePanic))

	// 正常なリクエスト
	resp := helloHandler(&Request{Path: "/hello", Method: "GET"})
	fmt.Printf("Response: %d - %s\n\n", resp.StatusCode, resp.Body)

	// panicが発生するリクエスト → recoverで500が返る
	resp = panicHandler(&Request{Path: "/panic", Method: "GET"})
	fmt.Printf("Response: %d - %s\n\n", resp.StatusCode, resp.Body)

	// サーバーはクラッシュしない
	fmt.Println("サーバーは稼働中...")
}
```

## 🎯 演習問題

安全なバッチ処理システムをdefer/recoverで実装してください。

**要件:**

1. `Job`構造体: `Name string`, `Fn func() error`を持つ
2. `RunJob(job Job) (err error)`: 1つのジョブを実行。panic時はrecoverしてエラーとして返す
3. `RunBatch(jobs []Job) []error`: 複数のジョブを順に実行。各ジョブのpanicが他に影響しないようにする
4. 結果レポートを出力: 成功数、失敗数、panic数

**期待される動作:**

- 正常なジョブ → 成功
- errorを返すジョブ → 失敗として記録
- panicするジョブ → recoverして失敗として記録、次のジョブは実行される

## ✅ 重要ポイント

- [ ] panicは致命的な状況（設定不備、不変条件違反）に使う。通常のエラーにはerror型
- [ ] recoverはdeferの中でのみ動作する
- [ ] `Must`で始まる関数はpanicの可能性を示す慣習
- [ ] HTTPサーバーではrecoveryミドルウェアで500を返すのが標準パターン

**次のテーマ:** [11. encoding/jsonと構造体タグ](./11-encoding-json.md)
