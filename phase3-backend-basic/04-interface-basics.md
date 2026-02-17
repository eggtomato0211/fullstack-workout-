# 04. Interfaceの基本 - interface定義、暗黙的な実装

## 🎯 このテーマで学ぶこと

- インターフェースの定義方法
- 暗黙的な実装（implements不要）の仕組み
- インターフェースを引数に取る関数の設計
- 標準ライブラリでのインターフェースの活用例

**なぜ重要か:** インターフェースはGoの最も強力な設計ツールです。暗黙的な実装により、依存性の逆転（DIP）や疎結合な設計を自然に実現できます。テスト時のモック差し替え、ストレージの切り替えなど、実務で必須の概念です。

## 📖 概念

インターフェースはメソッドの集合を定義する型です。Goのインターフェースは**暗黙的に実装**されます。つまり、必要なメソッドをすべて持っていれば、`implements`と宣言しなくても自動的にそのインターフェースを満たします。

**背景と設計意図:** JavaやC#では`implements`キーワードで明示的にインターフェースを実装しますが、Goは「ダックタイピング」を採用しました。「アヒルのように歩き、アヒルのように鳴くなら、それはアヒルだ」という考え方で、型の柔軟性と疎結合を実現します。

**実務での活用場面:** リポジトリパターン（DBをモックに差し替え）、ロガーの差し替え、外部API クライアントのモック化など。

**よくある誤解:**

- ❌ 「インターフェースは大きく定義する」→ 小さいインターフェース（1-2メソッド）が推奨（Go Proverbs）
- ❌ 「先にインターフェースを定義する」→ 必要になった時に定義する（Accept interfaces, return structs）
- ❌ 「implements宣言が必要」→ Goでは不要。メソッドが合えば自動的に満たす

## 💡 コード例

### 基本: インターフェースの定義と暗黙的な実装

インターフェースを定義し、複数の型で実装する基本パターンを学びます。

```go
package main

import (
	"fmt"
	"math"
)

// Shape は図形を表すインターフェース
// メソッドシグネチャのみを定義（実装は持たない）
type Shape interface {
	Area() float64
	Perimeter() float64
}

// ---- Rectangle は Shape を「暗黙的に」実装 ----

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// ---- Circle も Shape を「暗黙的に」実装 ----

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// printShapeInfo は Shape インターフェースを受け取る
// → Rectangle でも Circle でも、Shape を満たす任意の型を渡せる
func printShapeInfo(s Shape) {
	fmt.Printf("面積: %.2f, 周長: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	// どちらも Shape インターフェースを満たしているので渡せる
	printShapeInfo(rect)   // 面積: 50.00, 周長: 30.00
	printShapeInfo(circle) // 面積: 153.94, 周長: 43.98
}
```

> **💡 次のステップへ:** 基本的なインターフェースの定義と実装を学びました。次は実務的な「リポジトリパターン」でインターフェースを活用する方法を学びます。

### 応用: インターフェースによる依存性の逆転

インターフェースを使って具象型への依存を減らし、テスト時にモックに差し替えられる設計を学びます。

```go
package main

import "fmt"

// UserRepository はユーザーデータの永続化を抽象化するインターフェース
// メソッドが少ない（2つ）→ Goらしい小さなインターフェース
type UserRepository interface {
	FindByID(id int) (*User, error)
	Save(user *User) error
}

type User struct {
	ID   int
	Name string
}

// ---- 本番用: メモリ上にデータを保持する実装 ----

type InMemoryUserRepository struct {
	users map[int]*User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[int]*User)}
}

func (r *InMemoryUserRepository) FindByID(id int) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found: id=%d", id)
	}
	return user, nil
}

func (r *InMemoryUserRepository) Save(user *User) error {
	r.users[user.ID] = user
	return nil
}

// ---- UserService はインターフェースに依存 ----
// 具象型（InMemoryUserRepository）ではなくインターフェースに依存することで
// テスト時にモックに差し替えられる

type UserService struct {
	repo UserRepository // インターフェース型で保持
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(id int, name string) error {
	user := &User{ID: id, Name: name}
	return s.repo.Save(user)
}

func main() {
	// InMemoryUserRepository は UserRepository インターフェースを満たす
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)

	// ユーザー作成
	service.CreateUser(1, "田中太郎")

	// ユーザー取得
	user, err := service.GetUser(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Found: %+v\n", user)
}
```

