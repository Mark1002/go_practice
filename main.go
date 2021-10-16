package main

import (
	"github.com/mark1002/go_practice/concurrent"
	"github.com/mark1002/go_practice/error"
	"github.com/mark1002/go_practice/hello"
	"github.com/mark1002/go_practice/slices"
)

func main() {
	error.Excute()
	concurrent.Excute()
	slices.Excute()
	hello.Excute()
}
