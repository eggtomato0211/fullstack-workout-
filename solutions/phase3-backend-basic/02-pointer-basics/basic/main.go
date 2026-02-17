package main

import "fmt"

func main() {
	x := 42

	// &x でメモリアドレスを取得
	p := &x
	fmt.Println("xの値:", x)
	fmt.Println("xのアドレス:", p)
	fmt.Println("pが指す値:", *p)

	// ポインタ経由で元の値を変更
	*p = 100
	fmt.Println("変更後のx:", x)

	// ポインタのゼロ値はnil
	var q *int
	fmt.Println("qの値:", q)

	// nilチェック
	if q != nil {
		fmt.Println(*q)
	} else {
		fmt.Println("qはnilです")
	}
}
