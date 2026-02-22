# 04. Interfaceの基本 - interface定義、暗黙的な実装

## 🎯 このテーマで学ぶこと

- インターフェースの定義と暗黙的な実装
- インターフェースを引数に取る関数の設計
- 「Accept interfaces, return structs」の原則

## 📖 なぜインターフェースを理解する必要があるのか

インターフェースはGoの最も強力な設計ツールです。「この関数はDBに依存している → テストでDBを使いたくない → インターフェースでモックに差し替えよう」という場面で必須になります。

### こう書かないとどうなるか

```go
// インターフェースを使わない場合
type UserService struct {
    db *PostgresDB  // 具象型に直接依存
}

// テストしたいのに、本物のDBが必要になってしまう
// → テストが遅い、セットアップが面倒、CIで動かない
```

インターフェースを使えば、本番ではPostgres、テストではメモリ上のモックに差し替えられます。

### Goのインターフェースが特殊な理由

JavaやC#では`implements`キーワードで明示的にインターフェースを実装しますが、Goは**暗黙的に実装**されます。必要なメソッドを全て持っていれば、宣言なしで自動的にそのインターフェースを満たします。これを「ダックタイピング」と呼びます。

### 小さく作るのがGo流

Goのインターフェースは1-2メソッドが理想です。標準ライブラリの`io.Reader`（`Read`メソッドだけ）、`io.Writer`（`Write`メソッドだけ）がお手本です。大きなインターフェースは実装が大変で、モックも面倒になります。

## 💡 コード例

### 基本: インターフェースの定義と暗黙的な実装

インターフェースを定義し、複数の型で実装する基本パターンです。

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

// Rectangle は Shape を「暗黙的に」実装
// implementsキーワードは不要 — メソッドが合えば自動的に満たす
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// Circle も Shape を「暗黙的に」実装
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

// 引数にインターフェースを取る → どの図形でも渡せる
func printShapeInfo(s Shape) {
	fmt.Printf("面積: %.2f, 周長: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	// どちらもShapeを満たしているので渡せる
	printShapeInfo(rect)   // 面積: 50.00, 周長: 30.00
	printShapeInfo(circle) // 面積: 153.94, 周長: 43.98
}
```

### 実践: リポジトリパターンで依存性を逆転する

インターフェースを使って具象型への依存を減らし、テスト時にモックに差し替えられる設計です。これが実務でインターフェースを使う最も典型的な場面です。

```go
package main

import "fmt"

// UserRepository はデータの永続化を抽象化するインターフェース
// メソッドが少ない（2つ）→ Goらしい小さなインターフェース
type UserRepository interface {
	FindByID(id int) (*User, error)
	Save(user *User) error
}

type User struct {
	ID   int
	Name string
}

// --- 本番用の実装 ---

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

// --- UserService はインターフェースに依存 ---
// 具象型ではなくインターフェースに依存することで、
// テスト時にモックに差し替えられる
type UserService struct {
	repo UserRepository // インターフェース型で保持
}

// 引数もインターフェース型（Accept interfaces）
// 戻り値は具象型（return structs）
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(id int, name string) error {
	return s.repo.Save(&User{ID: id, Name: name})
}

func main() {
	repo := NewInMemoryUserRepository()
	service := NewUserService(repo)

	service.CreateUser(1, "田中太郎")

	user, err := service.GetUser(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Found: %+v\n", user) // Found: &{ID:1 Name:田中太郎}

	// テスト時はモックを渡せる：
	// mockRepo := &MockUserRepository{...}
	// service := NewUserService(mockRepo)
}
```

## 🎯 演習問題

通知システムをインターフェースで設計してください。

**要件:**

1. `Notifier`インターフェース: `Notify(message string) error`メソッド
2. `EmailNotifier`構造体: `To string`フィールド。`fmt.Println`でメール送信をシミュレート
3. `SlackNotifier`構造体: `Channel string`フィールド。`fmt.Println`でSlack送信をシミュレート
4. `SendAll(notifiers []Notifier, message string) []error`: 全Notifierに送信

**ヒント:**

```go
type Notifier interface {
	Notify(message string) error
}

func SendAll(notifiers []Notifier, message string) []error {
	// 各notifierに対してNotifyを呼ぶ
}
```

## ✅ 重要ポイント

- [ ] インターフェースは小さく定義する（1-2メソッドが理想）
- [ ] Goのインターフェースは暗黙的に実装される（implements不要）
- [ ] 引数にインターフェースを使い、戻り値は具象型を返す
- [ ] テスト時のモック差し替えが最も典型的なユースケース

**次のテーマ:** [05. Interfaceの応用](./05-interface-advanced.md)
