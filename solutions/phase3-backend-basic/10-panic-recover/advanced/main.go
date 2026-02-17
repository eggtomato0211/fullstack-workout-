package main

import "fmt"

type Job struct {
	Name string
	Fn   func() error
}

// RunJob は1つのジョブを安全に実行する
func RunJob(job Job) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return job.Fn()
}

type BatchResult struct {
	Success int
	Failed  int
	Panics  int
	Errors  []string
}

// RunBatch は複数のジョブを順に実行する
func RunBatch(jobs []Job) BatchResult {
	result := BatchResult{}

	for _, job := range jobs {
		fmt.Printf("実行中: %s ... ", job.Name)
		err := RunJob(job)
		if err != nil {
			result.Failed++
			result.Errors = append(result.Errors, fmt.Sprintf("%s: %v", job.Name, err))
			if err.Error()[:5] == "panic" {
				result.Panics++
			}
			fmt.Println("失敗")
		} else {
			result.Success++
			fmt.Println("成功")
		}
	}

	return result
}

func main() {
	jobs := []Job{
		{Name: "データ取得", Fn: func() error { return nil }},
		{Name: "データ変換", Fn: func() error { return fmt.Errorf("invalid format") }},
		{Name: "データ保存", Fn: func() error { panic("nil pointer dereference") }},
		{Name: "通知送信", Fn: func() error { return nil }},
	}

	result := RunBatch(jobs)
	fmt.Println("\n=== バッチ結果 ===")
	fmt.Printf("成功: %d, 失敗: %d (うちpanic: %d)\n", result.Success, result.Failed, result.Panics)
	for _, e := range result.Errors {
		fmt.Println("  -", e)
	}
}
