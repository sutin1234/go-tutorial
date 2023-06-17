package hello

import (
	"fmt"

	Struct "github.com/sutin1234/go-tutorial/struct"
)

func SayHello(name string) string {
	return name
}

func GoVersion(p Struct.GoParams) string {
	if p.Name == "" {
		p.Name = "golang"
	}
	if p.Version == "" {
		p.Version = "1.20"
	}
	v := fmt.Sprintf("%s %s", p.Name, p.Version)
	return v
}
