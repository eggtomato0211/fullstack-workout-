package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{ Name string }
type Cat struct{ Name string }

func (d Dog) Speak() string { return "ワン！" }
func (d Dog) Fetch() string { return d.Name + "がボールを取ってきた！" }
func (c Cat) Speak() string { return "ニャー！" }
func (c Cat) Purr() string  { return c.Name + "がゴロゴロ言っている" }

func describeAnimal(a Animal) {
	fmt.Println("鳴き声:", a.Speak())

	// 型アサーション（2値返却パターン）
	if dog, ok := a.(Dog); ok {
		fmt.Println(dog.Fetch())
	}
	if cat, ok := a.(Cat); ok {
		fmt.Println(cat.Purr())
	}
}

func main() {
	animals := []Animal{
		Dog{Name: "ポチ"},
		Cat{Name: "タマ"},
	}

	for _, a := range animals {
		describeAnimal(a)
		fmt.Println("---")
	}
}
