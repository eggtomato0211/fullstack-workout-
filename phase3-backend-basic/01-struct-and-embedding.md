# 01. 構造体と埋め込み - struct定義、初期化、擬似継承

## 🎯 このテーマで学ぶこと

- 構造体（struct）の定義と初期化の複数パターン
- フィールドの可視性（大文字/小文字）の意味
- 構造体の埋め込み（embedding）による擬似継承
- コンストラクタ関数パターン
- 値型としての構造体の特性

**なぜ重要か:** Goにはクラスがありません。構造体がGoにおけるデータ構造の基本単位であり、メソッド、インターフェース、埋め込みと組み合わせることでオブジェクト指向的な設計を実現します。構造体の理解なくしてGoのプログラミングは成り立ちません。

## 📖 概念

構造体は複数のフィールドをまとめたデータ型です。Goでは構造体を起点にメソッドを定義し、インターフェースを満たすことで多態性を実現します。

**背景と設計意図:** JavaやPythonのような言語ではクラスが基本ですが、Goはシンプルさを重視し、クラスの代わりに構造体+メソッド+インターフェースという組み合わせを採用しました。継承の代わりに「埋め込み（embedding）」を使い、コンポジション（合成）でコードを再利用します。

**実務での活用場面:** APIのリクエスト/レスポンスの型定義、データベースのレコードマッピング、ドメインモデルの定義など、あらゆる場面で構造体を使います。

**よくある誤解:**

- ❌ 「埋め込みは継承と同じ」→ 埋め込みはフィールドの昇格（promotion）であり、is-a関係ではなくhas-a関係
- ❌ 「構造体はポインタで渡すべき」→ 小さい構造体は値渡しの方が効率的な場合もある
- ❌ 「ゼロ値は使えない」→ Goの構造体はゼロ値でも有効な状態になるよう設計するのがベストプラクティス

## 💡 コード例

### 基本: 構造体の定義と初期化

まずは構造体の定義方法と、複数の初期化パターンを学びます。Goの構造体はゼロ値で初期化されるため、`nil`チェックが不要な堅牢な設計ができます。

```go
package main

import "fmt"

// User は利用者を表す構造体
// フィールド名が大文字 → パッケージ外からアクセス可能（exported）
// フィールド名が小文字 → パッケージ内のみアクセス可能（unexported）
type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	// 初期化パターン1: フィールド名を指定（推奨）
	// → 順序に依存せず、可読性が高い
	u1 := User{
		Name:  "田中太郎",
		Email: "tanaka@example.com",
		Age:   30,
	}

	// 初期化パターン2: ゼロ値で初期化
	// string→"", int→0, bool→false, pointer→nil
	var u2 User
	u2.Name = "鈴木花子"
	u2.Email = "suzuki@example.com"

	// 初期化パターン3: new()でポインタを取得
	u3 := new(User)
	u3.Name = "佐藤一郎"

	fmt.Printf("u1: %+v\n", u1)  // {Name:田中太郎 Email:tanaka@example.com Age:30}
	fmt.Printf("u2: %+v\n", u2)  // {Name:鈴木花子 Email:suzuki@example.com Age:0}
	fmt.Printf("u3: %+v\n", *u3) // {Name:佐藤一郎 Email: Age:0}
}
```

> **💡 次のステップへ:** 基本では構造体の定義と初期化を学びました。次の応用では、実務で必須の**コンストラクタパターン**と**バリデーション付き初期化**を学びます。

### 応用: コンストラクタ関数パターン

Goにはコンストラクタ構文がありませんが、`New`で始まる関数を作る慣習があります。バリデーションやデフォルト値の設定を初期化時に行うことで、不正な状態の構造体が作られることを防ぎます。

```go
package main

import (
	"errors"
	"fmt"
)

type User struct {
	name  string // 小文字 → 外部パッケージから直接アクセス不可
	email string
	age   int
}

// NewUser はUserのコンストラクタ関数
// バリデーションを通過した場合のみUserを返す
func NewUser(name, email string, age int) (*User, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if age < 0 || age > 150 {
		return nil, errors.New("age must be between 0 and 150")
	}
	return &User{
		name:  name,
		email: email,
		age:   age,
	}, nil
}

// Getter: 外部から読み取り専用でアクセスできるようにする
func (u *User) Name() string  { return u.name }
func (u *User) Email() string { return u.email }
func (u *User) Age() int      { return u.age }

func main() {
	user, err := NewUser("田中太郎", "tanaka@example.com", 30)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Name: %s, Email: %s, Age: %d\n", user.Name(), user.Email(), user.Age())

	// バリデーションエラーのケース
	_, err = NewUser("", "test@example.com", 25)
	if err != nil {
		fmt.Println("Error:", err) // Error: name is required
	}
}
```

> **💡 次のステップへ:** コンストラクタパターンで安全な初期化ができるようになりました。次の実践では、**埋め込み（embedding）**を使って構造体を組み合わせ、コードの再利用を実現する方法を学びます。

### 実践: 構造体の埋め込み（embedding）

埋め込みはGoの「継承の代わり」となる機能です。構造体に別の構造体を埋め込むと、埋め込まれた構造体のフィールドやメソッドが「昇格」し、あたかも自分のフィールド・メソッドのようにアクセスできます。

```go
package main

import (
	"fmt"
	"time"
)

// Timestamps は作成日時・更新日時の共通フィールド
// 複数のモデルで使い回すための構造体
type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Touch は更新日時を現在時刻に更新するメソッド
func (t *Timestamps) Touch() {
	t.UpdatedAt = time.Now()
}

// User は Timestamps を埋め込んでいる
type User struct {
	Timestamps        // 埋め込み（フィールド名なし）
	ID         int
	Name       string
	Email      string
}

// Article も Timestamps を埋め込んでいる
type Article struct {
	Timestamps        // 同じTimestampsを再利用
	ID         int
	Title      string
	Body       string
	AuthorID   int
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

	// 埋め込みにより、Timestamps のフィールドに直接アクセス可能
	fmt.Println("Created:", user.CreatedAt)

	// 埋め込みにより、Timestamps のメソッドも直接呼び出し可能
	user.Touch()
	fmt.Println("Updated:", user.UpdatedAt)

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

	// Article でも同じように Timestamps のフィールドとメソッドを利用できる
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

**期待される動作:**

- `NewProduct("Goの本", 3000, 10)`が正常に作成される
- `NewProduct("", 3000, 10)`がエラーを返す
- `product.IsInStock()`が`true`/`false`を返す
- `product.CreatedAt`にBaseModelのフィールドとして直接アクセスできる

## ✅ 重要ポイント

- [ ] 構造体のフィールド名の大文字/小文字で可視性が決まる
- [ ] コンストラクタ関数（`NewXxx`）でバリデーション付き初期化を行う
- [ ] 埋め込みは継承ではなくコンポジション（has-a関係）
- [ ] ゼロ値が有効な状態になるよう構造体を設計する

**次のテーマ:** [02. ポインタの基本](./02-pointer-basics.md)
