# 09. defer文の活用 - リソース解放、実行順序、エラーハンドリングとの組み合わせ

## 🎯 このテーマで学ぶこと

- defer文の基本的な動作と実行順序（LIFO）
- リソース解放パターン（ファイル、DB接続など）
- deferとエラーハンドリングの組み合わせ
- deferの注意点（ループ内でのdefer、引数の評価タイミング）

**なぜ重要か:** ファイルのクローズ、DB接続の解放、ロックの解除など、「関数が終わったら必ず実行したい処理」はdeferで確実に行います。deferがないと、エラーリターンのたびに手動でクリーンアップコードを書く必要があり、リソースリークの原因になります。

## 📖 概念

`defer`は関数の終了時に実行される処理を予約する構文です。複数のdeferはLIFO（後入れ先出し）の順序で実行されます。deferはpanicが発生しても実行されるため、リソースのクリーンアップに最適です。

**背景と設計意図:** JavaやPythonでは`try/finally`でリソース解放を保証しますが、Goはdeferでよりシンプルに表現します。リソースの取得と解放を近い位置に書けるため、可読性も高くなります。

**よくある誤解:**

- ❌ 「deferは即座に実行される」→ 関数の終了時に実行される
- ❌ 「deferの引数は遅延評価される」→ 引数はdefer文の時点で評価される
- ❌ 「ループ内でdeferしてよい」→ ループ内のdeferはループ終了ではなく関数終了時に実行される

## 💡 コード例

### 基本: deferの動作と実行順序

deferの基本的な動作と、LIFO（後入れ先出し）の実行順序を学びます。

```go
package main

import "fmt"

func main() {
	fmt.Println("開始")

	// defer は関数の終了時に実行される
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("処理中...")
	fmt.Println("終了")

	// 出力順序:
	// 開始
	// 処理中...
	// 終了
	// defer 3  ← LIFO: 最後にdeferしたものが最初に実行
	// defer 2
	// defer 1
}
```

```go
package main

import "fmt"

// deferの引数は「defer文を実行した時点で」評価される
func deferArgEvaluation() {
	x := 10
	defer fmt.Println("deferされた x:", x) // この時点で x=10 が確定

	x = 20
	fmt.Println("現在の x:", x) // 20
	// 出力: 現在の x: 20 → deferされた x: 10
}

func main() {
	deferArgEvaluation()
}
```

> **💡 次のステップへ:** deferの基本動作を学びました。次はリソース管理での実用的な使い方を学びます。

### 応用: リソース管理パターン

ファイル操作やDBコネクションなど、リソースの取得直後にdeferでクリーンアップを予約するパターンです。

```go
package main

import (
	"fmt"
	"strings"
)

// ---- リソース管理のシミュレーション ----

type FileHandle struct {
	name   string
	closed bool
}

func OpenFile(name string) (*FileHandle, error) {
	fmt.Printf("ファイルを開く: %s\n", name)
	return &FileHandle{name: name}, nil
}

func (f *FileHandle) Read() string {
	return fmt.Sprintf("%sの内容", f.name)
}

func (f *FileHandle) Close() error {
	if f.closed {
		return fmt.Errorf("already closed: %s", f.name)
	}
	f.closed = true
	fmt.Printf("ファイルを閉じる: %s\n", f.name)
	return nil
}

// ---- 良いパターン: 取得直後にdefer ----

func readFile(name string) (string, error) {
	f, err := OpenFile(name)
	if err != nil {
		return "", fmt.Errorf("open failed: %w", err)
	}
	defer f.Close() // 取得直後にクリーンアップを予約

	// ここでエラーが起きても f.Close() は確実に実行される
	content := f.Read()
	if content == "" {
		return "", fmt.Errorf("empty file: %s", name)
	}

	return strings.ToUpper(content), nil
}

// ---- 複数リソースのクリーンアップ ----

func copyFile(src, dst string) error {
	srcFile, err := OpenFile(src)
	if err != nil {
		return err
	}
	defer srcFile.Close() // 1番目に予約

	dstFile, err := OpenFile(dst)
	if err != nil {
		return err // srcFile.Close() は実行される
	}
	defer dstFile.Close() // 2番目に予約

	content := srcFile.Read()
	fmt.Printf("コピー: %s → %s (%s)\n", src, dst, content)
	return nil

	// 関数終了時: dstFile.Close() → srcFile.Close() の順で実行
}

func main() {
	content, err := readFile("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("内容:", content)
	}

	fmt.Println("---")
	copyFile("source.txt", "dest.txt")
}
```

