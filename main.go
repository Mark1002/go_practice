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
	"github.com/mark1002/go_practice/generics"
	"github.com/mark1002/go_practice/hello"
	"github.com/mark1002/go_practice/interfaces"
	"github.com/mark1002/go_practice/maps"
	"github.com/mark1002/go_practice/slices"
)

func main() {
	args := os.Args
	var demo base.GoPracticer

	if len(args) == 1 {
		fmt.Println("please enter demo program name!")
		os.Exit(0)
	}
	switch args[1] {
	case "panicDemo":
		demo = error.PanicErrorDemo{}
	case "CustomErrorDemo":
		demo = error.CustomErrorDemo{}
	case "concurDemo":
		demo = concurrent.ConCurUpperDemo{}
	case "crawlerDemo":
		demo = concurrent.FakeCrawlerDemo{}
	case "crawlerHttpDemo":
		demo = concurrent.HttpCrawlerDemo{}
	case "urlpoll":
		demo = concurrent.UrlPollDemo{}
	case "slicesDemo":
		demo = slices.SliceDemo{}
	case "helloDemo":
		demo = hello.HelloDemo{}
	case "closureDemo":
		demo = fibclosure.FibClosureDemo{}
	case "mapDemo":
		demo = maps.MapDemo{}
	case "cmdDemo":
		demo = cmd.CmdDemo{}
	case "contextDemo":
		demo = context.ContextDemo{}
	case "contextCancelDemo":
		demo = context.ContextCancelDemo{}
	case "interfaceDemo":
		demo = interfaces.InterfaceDemo{}
	case "genericDemo":
		demo = generics.GenericDemo{}
	default:
		demo = base.NotFound{DemoName: args[1]}
	}
	demo.Execute()
}
