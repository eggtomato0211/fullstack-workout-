# 01. 構造体と埋め込み - struct定義、初期化、擬似継承

## 🎯 このテーマで学ぶこと

- 構造体（struct）の定義と初期化パターン
- コンストラクタ関数によるバリデーション付き初期化
- 構造体の埋め込み（embedding）によるコード再利用

## 📖 なぜ構造体を理解する必要があるのか

Goにはクラスがありません。「じゃあデータと振る舞いをどうまとめるの？」という疑問に対するGoの答えが**構造体**です。

他の言語（Java、Python等）では`class User`と書きますが、Goでは`type User struct`と書きます。やりたいことは同じ、「関連するデータをひとまとめにして、そのデータに対する操作を定義する」ことです。

### こう書かないとどうなるか

構造体を使わずに、バラバラの変数でユーザー情報を管理してみましょう：

```go
// 構造体を使わない場合
userName := "田中太郎"
userEmail := "tanaka@example.com"
userAge := 30

// ユーザーが増えると変数が爆発する
userName2 := "鈴木花子"
userEmail2 := "suzuki@example.com"
userAge2 := 25

// 関数に渡すときも引数がどんどん増える
func createUser(name string, email string, age int, role string, ...) { }
```

これでは「どの変数がどのユーザーに属するか」がコードから読み取れません。構造体は「このデータはセットで1つの意味を持つ」ということをコードで表現する仕組みです。

### フィールドの大文字・小文字が意味を持つ理由

Goでは**フィールド名の先頭が大文字なら外部パッケージからアクセス可能**、**小文字なら不可**というルールがあります。

なぜこんな仕組みがあるのか？　たとえば`User`構造体の`password`フィールドを外部から直接書き換えられたら困りますよね。「外から見える/見えない」をフィールド名だけで制御できるのがGoのシンプルな設計です。他の言語の`public`/`private`キーワードに相当しますが、Goでは名前の規則だけで実現しています。

### ゼロ値という安全装置

Goの構造体は宣言しただけで全フィールドがゼロ値（string→`""`、int→`0`、bool→`false`、pointer→`nil`）に初期化されます。「初期化し忘れて不定値が入っていた」というバグが原理的に起きません。これはGoが意図的に設計した安全策です。

## 💡 コード例

### 基本: 構造体の定義・初期化とコンストラクタ

構造体の定義方法と、実務で必須の**コンストラクタパターン**を合わせて学びます。Goにはコンストラクタ構文がありませんが、`NewXxx`という命名の関数を作る慣習があります。

なぜコンストラクタが必要なのか？　構造体リテラルで直接初期化すると、不正な値（名前が空、年齢がマイナス等）でもそのまま作れてしまいます。コンストラクタでバリデーションを挟むことで、「不正な状態のオブジェクトが存在しない」ことを保証できます。

```go
package main

import (
	"errors"
	"fmt"
)

// User はアプリケーションの利用者を表す
// Name, Email → 大文字なので外部パッケージからアクセス可能
type User struct {
	Name  string
	Email string
	age   int // 小文字 → 外部パッケージから直接アクセス不可
}

// NewUser はUserのコンストラクタ
// なぜポインタを返すのか → 構造体のコピーを避けるため（後のテーマで詳しく学ぶ）
// なぜerrorも返すのか → 不正な値で作れないようにするため
func NewUser(name, email string, age int) (*User, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if age < 0 || age > 150 {
		return nil, errors.New("age must be between 0 and 150")
	}
	return &User{
		Name:  name,
		Email: email,
		age:   age,
	}, nil
}

// Age はageフィールドのGetter
// なぜわざわざGetterを書くのか → ageを小文字にして外部から直接変更できなくし、
// 読み取りだけを許可するため（書き込みは NewUser 経由のバリデーションを通す）
func (u *User) Age() int { return u.age }

func main() {
	// コンストラクタ経由で作成（バリデーション付き）
	user, err := NewUser("田中太郎", "tanaka@example.com", 30)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Name: %s, Age: %d\n", user.Name, user.Age())

	// バリデーションエラーのケース
	_, err = NewUser("", "test@example.com", 25)
	fmt.Println("Error:", err) // Error: name is required

	// ゼロ値での初期化も可能（コンストラクタを通さない場合）
	var u2 User
	fmt.Printf("ゼロ値: %+v\n", u2) // {Name: Email: age:0}
}
```

