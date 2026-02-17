# 05. Interfaceの応用 - 型アサーション、型スイッチ、空インターフェース

## 🎯 このテーマで学ぶこと

- 型アサーション（type assertion）でインターフェースから具象型を取り出す
- 型スイッチ（type switch）で型に応じた分岐処理
- 空インターフェース（`any`/`interface{}`）の使いどころと注意点
- インターフェースの合成（embedding）

**なぜ重要か:** インターフェースを実際に活用する場面では、「このインターフェースの中身は何型か？」を調べる必要が出てきます。JSONのパース結果の処理、ミドルウェアでの値取り出し、エラーの詳細取得など、実務で頻出するパターンです。

## 📖 概念

型アサーションはインターフェース値から具象型を取り出す操作です。型スイッチは複数の型に対して分岐する構文です。空インターフェース（`any`）は全ての型を受け入れますが、型安全性を失うため使いどころを限定すべきです。

**背景と設計意図:** Goはジェネリクスが後から追加された言語です。それ以前は`interface{}`が「任意の型」を表現する唯一の方法でした。現在はジェネリクスが使える場面では型パラメータを優先し、`any`は本当に任意の型を扱う場面に限定します。

**よくある誤解:**

- ❌ 「any をたくさん使えば柔軟になる」→ 型安全性を失うのでなるべく避ける
- ❌ 「型アサーションは安全」→ 失敗するとpanicになるので2値返却パターンを使う
- ❌ 「インターフェースの合成は継承」→ 合成は新しいインターフェースの定義であり継承ではない

## 💡 コード例

### 基本: 型アサーションと安全なチェック

インターフェース値から具象型を取り出す方法を学びます。

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

func describeAnimal(a Animal) {
	// 全ての Animal に共通の操作
	fmt.Println("鳴き声:", a.Speak())

	// 型アサーション（2値返却パターン）: 安全に具象型を取り出す
	// ok が false の場合、dog はゼロ値になる（panicしない）
	if dog, ok := a.(Dog); ok {
		fmt.Println(dog.Fetch()) // Dog固有のメソッドを呼べる
	}

	if cat, ok := a.(Cat); ok {
		fmt.Println(cat.Purr()) // Cat固有のメソッドを呼べる
	}
}

func main() {
	animals := []Animal{
		Dog{Name: "ポチ"},
		Cat{Name: "タマ"},
	}

	for _, a := range animals {
		describeAnimal(a)
		fmt.Println("---")
	}
}
```

> **💡 次のステップへ:** 型アサーションで個別の型をチェックする方法を学びました。型が多い場合は型スイッチを使うとすっきり書けます。

### 応用: 型スイッチ

複数の型に対して分岐処理を行う場合、型スイッチ（type switch）を使うと可読性が高くなります。

```go
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }
type Rectangle struct{ Width, Height float64 }
type Triangle struct{ Base, Height float64 }

func (c Circle) Area() float64    { return 3.14 * c.Radius * c.Radius }
func (r Rectangle) Area() float64 { return r.Width * r.Height }
func (t Triangle) Area() float64  { return t.Base * t.Height / 2 }

// describeShape は型スイッチで図形の種類に応じた説明を出力
func describeShape(s Shape) {
	// switch s.(type) で型に応じた分岐
	switch v := s.(type) {
	case Circle:
		fmt.Printf("円（半径: %.1f）面積: %.2f\n", v.Radius, v.Area())
	case Rectangle:
		fmt.Printf("長方形（%1.f x %.1f）面積: %.2f\n", v.Width, v.Height, v.Area())
	case Triangle:
		fmt.Printf("三角形（底辺: %.1f, 高さ: %.1f）面積: %.2f\n", v.Base, v.Height, v.Area())
	default:
		fmt.Printf("未知の図形: 面積 %.2f\n", s.Area())
	}
}

// processValues は空インターフェース（any）の型スイッチの例
func processValues(values []any) {
	for _, v := range values {
		switch val := v.(type) {
		case int:
			fmt.Printf("整数: %d (2倍: %d)\n", val, val*2)
		case string:
			fmt.Printf("文字列: %q (長さ: %d)\n", val, len(val))
		case bool:
			fmt.Printf("真偽値: %t\n", val)
		case nil:
			fmt.Println("nil値")
		default:
			fmt.Printf("その他の型: %T = %v\n", val, val)
		}
	}
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 3},
		Triangle{Base: 8, Height: 6},
	}

	for _, s := range shapes {
		describeShape(s)
	}

	fmt.Println("\n--- any の型スイッチ ---")
	values := []any{42, "hello", true, nil, 3.14}
	processValues(values)
}
```

> **💡 次のステップへ:** 型スイッチの使い方を学びました。次はインターフェースの合成と実務的な設計パターンを学びます。

### 実践: インターフェースの合成と実務パターン

複数の小さなインターフェースを合成して、柔軟な設計を行う方法を学びます。

```go
package main

import (
	"fmt"
	"strings"
)

// ---- 小さなインターフェースを定義 ----

type Reader interface {
	Read() (string, error)
}

type Writer interface {
	Write(data string) error
}

// ReadWriter はReaderとWriterを合成したインターフェース
// 標準ライブラリの io.ReadWriter と同じパターン
type ReadWriter interface {
	Reader
	Writer
}

// Closer はリソースを閉じるインターフェース
type Closer interface {
	Close() error
}

// ReadWriteCloser は3つのインターフェースを合成
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// ---- 具象型の実装 ----

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

// copyData はReaderからWriterにデータをコピーする
// 引数にインターフェースを取るため、どの実装でも渡せる
func copyData(r Reader, w Writer) error {
	data, err := r.Read()
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}
	return w.Write(data)
}

// processAndClose はReadWriteCloserを受け取って処理して閉じる
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

	// ReadWriteCloser として扱える
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

1. `PaymentMethod`インターフェース: `Pay(amount int) error`と`Name() string`メソッドを持つ
2. `CreditCard`構造体: `Number string`, `Limit int`フィールド。Limit超過でエラー
3. `BankTransfer`構造体: `AccountNumber string`, `Balance int`フィールド。残高不足でエラー
4. `ProcessPayment(method PaymentMethod, amount int) string`: 型スイッチで支払い方法に応じたメッセージを返す
5. `PaymentLogger`インターフェース: `PaymentMethod`を埋め込み、`Log() string`メソッドを追加

**期待される動作:**

- `CreditCard{Limit: 10000}.Pay(5000)` → 成功
- `CreditCard{Limit: 10000}.Pay(15000)` → 限度額超過エラー
- `ProcessPayment`が型スイッチでカード番号や口座番号を含むメッセージを返す

## ✅ 重要ポイント

- [ ] 型アサーションは必ず2値返却パターン（`v, ok := i.(Type)`）を使う
- [ ] 型スイッチは複数の型に対する分岐で可読性が高い
- [ ] `any`の使用は最小限にし、可能ならジェネリクスを検討する
- [ ] インターフェースの合成で、小さなインターフェースを組み合わせる

**次のテーマ:** [06. エラーハンドリングの基本](./06-error-handling-basics.md)
