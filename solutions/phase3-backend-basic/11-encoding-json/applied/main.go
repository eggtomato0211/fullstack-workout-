package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description,omitempty"` // 空文字列なら省略
	Discount    int    `json:"discount,omitempty"`    // 0なら省略
	InternalKey string `json:"-"`                     // JSONに含めない
	Stock       *int   `json:"stock,omitempty"`       // nilなら省略
}

func intPtr(i int) *int { return &i }

func main() {
	// omitempty の動作確認
	products := []Product{
		{
			ID: 1, Name: "Go入門書", Price: 3000,
			Description: "Goの基礎を学ぶ本", Discount: 500,
			InternalKey: "INTERNAL_001", Stock: intPtr(10),
		},
		{
			ID: 2, Name: "キーボード", Price: 15000,
			Description: "", Discount: 0,
			InternalKey: "INTERNAL_002", Stock: nil,
		},
	}

	for _, p := range products {
		jsonBytes, _ := json.MarshalIndent(p, "", "  ")
		fmt.Println(string(jsonBytes))
		fmt.Println("---")
	}

	// 未知のフィールドは無視される
	jsonStr := `{"id":3,"name":"マウス","price":5000,"color":"black"}`
	var product Product
	json.Unmarshal([]byte(jsonStr), &product)
	fmt.Printf("Decoded: %+v\n", product)
}
