package main

import "fmt"

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()

	return a / b, nil
}

func main() {
	// 正常ケース
	result, err := safeDivide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 3 =", result)
	}

	// panicが発生するケース
	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// プログラムは続行できる
	fmt.Println("プログラムは続行中")
}
