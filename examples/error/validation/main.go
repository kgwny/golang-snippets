package main

import (
	"fmt"
)

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Msg)
}

func validate(name string) error {
	if name == "" {
		return &ValidationError{"name", "cannot be empty."}
	}
	return nil
}

// 独自のエラー構造体を作成して、状況を含むエラーを返却する
func main() {
	err := validate("")
	if err != nil {
		fmt.Println(err)
	}
}
