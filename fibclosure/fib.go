package fibclosure

import "fmt"

func fibonacci() func() int {
	f0 := -1
	f1 := 1
	return func() int {
		// make sure the first return is 0
		if f0 == -1 {
			f0 += 1
			return f0
		}
		temp := f0 + f1
		f0 = f1
		f1 = temp
		return f0
	}
}

func Execute() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
