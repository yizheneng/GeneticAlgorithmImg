// genetic_algorithm_image.go
package main

import (
	"fmt"
	"math/rand"
)

const (
	TrianglesNum  = 100 /// 每个个体的三角形的数目
	PeoplesNUM    = 100 /// 种群的个体数目
	PictureWidth  = 100 /// 图片的宽度
	PictureHeight = 100 /// 图片的高度
)

type Point struct {
	x uint
	y uint
}

type Color struct {
	r uint8
	g uint8
	b uint8
}

type Triangle struct {
	p1 Point
	p2 Point
	p3 Point

	color Color
}

type GeneticAlgorithmImgDNA struct {
	triangles []Triangle
}

type GeneticAlgorithmImg struct {
	peoples []GeneticAlgorithmImgDNA
}

/// 初始化族群
func NewGeneticAlgorithmImg() *GeneticAlgorithmImg {
	this := &GeneticAlgorithmImg{}

	for i := 0; i < PeoplesNUM; i++ {
		var dna GeneticAlgorithmImgDNA
		for j := 0; j < TrianglesNum; j++ {
			dna.triangles = append(dna.triangles, Triangle{Point{uint(rand.Uint32() % PictureWidth), uint(rand.Uint32() % PictureHeight)},
				Point{uint(rand.Uint32() % PictureWidth), uint(rand.Uint32() % PictureHeight)},
				Point{uint(rand.Uint32() % PictureWidth), uint(rand.Uint32() % PictureHeight)},
				Color{uint8(rand.Uint32() % 255), uint8(rand.Uint32() % 255), uint8(rand.Uint32() % 255)}})
		}
		this.peoples = append(this.peoples, dna)
	}

	fmt.Println(this.peoples)

	return this
}

func (this *GeneticAlgorithmImg) Hello() string {
	return "hello GeneticAlgorithm"
}

/// 族群繁衍,变异
func (this *GeneticAlgorithmImg) MkChilds() {

}

/// 相似度排序
func (this *GeneticAlgorithmImg) Sort() {

}
