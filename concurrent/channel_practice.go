package concurrent

import (
	"fmt"
	"runtime"
)

type ChannelPracticeDemo struct{}

func (demo ChannelPracticeDemo) Execute() {
	fmt.Println("cpu core num: ", runtime.NumCPU())
	fmt.Println("goroutine num: ", runtime.NumGoroutine())
}
