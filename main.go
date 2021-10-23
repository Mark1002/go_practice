package main

import (
	"github.com/mark1002/go_practice/concurrent"
	"github.com/mark1002/go_practice/error"
	"github.com/mark1002/go_practice/fibclosure"
	"github.com/mark1002/go_practice/hello"
	"github.com/mark1002/go_practice/maps"
	"github.com/mark1002/go_practice/slices"
)

func main() {
	error.Execute()
	concurrent.Execute()
	slices.Execute()
	hello.Execute()
	fibclosure.Execute()
	maps.Execute()
}
