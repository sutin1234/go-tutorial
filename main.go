package main

import (
	"fmt"

	"github.com/sutin1234/go-tutorial/banana"
	"github.com/sutin1234/go-tutorial/hello"
)

func main() {
	fmt.Println("hello", hello.SayHello("world"))

	// // DateFormat
	// format_date := strings.ReplaceAll(cur, "-", "/")
	// fmt.Println(date.DateFormat(format_date, "2006/01/02"))
	fmt.Println("test BANANA ==> ", banana.Solution("BANANA"))                       // 1
	fmt.Println("test NAABXXAN ==> ", banana.Solution("NAABXXAN"))                   // 1
	fmt.Println("test NAANAAXNABABYNNBZ ==> ", banana.Solution("NAANAAXNABABYNNBZ")) // 2
	fmt.Println("test QABAAWOBL ==> ", banana.Solution("QABAAWOBL"))                 // 0
}
