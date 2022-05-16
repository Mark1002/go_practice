package interfaces

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (p Person) walk() {
	fmt.Println("i can walk")
}

func (p Person) secret() {
	fmt.Println("no one know the secret")
}

func (p Person) talk() {
	fmt.Println("i can talk")
}

func introducteYourSelf(h Human) {
	h.walk()
	h.talk()
	fmt.Println(h)
}

func onlyWalk(walker interface{ walk() }) {
	walker.walk()
}

func talkSecret(s secreter) {
	s.secret()
}

func Execute() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	introducteYourSelf(a)
	introducteYourSelf(z)
	onlyWalk(a)
	talkSecret(z)
}
