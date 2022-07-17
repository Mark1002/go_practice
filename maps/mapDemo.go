package maps

import "fmt"

type MapDemo struct{}

func (demo MapDemo) Execute() {
	userIDMap := map[string]int{
		"Peter": 1234,
		"Tony":  4321,
	}

	id, ok := userIDMap["Mark"]
	if ok {
		fmt.Println("userid is:", id)
	} else {
		fmt.Println("userid not found!")
	}
}
