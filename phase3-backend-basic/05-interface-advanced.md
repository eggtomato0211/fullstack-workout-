# 05. Interfaceの応用 - 型アサーション、型スイッチ、空インターフェース

## 🎯 このテーマで学ぶこと

- 型アサーションでインターフェースから具象型を取り出す
- 型スイッチで型に応じた分岐処理
- 空インターフェース（`any`）の注意点とインターフェースの合成

## 📖 なぜ型アサーションと型スイッチを理解する必要があるのか

インターフェースは便利ですが、「このインターフェースの中身は実際に何型なのか？」を知りたい場面が出てきます。JSONのパース結果、エラーの詳細取得、ミドルウェアでの値取り出しなどです。

### こう書かないとどうなるか

```go
// 型アサーションを1値で受け取ると、失敗時にpanicする
dog := animal.(Dog) // animalがDogでなければ → panic!

// 安全な方法：2値返却パターン
dog, ok := animal.(Dog) // okがfalseならdogはゼロ値（panicしない）
```

型アサーションは**必ず2値返却パターン**を使うのが鉄則です。

### `any`は最後の手段

`any`（= `interface{}`）は全ての型を受け入れますが、型安全性を失います。Go 1.18以降はジェネリクスが使えるため、「任意の型」が必要な場面ではまずジェネリクスを検討し、`any`は本当に型が不定な場面（JSONの動的パースなど）に限定します。

## 💡 コード例

### 基本: 型アサーションと型スイッチ

インターフェースから具象型を安全に取り出す方法と、型スイッチで複数の型に分岐する方法を学びます。

```go
package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{ Name string }
type Cat struct{ Name string }

func (d Dog) Speak() string { return "ワン！" }
func (d Dog) Fetch() string { return d.Name + "がボールを取ってきた！" }
func (c Cat) Speak() string { return "ニャー！" }
func (c Cat) Purr() string  { return c.Name + "がゴロゴロ言っている" }

// 型アサーション：2値返却パターンで安全に取り出す
func describeAnimal(a Animal) {
	fmt.Println("鳴き声:", a.Speak())

	// okがfalseならdogはゼロ値（panicしない）
	if dog, ok := a.(Dog); ok {
		fmt.Println(dog.Fetch()) // Dog固有のメソッド
	}
	if cat, ok := a.(Cat); ok {
		fmt.Println(cat.Purr()) // Cat固有のメソッド
	}
}

// 型スイッチ：型が多い場合はこちらの方が可読性が高い
func processValue(v any) {
	switch val := v.(type) {
	case int:
		fmt.Printf("整数: %d (2倍: %d)\n", val, val*2)
	case string:
		fmt.Printf("文字列: %q (長さ: %d)\n", val, len(val))
	case bool:
		fmt.Printf("真偽値: %t\n", val)
	default:
		fmt.Printf("その他: %T = %v\n", val, val)
	}
}

func main() {
	animals := []Animal{Dog{Name: "ポチ"}, Cat{Name: "タマ"}}
	for _, a := range animals {
		describeAnimal(a)
		fmt.Println("---")
	}

	fmt.Println("--- 型スイッチ ---")
	values := []any{42, "hello", true, 3.14}
	for _, v := range values {
		processValue(v)
	}
}
```

### 実践: インターフェースの合成

小さなインターフェースを組み合わせて、柔軟で再利用性の高い設計を行います。標準ライブラリの`io.ReadWriter`と同じパターンです。

```go
package main

import (
	"fmt"
	"strings"
)

// 小さなインターフェースを定義
type Reader interface {
	Read() (string, error)
}

type Writer interface {
	Write(data string) error
}

// 合成：Reader + Writer = ReadWriter
// 標準ライブラリの io.ReadWriter と同じパターン
type ReadWriter interface {
	Reader
	Writer
}

type Closer interface {
	Close() error
}

// 3つを合成
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// --- 具象型 ---

type StringBuffer struct {
	data   strings.Builder
	closed bool
}

func (b *StringBuffer) Read() (string, error) {
	if b.closed {
		return "", fmt.Errorf("buffer is closed")
	}
	return b.data.String(), nil
}

func (b *StringBuffer) Write(data string) error {
	if b.closed {
		return fmt.Errorf("buffer is closed")
	}
	b.data.WriteString(data)
	return nil
}

func (b *StringBuffer) Close() error {
	b.closed = true
	return nil
}

// 引数にインターフェースを取る → 実装に依存しない
func copyData(r Reader, w Writer) error {
	data, err := r.Read()
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}
	return w.Write(data)
}

func processAndClose(rwc ReadWriteCloser) error {
	if err := rwc.Write("処理済みデータ"); err != nil {
		return err
	}
	data, err := rwc.Read()
	if err != nil {
		return err
	}
	fmt.Println("データ:", data)
	return rwc.Close()
}

func main() {
	buf := &StringBuffer{}
	if err := processAndClose(buf); err != nil {
		fmt.Println("Error:", err)
	}

	// Close後は操作できない
	_, err := buf.Read()
	if err != nil {
		fmt.Println("Expected error:", err) // buffer is closed
	}
}
```

## 🎯 演習問題

支払いシステムをインターフェースと型スイッチで設計してください。

**要件:**

1. `PaymentMethod`インターフェース: `Pay(amount int) error`と`Name() string`
2. `CreditCard`構造体: `Number string`, `Limit int`。Limit超過でエラー
3. `BankTransfer`構造体: `AccountNumber string`, `Balance int`。残高不足でエラー
4. `ProcessPayment(method PaymentMethod, amount int) string`: 型スイッチで支払い方法に応じたメッセージを返す

## ✅ 重要ポイント

- [ ] 型アサーションは必ず2値返却パターン（`v, ok := i.(Type)`）を使う
- [ ] 型スイッチは複数の型に対する分岐で可読性が高い
- [ ] `any`の使用は最小限にし、ジェネリクスを優先する
- [ ] インターフェースの合成で、小さなインターフェースを組み合わせる

**次のテーマ:** [06. エラーハンドリングの基本](./06-error-handling-basics.md)
