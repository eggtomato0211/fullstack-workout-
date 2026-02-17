package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID         string
	committed  bool
	rolledback bool
}

func BeginTx() *Transaction {
	id := fmt.Sprintf("tx_%d", time.Now().UnixNano()%10000)
	fmt.Printf("[%s] BEGIN\n", id)
	return &Transaction{ID: id}
}

func (tx *Transaction) Commit() error {
	if tx.committed {
		return fmt.Errorf("already committed")
	}
	if tx.rolledback {
		return fmt.Errorf("already rolled back")
	}
	tx.committed = true
	fmt.Printf("[%s] COMMIT\n", tx.ID)
	return nil
}

func (tx *Transaction) Rollback() error {
	if tx.committed {
		fmt.Printf("[%s] ROLLBACK skipped (already committed)\n", tx.ID)
		return nil
	}
	if tx.rolledback {
		return nil
	}
	tx.rolledback = true
	fmt.Printf("[%s] ROLLBACK\n", tx.ID)
	return nil
}

// ExecuteInTx はトランザクション内で処理を実行する
func ExecuteInTx(fn func(tx *Transaction) error) error {
	tx := BeginTx()

	defer func() {
		if tx.committed {
			return
		}
		tx.Rollback()
	}()

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit()
}

func main() {
	// 正常ケース
	fmt.Println("=== 正常ケース ===")
	err := ExecuteInTx(func(tx *Transaction) error {
		fmt.Printf("[%s] INSERT INTO users ...\n", tx.ID)
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	// エラーケース
	fmt.Println("\n=== エラーケース ===")
	err = ExecuteInTx(func(tx *Transaction) error {
		fmt.Printf("[%s] INSERT INTO users ...\n", tx.ID)
		return fmt.Errorf("duplicate key")
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
}
