package concurrent

import (
	"fmt"
	"strings"
)

type ConCurUpperDemo struct{}

func (demo ConCurUpperDemo) to_upper(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for lang := range in {
			upper_lang := strings.ToUpper(lang)
			out <- upper_lang
		}
		close(out)
	}()
	return out
}

func (demo ConCurUpperDemo) producer(langs []string) <-chan string {
	out := make(chan string)
	go func() {
		for _, lang := range langs {
			out <- lang
		}
		close(out)
	}()
	return out
}

func (demo ConCurUpperDemo) Execute() {
	langs := []string{"python", "java", "ruby", "go", "c++"}
	in := demo.producer(langs)
	out1 := demo.to_upper(in)
	out2 := demo.to_upper(in)
	for lang := range out2 {
		fmt.Println("out2: " + lang)
	}
	for lang := range out1 {
		fmt.Println("out1: " + lang)
	}

	fmt.Println("finish program.")
}
