package main

import "fmt"

type Counter struct {
	value int
	name  string
}

func NewCounter(name string) *Counter {
	return &Counter{name: name, value: 0}
}

// ポインタレシーバー: 値を変更する
func (c *Counter) Increment() { c.value++ }
func (c *Counter) Add(n int)  { c.value += n }
func (c *Counter) Reset()     { c.value = 0 }

// 値レシーバー: 読み取り専用
func (c Counter) Value() int { return c.value }
func (c Counter) String() string {
	return fmt.Sprintf("%s: %d", c.name, c.value)
}

func main() {
	counter := NewCounter("アクセス数")
	fmt.Println(counter)

	counter.Increment()
	counter.Increment()
	counter.Add(10)
	fmt.Println(counter)
	fmt.Println("現在値:", counter.Value())

	counter.Reset()
	fmt.Println(counter)
}
