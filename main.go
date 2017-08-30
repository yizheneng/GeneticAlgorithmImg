// main.go
package main

type GeneticAlgorithm interface {
	Hello() string
}

func main() {
	a := NewGeneticAlgorithmImg()

	for i := 0; i < 3000; i++ {
		a.Sort()
		a.MkChilds()
	}
}
