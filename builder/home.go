package main

import "fmt"

type Home struct {
}

func (home Home) build(builder HomeBuilder) {
	fmt.Printf("The home was successfully built : %+v\n", builder)
}

func NewHome() *Home {
	return &Home{}
}
