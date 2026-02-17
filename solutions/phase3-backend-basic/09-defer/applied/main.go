package main

import (
	"fmt"
	"strings"
)

// FileHandle はファイル操作のシミュレーション
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

// 取得直後にdeferでクリーンアップ
func readFile(name string) (string, error) {
	f, err := OpenFile(name)
	if err != nil {
		return "", fmt.Errorf("open failed: %w", err)
	}
	defer f.Close()

	content := f.Read()
	return strings.ToUpper(content), nil
}

// 複数リソースのクリーンアップ
func copyFile(src, dst string) error {
	srcFile, err := OpenFile(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := OpenFile(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	content := srcFile.Read()
	fmt.Printf("コピー: %s → %s (%s)\n", src, dst, content)
	return nil
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
