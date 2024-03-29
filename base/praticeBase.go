package base

import "fmt"

type GoPracticer interface {
	Execute()
}

type NotFound struct {
	DemoName string
}

func (demo NotFound) Execute() {
	fmt.Printf("not found program \"%v\"\n", demo.DemoName)
}
