package main

import(
	"fmt"
)

type Demo struct {
	Name string
	Age int
	Sex string
}

func NewDemo() *Demo {
	demo := new(Demo)
	demo.Name = "xiaoming"
	demo.Age = 15
	demo.Sex = "famle"
	return demo
}

func main() {
	a := NewDemo()
	fmt.Println(a.Name)
}


