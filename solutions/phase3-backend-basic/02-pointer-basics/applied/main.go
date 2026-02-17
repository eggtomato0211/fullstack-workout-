package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 値渡し: コピーが作られるため元の値は変わらない
func incrementAgeByValue(u User) {
	u.Age++
	fmt.Println("関数内(値渡し):", u.Age)
}

// ポインタ渡し: 元の値を変更できる
func incrementAgeByPointer(u *User) {
	u.Age++
	fmt.Println("関数内(ポインタ渡し):", u.Age)
}

func main() {
	user := User{Name: "田中太郎", Age: 30}

	// 値渡し
	incrementAgeByValue(user)
	fmt.Println("値渡し後:", user.Age) // 30（変わらない）

	// ポインタ渡し
	incrementAgeByPointer(&user)
	fmt.Println("ポインタ渡し後:", user.Age) // 31（変わった）
}
