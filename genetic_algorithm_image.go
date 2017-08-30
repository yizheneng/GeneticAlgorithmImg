// genetic_algorithm_image.go
package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
)

const (
	TrianglesNum   = 100 /// 每个个体的三角形的数目
	PeoplesNUM     = 100 /// 种群的个体数目
	OutScale       = 0.9 /// 每次淘汰的比例
	PictureWidth   = 268 /// 图片的宽度
	PictureHeight  = 268 /// 图片的高度
	VariationLimit = 0.1 /// 变异上限
	SrcImgPath     = "e850352ac65c1038f5e242beb0119313b07e892a.jpg"
)

type Point struct {
	x uint
	y uint
}

type Color struct {
	r int
	g int
	b int
	a int
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
	peoples   []GeneticAlgorithmImgDNA
	srcimg    *gui.QImage
	numerical int
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
				Color{int(rand.Uint32() % 255), int(rand.Uint32() % 255), int(rand.Uint32() % 255), int(rand.Uint32() % 255)}})
		}
		this.peoples = append(this.peoples, dna)

		//		img := DrawTriangle(dna)
		//		img.Save(fmt.Sprintf("test/%d.jpg", i), "jpg", -1)

	}

	this.srcimg = gui.NewQImage9(SrcImgPath, "")
	fmt.Println(this.srcimg.Width(), this.srcimg.Height())
	this.srcimg.Save("nihao.jpg", "jpg", -1)

	return this
}

func (this *GeneticAlgorithmImg) Hello() string {
	return "hello GeneticAlgorithm"
}

/// 族群繁衍,变异
func (this *GeneticAlgorithmImg) MkChilds() {
	var babys []GeneticAlgorithmImgDNA
	for i := 0; i < PeoplesNUM; i++ {
		father := rand.Uint32() % (PeoplesNUM * OutScale)
		mother := rand.Uint32() % (PeoplesNUM * OutScale)
		baby := this.peoples[father].triangles[:TrianglesNum/2]
		baby = append(baby, this.peoples[mother].triangles[TrianglesNum/2:]...)
		babys = append(babys, GeneticAlgorithmImgDNA{baby})
	}

	babys[rand.Uint32()%(PeoplesNUM*OutScale)].triangles[rand.Uint32()%TrianglesNum].p1 = Point{uint(rand.Uint32() % PictureWidth), uint(rand.Uint32() % PictureHeight)}
	babys[rand.Uint32()%(PeoplesNUM*OutScale)].triangles[rand.Uint32()%TrianglesNum].p2 = Point{uint(rand.Uint32() % PictureWidth), uint(rand.Uint32() % PictureHeight)}
	babys[rand.Uint32()%(PeoplesNUM*OutScale)].triangles[rand.Uint32()%TrianglesNum].p3 = Point{uint(rand.Uint32() % PictureWidth), uint(rand.Uint32() % PictureHeight)}
	babys[rand.Uint32()%(PeoplesNUM*OutScale)].triangles[rand.Uint32()%TrianglesNum].color = Color{int(rand.Uint32() % 255), int(rand.Uint32() % 255), int(rand.Uint32() % 255), int(rand.Uint32() % 255)}

	this.peoples = babys
}

/// 相似度排序
func (this *GeneticAlgorithmImg) Sort() {
	prioritys := []int{}

	for _, people := range this.peoples {
		img := DrawTriangle(people)
		prioritys = append(prioritys, SimilarityCalculation(this.srcimg, img))
		//		img.DestroyQImage()
		img = gui.NewQImage3(PictureWidth, PictureHeight, gui.QImage__Format_RGB888)
	}

	for i := 1; i < len(prioritys); i++ {
		for j := 0; j < (len(prioritys) - i); j++ {
			if prioritys[j] > prioritys[j+1] {
				temp := prioritys[j]
				prioritys[j] = prioritys[j+1]
				prioritys[j+1] = temp

				temp_people := this.peoples[j]
				this.peoples[j] = this.peoples[j+1]
				this.peoples[j+1] = temp_people
			}
		}
	}
	this.numerical++
	DrawTriangle(this.peoples[0]).Save(fmt.Sprintf("test/%d.jpg", this.numerical), "jpg", -1)
	fmt.Println("generation:", this.numerical)
	fmt.Println("      priority:", prioritys[0])
}

func DrawTriangle(dna GeneticAlgorithmImgDNA) (img *gui.QImage) {
	img = gui.NewQImage3(PictureWidth, PictureHeight, gui.QImage__Format_RGB888)
	img.Fill2(gui.NewQColor3(255, 255, 255, 255))
	drawer := gui.NewQPainter2(img)

	for _, triangle := range dna.triangles {
		drawer.SetBrush(gui.NewQBrush3(gui.NewQColor3(triangle.color.r, triangle.color.g, triangle.color.b, triangle.color.a), core.Qt__SolidPattern))
		drawer.SetPen2(gui.NewQColor3(0, 0, 0, 0))
		points := []*core.QPointF{}
		points = append(points, core.NewQPointF3(float64(triangle.p1.x), float64(triangle.p1.y)), core.NewQPointF3(float64(triangle.p2.x), float64(triangle.p2.y)), core.NewQPointF3(float64(triangle.p3.x), float64(triangle.p3.y)))
		drawer.DrawPolygon2(gui.NewQPolygonF3(points), core.Qt__OddEvenFill)
	}

	return
}

/// 相似度计算值越小相似度越高
func SimilarityCalculation(img1, img2 *gui.QImage) int {
	temp := 0
	for w := 0; w < img1.Width(); w++ {
		for h := 0; h < img1.Height(); h++ {
			color1 := img1.PixelColor2(w, h)
			color2 := img2.PixelColor2(w, h)
			temp += int(math.Abs(float64(color1.Red() - color2.Red())))
			temp += int(math.Abs(float64(color1.Green() - color2.Green())))
			temp += int(math.Abs(float64(color1.Blue() - color2.Blue())))
		}
	}

	return temp
}
