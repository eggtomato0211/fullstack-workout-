package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// fmt.Fprintf: io.Writer に書式付きで書き込む
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "名前: %s, 年齢: %d", "田中太郎", 30)
	fmt.Println("Fprintf:", buf.String())

	// io.MultiWriter: 複数のWriterに同時書き込み
	var log bytes.Buffer
	multi := io.MultiWriter(os.Stdout, &log)
	fmt.Fprintln(multi, "このメッセージは画面とログの両方に出力されます")
	fmt.Println("ログ内容:", log.String())

	// io.LimitReader: 読み取りサイズを制限
	longText := strings.NewReader("これは長いテキストです。最初の10バイトだけ読みます。")
	limited := io.LimitReader(longText, 10)
	data, _ := io.ReadAll(limited)
	fmt.Printf("LimitReader: %q\n", string(data))

	// io.TeeReader: 読み取り時に別のWriterにもコピー
	var teeLog bytes.Buffer
	original := strings.NewReader("TeeReaderのテスト")
	tee := io.TeeReader(original, &teeLog)
	result, _ := io.ReadAll(tee)
	fmt.Printf("TeeReader result: %q\n", string(result))
	fmt.Printf("TeeReader log:    %q\n", teeLog.String())
}
