package main

import "fmt"

// AppError はHTTPステータスコード付きのエラー
type AppError struct {
	Code    int
	Message string
	Detail  string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%d] %s: %s", e.Code, e.Message, e.Detail)
}

func NewBadRequestError(detail string) *AppError {
	return &AppError{Code: 400, Message: "Bad Request", Detail: detail}
}

func NewNotFoundError(detail string) *AppError {
	return &AppError{Code: 404, Message: "Not Found", Detail: detail}
}

func NewInternalError(detail string) *AppError {
	return &AppError{Code: 500, Message: "Internal Server Error", Detail: detail}
}

var products = map[int]string{1: "Go入門書"}

func getProduct(id int) (string, error) {
	if id <= 0 {
		return "", NewBadRequestError(fmt.Sprintf("invalid product id: %d", id))
	}
	name, ok := products[id]
	if !ok {
		return "", NewNotFoundError(fmt.Sprintf("product id=%d", id))
	}
	return name, nil
}

func handleRequest(productID int) {
	product, err := getProduct(productID)
	if err != nil {
		if appErr, ok := err.(*AppError); ok {
			fmt.Printf("HTTP %d: %s (%s)\n", appErr.Code, appErr.Message, appErr.Detail)
		} else {
			fmt.Printf("HTTP 500: %v\n", err)
		}
		return
	}
	fmt.Printf("HTTP 200: %s\n", product)
}

func main() {
	handleRequest(1)  // 200
	handleRequest(99) // 404
	handleRequest(-1) // 400
}