> **💡 次のステップへ:** リソース管理パターンを学びました。次はdeferの注意点と実務的なパターンを学びます。

### 実践: deferの注意点と実務パターン

deferの落とし穴と、エラーハンドリングと組み合わせた実務パターンを学びます。

```go
package main

import "fmt"

// ---- 注意点1: ループ内のdefer ----

type Resource struct {
	id int
}

func (r *Resource) Close() {
	fmt.Printf("Resource %d closed\n", r.id)
}

// 悪い例: ループ内のdeferは関数終了まで実行されない
func badLoopDefer() {
	fmt.Println("=== 悪い例: ループ内のdefer ===")
	for i := 0; i < 3; i++ {
		r := &Resource{id: i}
		defer r.Close() // 関数終了まで閉じられない！
		fmt.Printf("Resource %d を使用中\n", r.id)
	}
	fmt.Println("関数終了")
	// ここで全てのClose()が実行される（リソースが溜まる）
}

// 良い例: 無名関数でスコープを作る
func goodLoopDefer() {
	fmt.Println("\n=== 良い例: 無名関数でスコープを作る ===")
	for i := 0; i < 3; i++ {
		func() {
			r := &Resource{id: i}
			defer r.Close() // この無名関数の終了時に実行される
			fmt.Printf("Resource %d を使用中\n", r.id)
		}()
	}
	fmt.Println("関数終了")
}

// ---- 注意点2: deferでエラーを扱う ----

type DBConn struct {
	name string
}

func (c *DBConn) Close() error {
	fmt.Printf("DB接続 %s を閉じる\n", c.name)
	return nil
}

func (c *DBConn) Query(sql string) (string, error) {
	return "結果データ", nil
}

// deferでCloseのエラーを処理するパターン
func queryWithCleanup(connName, sql string) (result string, err error) {
	conn := &DBConn{name: connName}

	// 名前付き戻り値(err)を使って、deferでCloseのエラーを返す
	defer func() {
		closeErr := conn.Close()
		if err == nil {
			err = closeErr // メインの処理が成功した場合のみCloseのエラーを返す
		}
	}()

	result, err = conn.Query(sql)
	if err != nil {
		return "", fmt.Errorf("query failed: %w", err)
	}

	return result, nil
}

func main() {
	badLoopDefer()
	goodLoopDefer()

	fmt.Println("\n=== deferでエラーを扱う ===")
	result, err := queryWithCleanup("main-db", "SELECT * FROM users")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
```

## 🎯 演習問題

トランザクション管理をdeferで実装してください。

**要件:**

1. `Transaction`構造体: `ID string`, `committed bool`, `rolledback bool`フィールドを持つ
2. `BeginTx() *Transaction`: トランザクションを開始する
3. `Commit() error`: コミットする（2回目以降はエラー）
4. `Rollback() error`: ロールバックする（コミット済みならスキップ）
5. `ExecuteInTx(fn func(tx *Transaction) error) error`: トランザクション内で処理を実行し、エラー時はRollback、成功時はCommitをdeferで行う

**ヒント:**

```go
func ExecuteInTx(fn func(tx *Transaction) error) error {
	tx := BeginTx()
	defer func() {
		if /* エラーがあったら */ {
			tx.Rollback()
		}
	}()
	// fn を実行して、成功したら Commit
}
```

**期待される動作:**

- 正常時: Begin → 処理 → Commit
- エラー時: Begin → 処理(エラー) → Rollback

## ✅ 重要ポイント

- [ ] deferは関数の終了時にLIFO順で実行される
- [ ] リソース取得の直後にdeferでクリーンアップを予約する
- [ ] ループ内のdeferは無名関数でスコープを作って対応する
- [ ] 名前付き戻り値でdeferからエラーを返せる

**次のテーマ:** [10. panic/recoverの理解](./10-panic-recover.md)
