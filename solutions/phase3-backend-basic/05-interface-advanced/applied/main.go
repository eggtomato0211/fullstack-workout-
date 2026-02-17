package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }
type Rectangle struct{ Width, Height float64 }
type Triangle struct{ Base, Height float64 }

func (c Circle) Area() float64    { return 3.14 * c.Radius * c.Radius }
func (r Rectangle) Area() float64 { return r.Width * r.Height }
func (t Triangle) Area() float64  { return t.Base * t.Height / 2 }

// 型スイッチで図形の種類に応じた説明を出力
func describeShape(s Shape) {
	switch v := s.(type) {
	case Circle:
		fmt.Printf("円（半径: %.1f）面積: %.2f\n", v.Radius, v.Area())
	case Rectangle:
		fmt.Printf("長方形（%.1f x %.1f）面積: %.2f\n", v.Width, v.Height, v.Area())
	case Triangle:
		fmt.Printf("三角形（底辺: %.1f, 高さ: %.1f）面積: %.2f\n", v.Base, v.Height, v.Area())
	default:
		fmt.Printf("未知の図形: 面積 %.2f\n", s.Area())
	}
}

// any型の型スイッチ
func processValues(values []any) {
	for _, v := range values {
		switch val := v.(type) {
		case int:
			fmt.Printf("整数: %d\n", val)
		case string:
			fmt.Printf("文字列: %q\n", val)
		case bool:
			fmt.Printf("真偽値: %t\n", val)
		default:
			fmt.Printf("その他: %T = %v\n", val, val)
		}
	}
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 10, Height: 3},
		Triangle{Base: 8, Height: 6},
	}

	for _, s := range shapes {
		describeShape(s)
	}

	fmt.Println("\n--- any の型スイッチ ---")
	processValues([]any{42, "hello", true, 3.14})
}
