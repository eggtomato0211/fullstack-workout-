# 12. io.Reader/io.Writerの理解 - Goのインターフェースパターンの代表例

## 🎯 このテーマで学ぶこと

- `io.Reader`と`io.Writer`インターフェースの仕組み
- `io.Copy`、`io.ReadAll`などのユーティリティ関数
- 自作のReader/Writerの実装

## 📖 なぜio.Reader/io.Writerを理解する必要があるのか

`io.Reader`と`io.Writer`はGoで**最も重要なインターフェース**です。ファイル操作、HTTP通信、暗号化、圧縮など、あらゆるI/O操作がこの2つに統一されています。UNIXの「全てはファイル」という哲学を、Goはインターフェースで実現しました。

### こう書かないとどうなるか

```go
// io.Readerを使わない場合 → データソースごとに別の関数が必要
func processFromFile(path string) { ... }
func processFromNetwork(url string) { ... }
func processFromString(s string) { ... }

// io.Readerを使う場合 → 1つの関数で全てに対応
func process(r io.Reader) {
    data, _ := io.ReadAll(r)
    // ファイルでもネットワークでも文字列でも同じ処理
}

process(file)                          // *os.File は io.Reader
process(resp.Body)                     // http.Response.Body は io.Reader
process(strings.NewReader("hello"))    // strings.Reader は io.Reader
```

### よくある誤解

- `io.EOF`はエラーではなく、データの終端を示す**正常なシグナル**
- `Read`は1回で全データを読めるとは限らない。`io.ReadAll`で一括読み取りが安全

## 💡 コード例

### 基本: io.Reader/io.Writerの基礎

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// --- io.Reader: データの読み取り ---

	// strings.NewReader: 文字列を io.Reader として扱う
	reader := strings.NewReader("Hello, Go!")

	// io.ReadAll: Reader から全データを一括読み取り
	data, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Println("ReadAll:", string(data))

	// --- io.Writer: データの書き込み ---

	// bytes.Buffer は io.Writer を実装している
	var buf bytes.Buffer
	buf.Write([]byte("Hello, "))
	buf.WriteString("World!")
	fmt.Println("Buffer:", buf.String())

	// --- io.Copy: Reader → Writer にコピー ---
	// ファイルコピー、HTTPレスポンスの転送などで頻出
	src := strings.NewReader("コピーされるデータ")
	var dst bytes.Buffer

	n, err := io.Copy(&dst, src)
	if err != nil {
		fmt.Println("Copy error:", err)
		return
	}
	fmt.Printf("Copy: %d bytes → %q\n", n, dst.String())

	// --- io.MultiWriter: 複数のWriterに同時に書き込む ---
	// ログを画面とファイルの両方に出力する場面で便利
	var log1, log2 bytes.Buffer
	multi := io.MultiWriter(&log1, &log2)

	fmt.Fprintln(multi, "このメッセージは両方に書き込まれる")
	fmt.Println("log1:", log1.String())
	fmt.Println("log2:", log2.String())
}
```

### 実践: カスタムReader/Writerの実装

`io.Reader`/`io.Writer`を実装して、データの変換パイプラインを作るパターンです。

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"unicode"
)

// UpperReader は読み取ったデータを大文字に変換するReader
// io.Readerを満たす → io.Copy, io.ReadAll などで使える
type UpperReader struct {
	reader io.Reader
}

func NewUpperReader(r io.Reader) *UpperReader {
	return &UpperReader{reader: r}
}

func (u *UpperReader) Read(p []byte) (int, error) {
	n, err := u.reader.Read(p)
	for i := 0; i < n; i++ {
		p[i] = byte(unicode.ToUpper(rune(p[i])))
	}
	return n, err
}

// CountWriter は書き込まれたバイト数をカウントするWriter
type CountWriter struct {
	writer    io.Writer
	ByteCount int
}

func NewCountWriter(w io.Writer) *CountWriter {
	return &CountWriter{writer: w}
}

func (c *CountWriter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	c.ByteCount += n
	return n, err
}

// PrefixWriter は各行の先頭にプレフィックスを追加するWriter
type PrefixWriter struct {
	writer      io.Writer
	prefix      string
	atLineStart bool
}

func NewPrefixWriter(w io.Writer, prefix string) *PrefixWriter {
	return &PrefixWriter{writer: w, prefix: prefix, atLineStart: true}
}

func (pw *PrefixWriter) Write(p []byte) (int, error) {
	var written int
	for _, b := range p {
		if pw.atLineStart {
			pw.writer.Write([]byte(pw.prefix))
			pw.atLineStart = false
		}
		n, err := pw.writer.Write([]byte{b})
		written += n
		if err != nil {
			return written, err
		}
		if b == '\n' {
			pw.atLineStart = true
		}
	}
	return written, nil
}

func main() {
	// --- UpperReader ---
	src := strings.NewReader("hello, world!")
	upper := NewUpperReader(src)
	data, _ := io.ReadAll(upper)
	fmt.Println("UpperReader:", string(data)) // HELLO, WORLD!

	// --- CountWriter ---
	var buf bytes.Buffer
	counter := NewCountWriter(&buf)
	fmt.Fprint(counter, "Hello!")
	fmt.Fprint(counter, " How are you?")
	fmt.Printf("CountWriter: %q (%d bytes)\n", buf.String(), counter.ByteCount)

	// --- PrefixWriter ---
	var logBuf bytes.Buffer
	logger := NewPrefixWriter(&logBuf, "[LOG] ")
	fmt.Fprintln(logger, "サーバー起動")
	fmt.Fprintln(logger, "リクエスト受信")
	fmt.Print("PrefixWriter:\n", logBuf.String())

	// --- Reader/Writerの組み合わせ ---
	// カスタムReader/Writerはio.Copyで接続できる
	fmt.Println("\n--- 組み合わせ ---")
	src2 := strings.NewReader("go is awesome!")
	upperReader := NewUpperReader(src2)

	var result bytes.Buffer
	countWriter := NewCountWriter(&result)

	io.Copy(countWriter, upperReader)
	fmt.Printf("結果: %q (%d bytes)\n", result.String(), countWriter.ByteCount)
}
```

## 🎯 演習問題

ログシステムをio.Writerベースで設計してください。

**要件:**

1. `LogLevel`型: `INFO`, `WARN`, `ERROR`の定数を定義
2. `Logger`構造体: `io.Writer`と`LogLevel`を持つ
3. `NewLogger(w io.Writer, level LogLevel) *Logger`: コンストラクタ
4. `Info(msg string)`, `Warn(msg string)`, `Error(msg string)`: 設定レベル以上のログのみ出力
5. 出力先を`bytes.Buffer`に変えてテストできることを確認

**期待される動作:**

- `Logger{level: WARN}`の場合、`Info()`は出力されず、`Warn()`と`Error()`のみ出力
- 出力先を変えるだけで、標準出力・バッファ・ファイルに切り替え可能

## ✅ 重要ポイント

- [ ] `io.Reader`と`io.Writer`はGoで最も重要なインターフェース
- [ ] `io.ReadAll`、`io.Copy`で簡単にデータを読み書きできる
- [ ] カスタムReader/Writerでデータの変換パイプラインを作れる
- [ ] `io.MultiWriter`で複数の出力先にデータを流せる

**カテゴリAの完了です！おつかれさまでした！**
カテゴリBに進む場合: [13. Goroutineの基本](./13-goroutine-basics.md)