> **💡 次のステップへ:** インターフェースによる依存性の逆転を学びました。次は標準ライブラリのインターフェースを活用するパターンを見てみましょう。

### 実践: 標準ライブラリのインターフェースを活用する

Goの標準ライブラリは小さなインターフェースを多数定義しています。`fmt.Stringer`や`sort.Interface`を実装することで、標準ライブラリの機能を自分の型で活用できます。

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// Employee は従業員を表す構造体
type Employee struct {
	Name       string
	Department string
	Salary     int
}

// fmt.Stringer インターフェースの実装
// → fmt.Println等で自動的にこのメソッドが呼ばれる
func (e Employee) String() string {
	return fmt.Sprintf("%s (%s部門, %d万円)", e.Name, e.Department, e.Salary)
}

// ---- sort.Interface の実装 ----
// Len, Less, Swap の3メソッドを実装すると sort.Sort が使える

type BySalary []Employee

func (s BySalary) Len() int           { return len(s) }
func (s BySalary) Less(i, j int) bool { return s[i].Salary < s[j].Salary }
func (s BySalary) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Summarizer は集計機能を表すインターフェース
type Summarizer interface {
	Summary() string
}

// Department は部門別の集計を行う
type Department struct {
	Name      string
	Employees []Employee
}

func (d Department) Summary() string {
	total := 0
	for _, e := range d.Employees {
		total += e.Salary
	}
	return fmt.Sprintf("%s部門: %d名, 合計給与%d万円", d.Name, len(d.Employees), total)
}

func printSummary(s Summarizer) {
	fmt.Println(s.Summary())
}

func main() {
	employees := []Employee{
		{Name: "田中", Department: "開発", Salary: 600},
		{Name: "鈴木", Department: "営業", Salary: 500},
		{Name: "佐藤", Department: "開発", Salary: 700},
	}

	// fmt.Stringer の活用
	for _, e := range employees {
		fmt.Println(e) // String() メソッドが自動的に呼ばれる
	}

	// sort.Interface の活用
	sort.Sort(BySalary(employees))
	fmt.Println("\n--- 給与順 ---")
	for _, e := range employees {
		fmt.Println(e)
	}

	// カスタムインターフェースの活用
	fmt.Println("\n--- 部門集計 ---")
	dept := Department{Name: "開発", Employees: employees}
	printSummary(dept)

	_ = strings.NewReader("dummy") // strings パッケージのインポートエラー回避
}
```

## 🎯 演習問題

通知システムをインターフェースで設計してください。

**要件:**

1. `Notifier`インターフェース: `Notify(message string) error`メソッドを持つ
2. `EmailNotifier`構造体: `To string`フィールドを持ち、`Notifier`を実装。`fmt.Println`でメール送信をシミュレート
3. `SlackNotifier`構造体: `Channel string`フィールドを持ち、`Notifier`を実装。`fmt.Println`でSlack送信をシミュレート
4. `SendAll(notifiers []Notifier, message string) []error`: 全Notifierにメッセージを送信し、エラーを収集

**ヒント:**

```go
type Notifier interface {
	Notify(message string) error
}

func SendAll(notifiers []Notifier, message string) []error {
	// 各notifierに対してNotifyを呼ぶ
}
```

**期待される動作:**

- `EmailNotifier{To: "user@example.com"}.Notify("Hello")` → "メール送信: user@example.com - Hello"
- `SlackNotifier{Channel: "#general"}.Notify("Hello")` → "Slack送信: #general - Hello"
- `SendAll`で複数のNotifierにまとめて送信

## ✅ 重要ポイント

- [ ] インターフェースは小さく定義する（1-2メソッドが理想）
- [ ] Goのインターフェースは暗黙的に実装される（implements不要）
- [ ] 関数の引数にインターフェースを使うことで疎結合を実現する
- [ ] 「Accept interfaces, return structs」が基本原則

**次のテーマ:** [05. Interfaceの応用](./05-interface-advanced.md)
