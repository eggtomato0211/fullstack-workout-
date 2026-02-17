# 12. io.Reader/io.Writerの理解 - Goのインターフェースパターンの代表例、ストリーム処理

## 🎯 このテーマで学ぶこと

- `io.Reader`と`io.Writer`インターフェースの仕組み
- 標準ライブラリでの`io.Reader`/`io.Writer`の活用
- `io.Copy`、`io.ReadAll`などのユーティリティ関数
- 自作のReader/Writerの実装

**なぜ重要か:** `io.Reader`と`io.Writer`はGoで最も重要なインターフェースです。ファイル操作、HTTP通信、暗号化、圧縮など、あらゆるI/O操作がこの2つのインターフェースで統一されています。このパターンを理解することで、Goの標準ライブラリを効果的に活用できるようになります。

## 📖 概念

`io.Reader`は`Read(p []byte) (n int, err error)`、`io.Writer`は`Write(p []byte) (n int, err error)`というたった1つのメソッドを持つインターフェースです。この小さなインターフェースにより、ファイル、ネットワーク、メモリバッファなど、異なるデータソースを同一の方法で扱えます。

**背景と設計意図:** UNIXの「全てはファイル」という哲学をGoはインターフェースで実現しました。`io.Reader`を満たすものは全て同じ方法で読み取れるため、データソースの差し替えが容易です。

**よくある誤解:**

- ❌ 「io.Readerは1回のReadで全データが読める」→ バッファサイズ分ずつ読み取る。`io.ReadAll`で一括読み取り
- ❌ 「io.EOF はエラー」→ データの終端を示す正常なシグナル
- ❌ 「strings.NewReader は不要」→ 文字列をio.Readerとして扱いたい場面で必須

## 💡 コード例

### 基本: io.Reader/io.Writerの基礎

`io.Reader`と`io.Writer`の基本的な使い方を学びます。

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// ---- io.Reader: データの読み取り ----

	// strings.NewReader: 文字列を io.Reader として扱う
	reader := strings.NewReader("Hello, Go!")

	// io.ReadAll: Reader から全データを読み取る
	data, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Println("ReadAll:", string(data)) // Hello, Go!

	// ---- io.Writer: データの書き込み ----

	// bytes.Buffer は io.Writer を実装している
	var buf bytes.Buffer

	// Write メソッドで書き込み
	buf.Write([]byte("Hello, "))
	buf.WriteString("World!")
	fmt.Println("Buffer:", buf.String()) // Hello, World!

	// ---- io.Copy: Reader → Writer にコピー ----

	src := strings.NewReader("コピーされるデータ")
	var dst bytes.Buffer

	n, err := io.Copy(&dst, src)
	if err != nil {
		fmt.Println("Copy error:", err)
		return
	}
	fmt.Printf("Copy: %d bytes → %q\n", n, dst.String())
}
```

> **💡 次のステップへ:** 基本的なReader/Writerの使い方を学びました。次は標準ライブラリでの実用的な活用パターンを学びます。

### 応用: 標準ライブラリでの活用

`fmt.Fprintf`、`io.MultiWriter`など、Reader/Writerを活用する標準ライブラリの関数を学びます。

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// ---- fmt.Fprintf: io.Writer に書式付きで書き込む ----
	var buf bytes.Buffer

	// os.Stdout も bytes.Buffer も io.Writer を満たす
	fmt.Fprintf(&buf, "名前: %s, 年齢: %d", "田中太郎", 30)
	fmt.Println("Fprintf:", buf.String())

	// ---- io.MultiWriter: 複数のWriterに同時に書き込む ----
	var log bytes.Buffer

	// stdout と log の両方に書き込む
	multi := io.MultiWriter(os.Stdout, &log)
	fmt.Fprintln(multi, "このメッセージは画面とログの両方に出力されます")
	fmt.Println("ログ内容:", log.String())

	// ---- io.LimitReader: 読み取りサイズを制限 ----
	longText := strings.NewReader("これは長いテキストです。最初の10バイトだけ読みます。")
	limited := io.LimitReader(longText, 10)

	data, _ := io.ReadAll(limited)
	fmt.Printf("LimitReader: %q\n", string(data))

	// ---- io.TeeReader: 読み取り時に別のWriterにもコピー ----
	var teeLog bytes.Buffer
	original := strings.NewReader("TeeReaderのテスト")
	tee := io.TeeReader(original, &teeLog)

	// tee から読むと、同じデータが teeLog にも書き込まれる
	result, _ := io.ReadAll(tee)
	fmt.Printf("TeeReader result: %q\n", string(result))
	fmt.Printf("TeeReader log:    %q\n", teeLog.String())
}
```

