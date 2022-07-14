package main

import (
	"fmt"
	"os"

	"github.com/mark1002/go_practice/cmd"
	"github.com/mark1002/go_practice/concurrent"
	"github.com/mark1002/go_practice/context"
	"github.com/mark1002/go_practice/error"
	"github.com/mark1002/go_practice/fibclosure"
	"github.com/mark1002/go_practice/hello"
	"github.com/mark1002/go_practice/interfaces"
	"github.com/mark1002/go_practice/maps"
	"github.com/mark1002/go_practice/slices"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("please enter demo program name!")
		os.Exit(0)
	}
	switch args[1] {
	case "panicDemo":
		error.ExecutePanic()
	case "CustomErrorDemo":
		error.ExecuteCustomError()
	case "concurDemo":
		concurrent.Execute()
	case "crawlerDemo":
		concurrent.ExecuteCrawler()
	case "crawlerHttpDemo":
		concurrent.ExecuteHttp()
	case "urlpoll":
		concurrent.ExecuteUrlPoll()
	case "slicesDemo":
		slices.Execute()
	case "helloDemo":
		hello.Execute()
	case "closureDemo":
		fibclosure.Execute()
	case "mapDemo":
		maps.Execute()
	case "cmdDemo":
		cmd.Execute()
	case "contextDemo":
		context.Execute()
	case "contextCancelDemo":
		context.ExecuteCancel()
	case "interfaceDemo":
		interfaces.Execute()
	default:
		fmt.Printf("not found program \"%v\"\n", args[1])
	}
}
