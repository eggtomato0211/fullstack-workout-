package main

import "fmt"

func main() {
	fmt.Println("開始")

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("処理中...")
	fmt.Println("終了")
	// 出力: 開始 → 処理中... → 終了 → defer 3 → defer 2 → defer 1

	fmt.Println("\n--- 引数の評価タイミング ---")
	deferArgEvaluation()
}

func deferArgEvaluation() {
	x := 10
	defer fmt.Println("deferされた x:", x) // この時点で x=10 が確定

	x = 20
	fmt.Println("現在の x:", x)
}
