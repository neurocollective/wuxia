package components

import "fmt"

type Components struct {
	stuff string
}

func BuildComponents(components any) Components {
	fmt.Println("this is where i will build component definitions in memory?")

	return Components{}
}