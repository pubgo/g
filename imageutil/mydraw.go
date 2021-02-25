package imageutil

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type Pen struct {
	FontSize   float64
	Dpi        float64
	Font       *truetype.Font
	StartPoint image.Point
	Color      *image.Uniform
}

type HDC struct {
	//Bg   image.Image
	Rgba *image.RGBA
}

//获取画笔
func OnGetPen(fontPath string, R, G, B, A uint8) (pen Pen, b bool) {
	b = false
	pen.Color = image.NewUniform(color.RGBA{R: R, G: G, B: B, A: A})
	pen.Dpi = 72
	pen.FontSize = 10
	pen.StartPoint = image.Point{0, 0}
	// 读字体数据
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		log.Println(err)
		return
	}
	pen.Font, err = freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	b = true
	return
}

func (this *HDC) SetBg(imagePath string) bool {
	file, _ := os.Open(imagePath)

	//var err error
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("err = ", err)
		return false
	}

	this.Rgba = image.NewRGBA(img.Bounds())
	draw.Draw(this.Rgba, this.Rgba.Bounds(), img, image.ZP, draw.Src)
	return true
}

func (this *HDC) GetBgSize() (w, h int) {
	b := this.Rgba.Bounds()
	w = b.Max.X
	h = b.Max.Y
	return
}

//图片上画文字
func (this *HDC) DrawText(pen Pen, text string) bool {
	if this.Rgba == nil {
		return false
	}

	c := freetype.NewContext()
	c.SetDPI(pen.Dpi)
	c.SetFont(pen.Font)
	c.SetFontSize(pen.FontSize)
	c.SetClip(this.Rgba.Bounds())
	c.SetDst(this.Rgba)
	//c.SetSrc(image.NewUniform(color.RGBA{255, 255, 255, 255}))
	c.SetSrc(pen.Color)

	// Draw the text.
	pt := freetype.Pt(pen.StartPoint.X, pen.StartPoint.Y+int(c.PointToFixed(pen.FontSize)>>6))
	for _, s := range strings.Split(text, "\r\n") {
		_, err := c.DrawString(s, pt)
		if err != nil {
			log.Printf("c.DrawString(%s) error(%v)\n", s, err)
			return false
		}
		pt.Y += c.PointToFixed(pen.FontSize * 1.5)
	}
	return false
}

//保存图片
func (this *HDC) Save(imagePath string) bool {
	output, err := os.OpenFile(imagePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Print(err)
		return false
	}

	if strings.HasSuffix(imagePath, ".png") || strings.HasSuffix(imagePath, ".PNG") {
		err = png.Encode(output, this.Rgba)
	} else {
		err = jpeg.Encode(output, this.Rgba, nil)
	}
	if err != nil {

		log.Printf("image encode error(%v)", err)
		//mylog.Error(err)
		return false
	}
	return true
}

// Rgb2Gray function converts RGB to a gray scale array.
func Rgb2Gray(colorImg image.Image) [][]float64 {
	bounds := colorImg.Bounds()
	w, h := bounds.Max.X-bounds.Min.X, bounds.Max.Y-bounds.Min.Y
	pixels := make([][]float64, h)

	for i := range pixels {
		pixels[i] = make([]float64, w)
		for j := range pixels[i] {
			r, g, b, _ := colorImg.At(j, i).RGBA()
			pixels[i][j] = 0.299*float64(r/257) + 0.587*float64(g/257) + 0.114*float64(b/256)
		}
	}

	return pixels
}

// FlattenPixels function flattens 2d array into 1d array.
func FlattenPixels(pixels [][]float64, x int, y int) []float64 {
	flattens := make([]float64, x*y)
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			flattens[y*i+j] = pixels[i][j]
		}
	}
	return flattens
}

// MeanOfPixels function returns a mean of pixels.
func MeanOfPixels(pixels []float64) float64 {
	m := 0.0
	lens := len(pixels)
	if lens == 0 {
		return 0
	}

	for _, p := range pixels {
		m += p
	}

	return m / float64(lens)
}

// MedianOfPixels function returns a median value of pixels.
// It uses quick selection algorithm.
func MedianOfPixels(pixels []float64) float64 {
	tmp := make([]float64, len(pixels))
	copy(tmp, pixels)
	l := len(tmp)
	pos := l / 2
	v := quickSelectMedian(tmp, 0, l-1, pos)
	return v
}

func quickSelectMedian(sequence []float64, low int, hi int, k int) float64 {
	if low == hi {
		return sequence[k]
	}

	for low < hi {
		pivot := low/2 + hi/2
		pivotValue := sequence[pivot]
		storeIdx := low
		sequence[pivot], sequence[hi] = sequence[hi], sequence[pivot]
		for i := low; i < hi; i++ {
			if sequence[i] < pivotValue {
				sequence[storeIdx], sequence[i] = sequence[i], sequence[storeIdx]
				storeIdx++
			}
		}
		sequence[hi], sequence[storeIdx] = sequence[storeIdx], sequence[hi]
		if k <= storeIdx {
			hi = storeIdx
		} else {
			low = storeIdx + 1
		}
	}

	if len(sequence)%2 == 0 {
		return sequence[k-1]/2 + sequence[k]/2
	}
	return sequence[k]
}
