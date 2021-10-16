package concurrent

import (
	"fmt"
	"strings"
)

func to_upper(in <-chan string) <-chan string {
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

func producer(langs []string) <-chan string {
	out := make(chan string)
	go func() {
		for _, lang := range langs {
			out <- lang
		}
		close(out)
	}()
	return out
}

func Excute() {
	langs := []string{"python", "java", "ruby", "go", "c++"}
	in := producer(langs)
	out1 := to_upper(in)
	out2 := to_upper(in)
	for lang := range out2 {
		fmt.Println("out2: " + lang)
	}
	for lang := range out1 {
		fmt.Println("out1: " + lang)
	}

	fmt.Println("finish program.")
}
