package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// 値レシーバー: 読み取り専用
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.1f x %.1f)", r.Width, r.Height)
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}

	fmt.Println(rect)
	fmt.Println("面積:", rect.Area())
	fmt.Println("周長:", rect.Perimeter())
}
