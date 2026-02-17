# 10. panic/recoverの理解 - 例外的状況の処理、ミドルウェアでの活用

## 🎯 このテーマで学ぶこと

- panicの発生とプログラムの中断
- recoverによるpanicの捕捉
- deferと組み合わせたrecoverパターン
- panicを使うべき場面と避けるべき場面

**なぜ重要か:** Goの通常のエラーハンドリングは`error`型を使いますが、プログラムの続行が不可能な致命的エラー（設定の不備、不変条件の違反など）にはpanicが使われます。また、HTTPサーバーのミドルウェアでpanicをrecoverし、500エラーとして返すのは実務での標準パターンです。

## 📖 概念

`panic`はプログラムの実行を中断し、スタックを巻き戻す仕組みです。`recover`はpanic中のスタック巻き戻しを止め、正常な実行に復帰させます。recoverは`defer`の中でのみ動作します。

**背景と設計意図:** Goはエラーを戻り値で扱うことを推奨しますが、「プログラムが回復不能な状態」ではpanicが適切です。recoverはサーバーアプリケーションでリクエスト単位のpanicをキャッチし、サーバー全体がクラッシュしないようにするために使います。

**よくある誤解:**

- ❌ 「panicはtry/catchの代わりに使える」→ 通常のエラーにはerror型を使う。panicは例外的状況のみ
- ❌ 「recoverでどこからでもpanicを捕捉できる」→ deferの中でのみ動作する
- ❌ 「panicは使ってはいけない」→ 適切な場面（初期化の失敗、不変条件の違反）では有用

## 💡 コード例

### 基本: panicとrecoverの動作

panicの発生とrecoverによる捕捉の基本パターンを学びます。

```go
package main

import "fmt"

// safeDivide はゼロ除算でもpanicしないようにrecoverで保護
func safeDivide(a, b int) (result int, err error) {
	// deferの中でrecoverを呼ぶ
	defer func() {
		if r := recover(); r != nil {
			// panicから回復し、エラーとして返す
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()

	// ゼロ除算は runtime panic を引き起こす
	return a / b, nil
}

func main() {
	// 正常ケース
	result, err := safeDivide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 3 =", result) // 3
	}

	// panicが発生するケース（ゼロ除算）
	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // recovered from panic: runtime error: integer divide by zero
	} else {
		fmt.Println("結果:", result)
	}

	// recoverしたのでプログラムは続行できる
	fmt.Println("プログラムは続行中")
}
```

> **💡 次のステップへ:** recoverの基本を学びました。次は「panicを使うべき場面」と「使うべきでない場面」を学びます。

### 応用: panicを使うべき場面

panicは「プログラムの設定が不正」「不変条件が満たされていない」など、プログラマのミスを示す場面で使います。

```go
package main

import (
	"fmt"
	"os"
)

// ---- panicが適切な場面 ----

type Config struct {
	DatabaseURL string
	Port        int
}

// MustLoadConfig は設定の読み込みに失敗したらpanicする
// 命名規則: Must で始まる関数はpanic可能性を示す
func MustLoadConfig() *Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// アプリケーション起動時に必須の設定がない → 続行不可 → panic
		panic("DATABASE_URL environment variable is required")
	}
	return &Config{DatabaseURL: dbURL, Port: 8080}
}

// MustParseTemplate はテンプレートのパースに失敗したらpanicする
func MustParseTemplate(name, content string) string {
	if content == "" {
		panic(fmt.Sprintf("template %q is empty", name))
	}
	return content
}

// ---- panicが不適切な場面 ----

// findUser でのpanicは不適切: ユーザーが見つからないのは通常の事象
// → error で返すべき
func findUser(id int) (string, error) {
	users := map[int]string{1: "田中太郎"}
	name, ok := users[id]
	if !ok {
		return "", fmt.Errorf("user not found: id=%d", id) // error で返す（panicしない）
	}
	return name, nil
}

func main() {
	// Must関数の使用例（環境変数が未設定だとpanic）
	// 実際にはプログラム起動時に呼ぶ
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("起動エラー:", r)
		}
	}()

	// panicが不適切な場面: 通常のエラーとして処理
	_, err := findUser(99)
	if err != nil {
		fmt.Println("通常のエラー:", err) // error で処理
	}

	// panicが適切な場面: 設定の読み込み失敗
	fmt.Println("設定を読み込みます...")
	_ = MustLoadConfig() // DATABASE_URL が未設定なら panic
}
```

> **💡 次のステップへ:** panicの使い分けを学びました。次はHTTPサーバーでのrecoverミドルウェアの実装パターンを学びます。

### 実践: HTTPミドルウェアでのrecover

実際のWebサーバーでpanic recoveryミドルウェアを実装するパターンを学びます。

```go
package main

import "fmt"

// ---- HTTPをシミュレートする簡易的な型 ----

type Request struct {
	Path   string
	Method string
}

type Response struct {
	StatusCode int
	Body       string
}

type Handler func(req *Request) *Response

// ---- recoveryミドルウェア ----

// withRecovery はpanicが発生しても500エラーを返すミドルウェア
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

// withLogging はリクエストログを出力するミドルウェア
func withLogging(next Handler) Handler {
	return func(req *Request) *Response {
		fmt.Printf("[LOG] %s %s\n", req.Method, req.Path)
		resp := next(req)
		fmt.Printf("[LOG] %s %s → %d\n", req.Method, req.Path, resp.StatusCode)
		return resp
	}
}

// ---- ハンドラー ----

func handleHello(req *Request) *Response {
	return &Response{StatusCode: 200, Body: "Hello, World!"}
}

func handlePanic(req *Request) *Response {
	// 意図しないpanicが発生するケース
	var data map[string]string
	_ = data["key"] // nil map からの読み込みはpanicしない（ゼロ値が返る）

	// これはpanicする
	panic("something unexpected happened!")
}

func main() {
	// ミドルウェアを適用: recovery → logging → handler
	helloHandler := withRecovery(withLogging(handleHello))
	panicHandler := withRecovery(withLogging(handlePanic))

	// 正常なリクエスト
	resp := helloHandler(&Request{Path: "/hello", Method: "GET"})
	fmt.Printf("Response: %d - %s\n\n", resp.StatusCode, resp.Body)

	// panicが発生するリクエスト（recoverで500が返る）
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

- [ ] panicは通常のエラーではなく、致命的な状況（設定不備、不変条件違反）に使う
- [ ] recoverはdeferの中でのみ動作する
- [ ] `Must`で始まる関数はpanicの可能性を示す慣習
- [ ] HTTPサーバーではrecoveryミドルウェアで500を返すのが標準パターン

**次のテーマ:** [11. encoding/jsonと構造体タグ](./11-encoding-json.md)
