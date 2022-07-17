package main

import (
	"fmt"
	"os"

	"github.com/mark1002/go_practice/base"
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
	var demo base.GoPractice

	if len(args) == 1 {
		fmt.Println("please enter demo program name!")
		os.Exit(0)
	}
	switch args[1] {
	case "panicDemo":
		demo = error.PanicErrorDemo{}
		demo.Execute()
	case "CustomErrorDemo":
		demo = error.CustomErrorDemo{}
		demo.Execute()
	case "concurDemo":
		demo = concurrent.ConCurUpperDemo{}
		demo.Execute()
	case "crawlerDemo":
		demo = concurrent.FakeCrawlerDemo{}
		demo.Execute()
	case "crawlerHttpDemo":
		demo = concurrent.HttpCrawlerDemo{}
		demo.Execute()
	case "urlpoll":
		demo = concurrent.UrlPollDemo{}
		demo.Execute()
	case "slicesDemo":
		demo = slices.SliceDemo{}
		demo.Execute()
	case "helloDemo":
		hello.Execute()
	case "closureDemo":
		fibclosure.Execute()
	case "mapDemo":
		maps.Execute()
	case "cmdDemo":
		cmd.Execute()
	case "contextDemo":
		demo = context.ContextDemo{}
		demo.Execute()
	case "contextCancelDemo":
		demo = context.ContextCancelDemo{}
		demo.Execute()
	case "interfaceDemo":
		interfaces.Execute()
	default:
		fmt.Printf("not found program \"%v\"\n", args[1])
	}
}
