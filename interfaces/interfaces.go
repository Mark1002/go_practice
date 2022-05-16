package interfaces

type Human interface {
	walk()
	talk()
	String() string
}

type secreter interface {
	secret()
}
