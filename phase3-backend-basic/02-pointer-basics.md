# 02. ポインタの基本 - 参照と逆参照、値渡しとポインタ渡し

## 🎯 このテーマで学ぶこと

- ポインタの基本（`&`と`*`の使い方）
- 値渡しとポインタ渡しの違い
- ポインタを使うべき場面の判断基準
- nilポインタの扱いと安全なアクセス

**なぜ重要か:** Goは値渡しがデフォルトの言語です。構造体を関数に渡すとコピーが作られるため、「変更を呼び出し元に反映したい」「大きな構造体のコピーを避けたい」場面ではポインタが必須です。メソッドのレシーバーやJSON Unmarshalなど、実務で頻繁に使います。

## 📖 概念

ポインタは値が格納されているメモリアドレスを保持する変数です。`&`で変数のアドレスを取得し、`*`でアドレスが指す値にアクセスします。

**背景と設計意図:** CやC++のようなポインタ演算はGoにはありませんが、参照の概念は残しています。これにより、不必要なコピーを避けつつ、ポインタ演算の危険性を排除しています。

**実務での活用場面:** 構造体のメソッドレシーバー、関数への構造体の渡し方、JSONのデコード先指定、データベースのNULL表現（`*string`）など。

**よくある誤解:**

- ❌ 「常にポインタを使うべき」→ 小さい構造体やプリミティブは値渡しの方が効率的
- ❌ 「ポインタ = 参照型」→ Goのポインタは明示的で、暗黙の参照渡しはない（map/slice/channelは例外）
- ❌ 「nilポインタは怖い」→ 適切にチェックすれば安全に扱える

## 💡 コード例

### 基本: ポインタの基礎

`&`（アドレス演算子）と`*`（間接参照演算子）の基本的な使い方を学びます。

```go
package main

import "fmt"

func main() {
	x := 42

	// &x で x のメモリアドレスを取得
	p := &x
	fmt.Println("xの値:", x)       // 42
	fmt.Println("xのアドレス:", p)   // 0xc000018088（実行時に異なる）
	fmt.Println("pが指す値:", *p)   // 42

	// *p でポインタが指す値を変更 → 元の x も変わる
	*p = 100
	fmt.Println("xの値:", x) // 100

	// ポインタのゼロ値は nil
	var q *int
	fmt.Println("qの値:", q) // <nil>

	// nilポインタへのアクセスはpanicになるのでチェックが必要
	if q != nil {
		fmt.Println(*q)
	} else {
		fmt.Println("qはnilです")
	}
}
```

> **💡 次のステップへ:** ポインタの基本操作を学びました。次は関数に値を渡す際の「値渡し」と「ポインタ渡し」の違いを理解します。

### 応用: 値渡しとポインタ渡し

Goの関数は引数を常にコピーして渡します（値渡し）。ポインタを渡すことで、関数内での変更を呼び出し元に反映させることができます。

```go
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 値渡し: 引数のコピーが作られる → 元の値は変わらない
func incrementAgeByValue(u User) {
	u.Age++ // コピーのAgeを変更しているだけ
	fmt.Println("関数内:", u.Age)
}

// ポインタ渡し: アドレスを渡す → 元の値を変更できる
func incrementAgeByPointer(u *User) {
	u.Age++ // 元のUserのAgeを変更
	fmt.Println("関数内:", u.Age)
}

func main() {
	user := User{Name: "田中太郎", Age: 30}

	// 値渡し: 関数内で変更しても元のuserは変わらない
	incrementAgeByValue(user)
	fmt.Println("値渡し後:", user.Age) // 30（変わっていない）

	// ポインタ渡し: 関数内の変更が元のuserに反映される
	incrementAgeByPointer(&user)
	fmt.Println("ポインタ渡し後:", user.Age) // 31（変わった）
}
```

> **💡 次のステップへ:** 値渡しとポインタ渡しの違いを理解しました。次は実務的な場面で「ポインタを使うべき場面」を判断する基準を学びます。

### 実践: ポインタの使いどころ

実務でのポインタの使い方として、「値がないこと（nil）を表現する」パターンと、「大きな構造体のコピーを避ける」パターンを学びます。

```go
package main

import (
	"encoding/json"
	"fmt"
)

// UserProfile はAPIレスポンスを想定した構造体
// ポインタ型フィールドで「値がない（null）」を表現できる
type UserProfile struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Nickname *string `json:"nickname"` // nullの可能性がある → ポインタ型
	Age      *int    `json:"age"`      // nullの可能性がある → ポインタ型
}

// UpdateProfile はプロフィールを更新する
// 構造体が大きい場合やフィールドを変更する場合はポインタレシーバー
func (u *UserProfile) UpdateProfile(name, email string) {
	u.Name = name
	u.Email = email
}

// ヘルパー: stringのポインタを返す
func stringPtr(s string) *string { return &s }
func intPtr(i int) *int          { return &i }

func main() {
	// JSONからのデコード: nilフィールドの活用
	jsonData := `{"name": "田中太郎", "email": "tanaka@example.com", "nickname": null, "age": 30}`

	var profile UserProfile
	if err := json.Unmarshal([]byte(jsonData), &profile); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Name: %s\n", profile.Name)

	// ポインタ型フィールドはnilチェックが必要
	if profile.Nickname != nil {
		fmt.Printf("Nickname: %s\n", *profile.Nickname)
	} else {
		fmt.Println("Nickname: (未設定)")
	}

	if profile.Age != nil {
		fmt.Printf("Age: %d\n", *profile.Age)
	}

	// ヘルパー関数でポインタ型フィールドに値を設定
	profile.Nickname = stringPtr("たなちゃん")
	fmt.Printf("Updated Nickname: %s\n", *profile.Nickname)
}
```

## 🎯 演習問題

銀行口座を管理する構造体と関数を設計してください。

**要件:**

1. `BankAccount`構造体: `Owner string`, `Balance int`（残高）を持つ
2. `Deposit(amount int) error`: 入金処理（負の値はエラー）。呼び出し元のBalanceが変わること
3. `Withdraw(amount int) error`: 出金処理（残高不足・負の値はエラー）。呼び出し元のBalanceが変わること
4. `Transfer(to *BankAccount, amount int) error`: 送金処理。自分から引いて相手に足す

**ヒント:**

```go
type BankAccount struct {
	Owner   string
	Balance int
}

func (a *BankAccount) Deposit(amount int) error {
	// ポインタレシーバーで元の値を変更
}
```

**期待される動作:**

- `Deposit(1000)` → 残高が1000増える
- `Withdraw(500)` → 残高が500減る
- `Withdraw(10000)` → 残高不足エラー
- `Transfer(&otherAccount, 300)` → 自分-300、相手+300

## ✅ 重要ポイント

- [ ] `&`でアドレス取得、`*`で値にアクセス
- [ ] 値渡し：コピーが作られる。ポインタ渡し：元の値を変更できる
- [ ] ポインタ型フィールド（`*string`等）でnull/未設定を表現する
- [ ] nilポインタにアクセスするとpanicになるので必ずチェックする

**次のテーマ:** [03. メソッドとレシーバー](./03-methods-and-receivers.md)
