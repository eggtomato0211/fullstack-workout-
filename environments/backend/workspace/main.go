package main

import (
	"fmt"
	"sort"
	"strings"
)

// Employee は従業員を表す構造体
type Employee struct {
	Name       string
	Department string
	Salary     int
}

// fmt.Stringer インターフェースの実装
// → fmt.Println等で自動的にこのメソッドが呼ばれる
func (e Employee) String() string {
	return fmt.Sprintf("%s (%s部門, %d万円)", e.Name, e.Department, e.Salary)
}

// ---- sort.Interface の実装 ----
// Len, Less, Swap の3メソッドを実装すると sort.Sort が使える

type BySalary []Employee

func (s BySalary) Len() int           { return len(s) }
func (s BySalary) Less(i, j int) bool { return s[i].Salary < s[j].Salary }
func (s BySalary) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Summarizer は集計機能を表すインターフェース
type Summarizer interface {
	Summary() string
}

// Department は部門別の集計を行う
type Department struct {
	Name      string
	Employees []Employee
}

func (d Department) Summary() string {
	total := 0
	for _, e := range d.Employees {
		total += e.Salary
	}
	return fmt.Sprintf("%s部門: %d名, 合計給与%d万円", d.Name, len(d.Employees), total)
}

func printSummary(s Summarizer) {
	fmt.Println(s.Summary())
}

func main() {
	employees := []Employee{
		{Name: "田中", Department: "開発", Salary: 600},
		{Name: "鈴木", Department: "営業", Salary: 500},
		{Name: "佐藤", Department: "開発", Salary: 700},
	}

	// fmt.Stringer の活用
	for _, e := range employees {
		fmt.Println(e) // String() メソッドが自動的に呼ばれる
	}

	// sort.Interface の活用
	sort.Sort(BySalary(employees))
	fmt.Println("\n--- 給与順 ---")
	for _, e := range employees {
		fmt.Println(e)
	}

	// カスタムインターフェースの活用
	fmt.Println("\n--- 部門集計 ---")
	dept := Department{Name: "開発", Employees: employees}
	printSummary(dept)

	_ = strings.NewReader("dummy") // strings パッケージのインポートエラー回避
}