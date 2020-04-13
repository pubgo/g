package main

import (
	"github.com/vcaesar/imgo"
	"golang.org/x/image/bmp"
	"log"
	"os"
)

func main() {
	imgo.Save("2.png", imgo.ToBytes(imgo.ReadPNG("1.png"), "bmp"))

	f, err := os.Open("2.png")
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	img, derr := bmp.Decode(f)
	if nil != derr {
		log.Println(derr)
	}

	imgo.Save("3.png", imgo.ToBytes(img, "png"))

	//buf := new(bytes.Buffer)
	//xerror.Panic(bmp.Encode(buf, imgo.ReadPNG("1.png")))

	//buf1 := new(bytes.Buffer)
	//xerror.Panic(png.Encode(buf1, xerror.PanicErr(png.Decode(buf)).(image.Image)))

	//imgo.Save("3.png", buf1.Bytes())
	//imgo.DecodeFile()
	//img, s, err := robotgo.DecodeImg("2.png")
	//xerror.Panic(err)
	//fmt.Println(s, img.ColorModel())

	//image.NewRGBA()
}
