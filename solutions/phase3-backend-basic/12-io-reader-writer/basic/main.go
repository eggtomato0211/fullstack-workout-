package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// io.Reader: 文字列をReaderとして扱う
	reader := strings.NewReader("Hello, Go!")
	data, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Println("ReadAll:", string(data))

	// io.Writer: bytes.Buffer に書き込む
	var buf bytes.Buffer
	buf.Write([]byte("Hello, "))
	buf.WriteString("World!")
	fmt.Println("Buffer:", buf.String())

	// io.Copy: Reader → Writer にコピー
	src := strings.NewReader("コピーされるデータ")
	var dst bytes.Buffer
	n, err := io.Copy(&dst, src)
	if err != nil {
		fmt.Println("Copy error:", err)
		return
	}
	fmt.Printf("Copy: %d bytes → %q\n", n, dst.String())
}