> **💡 次のステップへ:** 標準ライブラリの活用を学びました。次は自分でReader/Writerを実装するパターンを学びます。

### 実践: カスタムReader/Writerの実装

`io.Reader`/`io.Writer`を実装して、データの変換や加工を行うパターンを学びます。

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
type UpperReader struct {
	reader io.Reader
}

func NewUpperReader(r io.Reader) *UpperReader {
	return &UpperReader{reader: r}
}

// Read は io.Reader インターフェースを実装
func (u *UpperReader) Read(p []byte) (int, error) {
	n, err := u.reader.Read(p)
	// 読み取ったデータを大文字に変換
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

// Write は io.Writer インターフェースを実装
func (c *CountWriter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	c.ByteCount += n
	return n, err
}

// PrefixWriter は各行の先頭にプレフィックスを追加するWriter
type PrefixWriter struct {
	writer    io.Writer
	prefix    string
	atLineStart bool
}

func NewPrefixWriter(w io.Writer, prefix string) *PrefixWriter {
	return &PrefixWriter{writer: w, prefix: prefix, atLineStart: true}
}

func (pw *PrefixWriter) Write(p []byte) (int, error) {
	var written int
	for _, b := range p {
		if pw.atLineStart {
			n, err := pw.writer.Write([]byte(pw.prefix))
			if err != nil {
				return written, err
			}
			_ = n
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
	// ---- UpperReader の使用 ----
	src := strings.NewReader("hello, world!")
	upper := NewUpperReader(src)

	data, _ := io.ReadAll(upper)
	fmt.Println("UpperReader:", string(data)) // HELLO, WORLD!

	// ---- CountWriter の使用 ----
	var buf bytes.Buffer
	counter := NewCountWriter(&buf)

	fmt.Fprint(counter, "Hello!")
	fmt.Fprint(counter, " How are you?")
	fmt.Printf("CountWriter: %q (%d bytes)\n", buf.String(), counter.ByteCount)

	// ---- PrefixWriter の使用 ----
	var logBuf bytes.Buffer
	logger := NewPrefixWriter(&logBuf, "[LOG] ")

	fmt.Fprintln(logger, "サーバー起動")
	fmt.Fprintln(logger, "リクエスト受信")
	fmt.Fprintln(logger, "レスポンス送信")
	fmt.Print("PrefixWriter:\n", logBuf.String())

	// ---- Reader/Writerの組み合わせ ----
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

**ヒント:**

```go
type Logger struct {
	out   io.Writer
	level LogLevel
}

func (l *Logger) log(level LogLevel, msg string) {
	if level >= l.level {
		fmt.Fprintf(l.out, "[%s] %s\n", level, msg)
	}
}
```

**期待される動作:**

- `Logger{level: WARN}`の場合、`Info()`は出力されず、`Warn()`と`Error()`のみ出力
- 出力先を変えるだけで、標準出力・バッファ・ファイルに切り替え可能

## ✅ 重要ポイント

- [ ] `io.Reader`と`io.Writer`はGoで最も重要なインターフェース
- [ ] `io.ReadAll`、`io.Copy`で簡単にデータを読み書きできる
- [ ] `io.MultiWriter`、`io.TeeReader`で複数の出力先にデータを流せる
- [ ] カスタムReader/Writerでデータの変換パイプラインを作れる

**カテゴリAの完了です！おつかれさまでした！**
カテゴリBに進む場合: [13. Goroutineの基本](./13-goroutine-basics.md)
