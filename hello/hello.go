package hello

import "fmt"

type GoParams struct {
	Name    string
	Version string
}

func SayHello(name string) string {
	return name
}

func GoVersion(p GoParams) string {
	if p.Name == "" {
		p.Name = "golang"
	}
	if p.Version == "" {
		p.Version = "1.20"
	}
	v := fmt.Sprintf("%s %s", p.Name, p.Version)
	return v
}
