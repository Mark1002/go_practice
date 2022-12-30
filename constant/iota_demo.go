package constant

import "fmt"

type IotaDemo struct{}

// https://github.com/golang/go/wiki/Iota
const (
	_ = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func (demo IotaDemo) Execute() {
	fmt.Printf("%d\t\t%b\n", KB, KB)
	fmt.Printf("%d\t\t%b\n", MB, MB)
	fmt.Printf("%d\t\t%b\n", GB, GB)
	fmt.Printf("%d\t\t%b\n", TB, TB)
	fmt.Printf("%d\t\t%b\n", PB, PB)
}
