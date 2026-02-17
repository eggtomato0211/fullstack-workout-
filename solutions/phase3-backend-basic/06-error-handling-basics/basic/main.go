package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// 標準ライブラリのエラーを受け取る
	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("変換エラー:", err)
		return
	}
	fmt.Println("変換成功:", num)

	// 変換失敗のケース
	_, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("変換エラー:", err)
	}

	// errors.New でエラーを作成
	err = errors.New("something went wrong")
	fmt.Println(err)

	// fmt.Errorf で書式付きエラーを作成
	name := "admin"
	err = fmt.Errorf("user not found: %s", name)
	fmt.Println(err)
}
