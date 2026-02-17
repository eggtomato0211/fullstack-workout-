package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// LogLevel はログレベルを表す型
type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERROR
)

func (l LogLevel) String() string {
	switch l {
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Logger はio.Writerベースのロガー
type Logger struct {
	out   io.Writer
	level LogLevel
}

func NewLogger(w io.Writer, level LogLevel) *Logger {
	return &Logger{out: w, level: level}
}

func (l *Logger) log(level LogLevel, msg string) {
	if level >= l.level {
		fmt.Fprintf(l.out, "[%s] %s\n", level, msg)
	}
}

func (l *Logger) Info(msg string)  { l.log(INFO, msg) }
func (l *Logger) Warn(msg string)  { l.log(WARN, msg) }
func (l *Logger) Error(msg string) { l.log(ERROR, msg) }

func main() {
	// 標準出力へのログ
	fmt.Println("=== 標準出力（レベル: INFO） ===")
	logger := NewLogger(os.Stdout, INFO)
	logger.Info("サーバー起動")
	logger.Warn("メモリ使用率が高い")
	logger.Error("DB接続失敗")

	// WARNレベル以上のみ出力
	fmt.Println("\n=== 標準出力（レベル: WARN） ===")
	warnLogger := NewLogger(os.Stdout, WARN)
	warnLogger.Info("これは出力されない")
	warnLogger.Warn("これは出力される")
	warnLogger.Error("これも出力される")

	// バッファへのログ（テスト用）
	fmt.Println("\n=== バッファ出力（テスト） ===")
	var buf bytes.Buffer
	testLogger := NewLogger(&buf, INFO)
	testLogger.Info("テストメッセージ1")
	testLogger.Warn("テストメッセージ2")
	testLogger.Error("テストメッセージ3")
	fmt.Print("バッファ内容:\n", buf.String())
}
