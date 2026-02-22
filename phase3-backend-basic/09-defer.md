# 09. defer文の活用 - リソース解放、実行順序、エラーハンドリングとの組み合わせ

## 🎯 このテーマで学ぶこと

- defer文の基本的な動作と実行順序（LIFO）
- リソース解放パターン（ファイル、DB接続など）
- deferの注意点（ループ内でのdefer、引数の評価タイミング）

## 📖 なぜdeferを理解する必要があるのか

ファイルのクローズ、DB接続の解放、ロックの解除など、「関数が終わったら必ず実行したい処理」はdeferで書きます。JavaやPythonの`try/finally`に相当しますが、Goのdeferは**リソース取得の直後に書ける**ため、取得と解放がセットで見えます。

### こう書かないとどうなるか

```go
// deferを使わない場合 → エラーリターンのたびにCloseが必要
func readFile(name string) (string, error) {
    f, err := os.Open(name)
    if err != nil {
        return "", err
    }

    data, err := io.ReadAll(f)
    if err != nil {
        f.Close() // ← ここでも閉じる必要がある
        return "", err
    }

    f.Close() // ← ここでも閉じる必要がある
    return string(data), nil
}

// deferを使う場合 → 1箇所書くだけで確実に実行される
func readFile(name string) (string, error) {
    f, err := os.Open(name)
    if err != nil {
        return "", err
    }
    defer f.Close() // ← どのリターンパスでも確実に実行される

    data, err := io.ReadAll(f)
    if err != nil {
        return "", err // f.Close()はdeferが呼ぶ
    }
    return string(data), nil
}
```

### deferの3つのルール

1. **関数の終了時に実行**される（即座ではない）
2. **LIFO順**で実行される（最後にdeferしたものが最初に実行）
3. **引数はdefer文の時点で評価**される（遅延評価ではない）

## 💡 コード例

### 基本: deferの動作とリソース管理

```go
package main

import (
	"fmt"
	"strings"
)

// --- deferの基本動作 ---

func deferBasics() {
	fmt.Println("開始")

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("終了")
	// 出力: 開始 → 終了 → defer 3 → defer 2 → defer 1（LIFO）
}

// --- 引数はdefer文の時点で評価される ---

func deferArgEvaluation() {
	x := 10
	defer fmt.Println("deferされた x:", x) // この時点で x=10 が確定

	x = 20
	fmt.Println("現在の x:", x) // 20
	// 出力: 現在の x: 20 → deferされた x: 10
}

// --- リソース管理: 取得直後にdeferで解放を予約 ---

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

func readFile(name string) (string, error) {
	f, err := OpenFile(name)
	if err != nil {
		return "", fmt.Errorf("open failed: %w", err)
	}
	defer f.Close() // 取得直後にクリーンアップを予約

	// ここでエラーが起きても f.Close() は確実に実行される
	content := f.Read()
	return strings.ToUpper(content), nil
}

func main() {
	deferBasics()
	fmt.Println("---")
	deferArgEvaluation()
	fmt.Println("---")

	content, err := readFile("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("内容:", content)
	}
}
```

### 実践: deferの落とし穴と実務パターン

ループ内のdeferや、deferでエラーを返すパターンなど、実務で遭遇する注意点を学びます。

```go
package main

import "fmt"

// --- 落とし穴: ループ内のdefer ---

type Resource struct{ id int }

func (r *Resource) Close() {
	fmt.Printf("Resource %d closed\n", r.id)
}

// 悪い例: ループ内のdeferは関数終了まで実行されない
func badLoopDefer() {
	fmt.Println("=== 悪い例 ===")
	for i := 0; i < 3; i++ {
		r := &Resource{id: i}
		defer r.Close() // 関数終了まで閉じられない → リソースが溜まる！
		fmt.Printf("Resource %d を使用中\n", r.id)
	}
	fmt.Println("関数終了")
}

// 良い例: 無名関数でスコープを作る
func goodLoopDefer() {
	fmt.Println("\n=== 良い例 ===")
	for i := 0; i < 3; i++ {
		func() {
			r := &Resource{id: i}
			defer r.Close() // この無名関数の終了時に実行される
			fmt.Printf("Resource %d を使用中\n", r.id)
		}()
	}
	fmt.Println("関数終了")
}

// --- 実務パターン: deferでCloseのエラーを処理 ---

type DBConn struct{ name string }

func (c *DBConn) Close() error {
	fmt.Printf("DB接続 %s を閉じる\n", c.name)
	return nil
}

func (c *DBConn) Query(sql string) (string, error) {
	return "結果データ", nil
}

// 名前付き戻り値を使って、deferからエラーを返す
func queryWithCleanup(connName, sql string) (result string, err error) {
	conn := &DBConn{name: connName}

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

**期待される動作:**

- 正常時: Begin → 処理 → Commit
- エラー時: Begin → 処理(エラー) → Rollback

## ✅ 重要ポイント

- [ ] deferは関数の終了時にLIFO順で実行される
- [ ] リソース取得の直後にdeferでクリーンアップを予約する
- [ ] ループ内のdeferは無名関数でスコープを作って対応する
- [ ] 名前付き戻り値でdeferからエラーを返せる

**次のテーマ:** [10. panic/recoverの理解](./10-panic-recover.md)