### 実践: 構造体の埋め込み（embedding）

「複数の構造体に共通するフィールドがある」場合、全部に同じフィールドを書くのは面倒ですし、修正時に全箇所を直す必要があります。

埋め込みを使えば、共通部分を1つの構造体に切り出して再利用できます。これはGoの「継承の代わり」ですが、重要な違いがあります：**埋め込みはis-a（〜である）ではなくhas-a（〜を持つ）の関係**です。UserはTimestampsを「継承」しているのではなく、Timestampsを「内包」しています。

```go
package main

import (
	"fmt"
	"time"
)

// Timestamps は作成日時・更新日時の共通フィールド
// User、Article、Comment... あらゆるモデルで必要になるフィールドを1箇所にまとめる
type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Touch は更新日時を現在時刻に更新する
func (t *Timestamps) Touch() {
	t.UpdatedAt = time.Now()
}

// User は Timestamps を埋め込んでいる
// 埋め込み = フィールド名なしで構造体を書く
type User struct {
	Timestamps        // ← これが埋め込み
	ID         int
	Name       string
	Email      string
}

// Article も同じ Timestamps を再利用
type Article struct {
	Timestamps
	ID       int
	Title    string
	Body     string
	AuthorID int
}

func main() {
	user := User{
		ID:    1,
		Name:  "田中太郎",
		Email: "tanaka@example.com",
		Timestamps: Timestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// 埋め込みの効果①: Timestampsのフィールドに直接アクセスできる
	// user.Timestamps.CreatedAt と書かなくていい
	fmt.Println("Created:", user.CreatedAt)

	// 埋め込みの効果②: Timestampsのメソッドも直接呼べる
	user.Touch()
	fmt.Println("Updated:", user.UpdatedAt)

	// ArticleでもUserと全く同じ使い方ができる
	// → Timestampsの実装を変えれば、全モデルに反映される
	article := Article{
		ID:       1,
		Title:    "Go入門",
		Body:     "Goの構造体について学びます",
		AuthorID: user.ID,
		Timestamps: Timestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	article.Touch()
	fmt.Printf("Article: %s (updated: %v)\n", article.Title, article.UpdatedAt)
}
```

## 🎯 演習問題

ECサイトの商品管理を想定した構造体を設計してください。

**要件:**

1. `BaseModel`構造体: `ID int`, `CreatedAt time.Time`, `UpdatedAt time.Time`フィールドを持つ
2. `Product`構造体: `BaseModel`を埋め込み、`Name string`, `Price int`, `Stock int`を持つ
3. `NewProduct`コンストラクタ: 名前が空、価格が0以下の場合にエラーを返す
4. `Product`に`IsInStock() bool`メソッド: 在庫があるかを返す
5. `Product`に`String() string`メソッド: 商品情報を文字列で返す

**ヒント:**

```go
type BaseModel struct {
	// 共通フィールドを定義
}

type Product struct {
	BaseModel // 埋め込み
	// 商品固有のフィールドを定義
}

func NewProduct(name string, price, stock int) (*Product, error) {
	// バリデーション + 初期化
}
```

## ✅ 重要ポイント

- [ ] フィールド名の大文字/小文字で可視性が決まる（public/privateの代わり）
- [ ] コンストラクタ関数（`NewXxx`）で不正な状態のオブジェクトを防ぐ
- [ ] 埋め込みは継承ではなくコンポジション（has-a関係）
- [ ] ゼロ値が有効な状態になるよう構造体を設計する

**次のテーマ:** [02. ポインタの基本](./02-pointer-basics.md)
