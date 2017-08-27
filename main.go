// main.go
package main

import (
	"fmt"
)

type GeneticAlgorithm interface {
	Hello() string
}

func main() {
	var a GeneticAlgorithm

	a = NewGeneticAlgorithmImg()

	fmt.Println("Hello World!  ", a.Hello())
}
