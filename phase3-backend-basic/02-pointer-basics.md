# 02. ポインタの基本 - 参照と逆参照、値渡しとポインタ渡し

## 🎯 このテーマで学ぶこと

- ポインタの基本（`&`と`*`の使い方）
- 値渡しとポインタ渡しの違い
- ポインタを使うべき場面の判断基準

## 📖 なぜポインタを理解する必要があるのか

Goは値渡しがデフォルトの言語です。構造体を関数に渡すと**コピーが作られる**ため、関数内で変更しても呼び出し元には反映されません。

### こう書かないとどうなるか

```go
func updateAge(u User) {
    u.Age = 31  // これはコピーを変更しているだけ
}

user := User{Name: "田中", Age: 30}
updateAge(user)
fmt.Println(user.Age) // 30 ← 変わっていない！
```

「変更したのに反映されない」というバグの原因はほぼこれです。変更を呼び出し元に反映するにはポインタ渡しが必要です。

### ポインタのもう一つの用途：「値がない」の表現

Goの`string`のゼロ値は空文字`""`、`int`のゼロ値は`0`です。しかし「空文字」と「未設定」は意味が違います。`*string`（ポインタ型）にすることで、`nil`=未設定、`""`=空文字と区別できます。データベースのNULL値やJSONのnullを扱う場面で必須です。

## 💡 コード例

### 基本: ポインタの基礎と値渡し vs ポインタ渡し

`&`でアドレスを取得し、`*`で値にアクセスする基本操作と、値渡し・ポインタ渡しの違いを学びます。

```go
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 値渡し：コピーを受け取るので、元の値は変わらない
func incrementByValue(u User) {
	u.Age++
	fmt.Println("関数内:", u.Age)
}

// ポインタ渡し：アドレスを受け取るので、元の値を変更できる
// なぜ*Userか → 「この関数はUserを変更する」ことを型で表現している
func incrementByPointer(u *User) {
	u.Age++
	fmt.Println("関数内:", u.Age)
}

func main() {
	// --- ポインタの基礎 ---
	x := 42
	p := &x          // &でアドレスを取得
	fmt.Println(*p)  // *で値にアクセス → 42
	*p = 100         // ポインタ経由で値を変更
	fmt.Println(x)   // 100（元の変数も変わる）

	// ポインタのゼロ値はnil
	var q *int
	fmt.Println(q)   // <nil>
	// *q とするとpanicになるので、必ずnilチェックする
	if q != nil {
		fmt.Println(*q)
	}

	// --- 値渡し vs ポインタ渡し ---
	user := User{Name: "田中太郎", Age: 30}

	incrementByValue(user)
	fmt.Println("値渡し後:", user.Age)    // 30（変わらない）

	incrementByPointer(&user)           // &でアドレスを渡す
	fmt.Println("ポインタ渡し後:", user.Age) // 31（変わった）
}
```

### 実践: nilで「値がない」を表現する

APIレスポンスやDBのNULL値を扱う場面では、ポインタ型で「値があるか・ないか」を区別します。

```go
package main

import (
	"encoding/json"
	"fmt"
)

// UserProfile はAPIレスポンスを想定した構造体
// ポインタ型フィールドでnull/未設定を表現
type UserProfile struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Nickname *string `json:"nickname"` // nullの可能性がある → ポインタ型
	Age      *int    `json:"age"`      // nullの可能性がある → ポインタ型
}

func (u *UserProfile) UpdateProfile(name, email string) {
	u.Name = name
	u.Email = email
}

// ヘルパー関数：リテラル値から直接ポインタを作れないGoの制約を回避
func stringPtr(s string) *string { return &s }

func main() {
	// JSONのnullはポインタ型ならnilに、非ポインタ型ならゼロ値になる
	jsonData := `{"name": "田中太郎", "email": "tanaka@example.com", "nickname": null, "age": 30}`

	var profile UserProfile
	if err := json.Unmarshal([]byte(jsonData), &profile); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Name: %s\n", profile.Name)

	// ポインタ型フィールドは必ずnilチェックしてからアクセス
	// nilのまま*profile.Nicknameとするとpanicになる
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

1. `BankAccount`構造体: `Owner string`, `Balance int`を持つ
2. `Deposit(amount int) error`: 入金処理（負の値はエラー）。呼び出し元のBalanceが変わること
3. `Withdraw(amount int) error`: 出金処理（残高不足・負の値はエラー）
4. `Transfer(to *BankAccount, amount int) error`: 送金処理。自分から引いて相手に足す

**ヒント:**

```go
// ポインタレシーバーで元の値を変更
func (a *BankAccount) Deposit(amount int) error {
	// ...
}
```

## ✅ 重要ポイント

- [ ] `&`でアドレス取得、`*`で値にアクセス
- [ ] 値渡しはコピー、ポインタ渡しは元の値を変更できる
- [ ] ポインタ型（`*string`等）でnull/未設定を表現する
- [ ] nilポインタにアクセスするとpanicするので必ずチェック

**次のテーマ:** [03. メソッドとレシーバー](./03-methods-and-receivers.md)
