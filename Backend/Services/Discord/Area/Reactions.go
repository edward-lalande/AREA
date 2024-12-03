package area

import "fmt"

func HelloWorld() {
	fmt.Println("Hello World")
}

func FindReactions(id int) {
	var Reactions map[int]func() = make(map[int]func())
	Reactions[0] = HelloWorld
	Reactions[id]()
}
