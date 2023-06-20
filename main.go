package main

import (
	"fmt"

	"github.com/sutin1234/go-tutorial/game"
)

func main() {
	// fmt.Println("hello", hello.SayHello("world"))
	s, e := game.Solution("20:00", "22:00")
	fmt.Printf("timeStart: %v, timrEnd: %v\n", s, e)
}
