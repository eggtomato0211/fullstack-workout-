package main

import (
	"errors"
	"fmt"
	"time"
)

// BaseModel は共通フィールドを持つ基底構造体
type BaseModel struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Touch は更新日時を現在時刻に更新する
func (b *BaseModel) Touch() {
	b.UpdatedAt = time.Now()
}

// Product はBaseModelを埋め込んだ商品構造体
type Product struct {
	BaseModel
	Name  string
	Price int
	Stock int
}

// NewProduct はProductのコンストラクタ
func NewProduct(name string, price, stock int) (*Product, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}
	if price <= 0 {
		return nil, errors.New("price must be positive")
	}
	now := time.Now()
	return &Product{
		BaseModel: BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name:  name,
		Price: price,
		Stock: stock,
	}, nil
}

// IsInStock は在庫があるかを返す
func (p Product) IsInStock() bool {
	return p.Stock > 0
}

// String は商品情報の文字列表現を返す
func (p Product) String() string {
	status := "在庫あり"
	if !p.IsInStock() {
		status = "在庫なし"
	}
	return fmt.Sprintf("%s - ¥%d (%s, 残り%d個)", p.Name, p.Price, status, p.Stock)
}

func main() {
	// 正常な商品作成
	product, err := NewProduct("Goの本", 3000, 10)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(product)
	fmt.Println("在庫あり:", product.IsInStock())

	// BaseModelのフィールドに直接アクセス（埋め込みによる昇格）
	fmt.Println("作成日時:", product.CreatedAt.Format("2006-01-02 15:04:05"))

	// Touchメソッドも直接呼べる
	product.Touch()
	fmt.Println("更新日時:", product.UpdatedAt.Format("2006-01-02 15:04:05"))

	// バリデーションエラー
	_, err = NewProduct("", 3000, 10)
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = NewProduct("テスト", 0, 10)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
