package error

import (
	"fmt"
)

type MarkError struct {
	Message string
}

// 繼承 error interface
func (e *MarkError) Error() string {
	return fmt.Sprintf("error in mark: %v", e.Message)
}

func throwError() error {
	return &MarkError{Message: "this is custom error!"}
}

func Foo() {
	panic("fatal error!")
}

func Execute() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("internal error: %v", p)
		}
	}()
	err := throwError()
	if err != nil {
		fmt.Println(err)
	}
	Foo()
	fmt.Println("program execute final")
}
