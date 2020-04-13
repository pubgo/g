package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// https://github.com/Nr90/imgsim
// https://github.com/corona10/goimagehash
// https://github.com/teran/imgsum
// https://github.com/nektro/dacite
// https://github.com/nicolasboulay/ghash
// https://github.com/disintegration/imaging
// https://github.com/brett-lempereur/ish
// https://github.com/devedge/imagehash
// https://github.com/devedge/imagehash
// https://github.com/Nr90/imgsim
// https://github.com/dgryski/go-minhash
// https://github.com/hbakhtiyor/strsim
// https://github.com/masatana/go-textdistance
// https://github.com/nivance/image-similarity

func main() {

	//openbit := robotgo.OpenBitmap()
	//xerror.Panic(robotgo.ActiveName("WeChat"))
	//robotgo.ShowAlert("", "")

	//for {
	//	x, y := robotgo.GetMousePos()
	//	fmt.Printf("%d,%d ok\n", x, y)
	//	robotgo.Sleep(1)
	//	//robotgo.ShowAlert("ss", fmt.Sprintf("%d,%d ok", x, y))
	//}

	//x, y := robotgo.GetMousePos()
	//bitmap := robotgo.CaptureScreen(x, y, 200, 50)
	// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	//defer robotgo.FreeBitmap(bitmap)

	//os.Remove("2.png")

	//robotgo.SaveCapture("2.png", 70, 0, 200, 50)

	//fmt.Println("...", bitmap)

	//robotgo.OpenBitmap("3.png")

	//x, y := robotgo.FindBitmap(robotgo.OpenBitmap("3.png"))

	//robotgo.SaveCapture("sss.png")

	//ioutil.WriteFile("a.txt",[]byte(hex.EncodeToString(imgo.ImgToBytes("3.png"))),0755)
	//ioutil.WriteFile("b.txt",[]byte(hex.EncodeToString(imgo.ImgToBytes("sss.png"))),0755)

	//fmt.Println(imgo.ReadPNG("3.png").Bounds())
	//fmt.Println(imgo.ReadPNG("sss.png").Bounds())

	//_img1 := imageutil.Rgb2Gray(imgo.ReadPNG("3.png"))
	//_img2 := imageutil.Rgb2Gray(imgo.ReadPNG("sss.png"))

	//resized := resize.Resize(1, 1, imgo.ReadPNG("3.png"), resize.Bilinear)
	//fmt.Println(imageutil.FlattenPixels(imageutil.Rgb2Gray(resized), 1, 1))
	//fmt.Println(string(xerror.PanicErr(json.MarshalIndent(imageutil.Rgb2Gray(resized)," ","\t")).([]byte)))
	//fmt.Println(imageutil.Rgb2Gray(imgo.ReadPNG("1.png")))
	//fmt.Println(imageutil.Rgb2Gray(imgo.ReadPNG("sss.png")))

	//resized1 := resize.Resize(6, 12, imgo.ReadPNG("sss.png"), resize.Bilinear)
	//fmt.Println(imageutil.Rgb2Gray(resized1))
	//fmt.Printf("%#v\n",imageutil.Rgb2Gray(resized1))

	//fmt.Println(string(xerror.PanicErr(json.MarshalIndent(imageutil.Rgb2Gray(resized1)," ","\t")).([]byte)))

	//for i := 1; i < 4; i++ {
	//fmt.Println(robotgo.FindPic(fmt.Sprintf("%d.png", 3)))
	//}

	x, y := robotgo.FindPic("2.png")
	fmt.Println(x, y)
	//robotgo.ShowAlert("ss", fmt.Sprintf("%d,%d ok", x, y))

	//x1, y1 := robotgo.GetImgSize("1.png")
	//robotgo.ShowAlert("ss", fmt.Sprintf("%d,%d ok", x1, y1))

	//fx, fy := robotgo.FindBitmap(bitmap)
	//bitmap1 := robotgo.CaptureScreen(x, y, 1, 1000)
	//defer robotgo.FreeBitmap(bitmap1)

	// 把bitmap转化为str 然后再转化为bitmap
	//bitmap1 = robotgo.BitmapStr(robotgo.TostringBitmap(bitmap1))
	//fx1, fy1 := robotgo.FindBitmap(bitmap1)
	//fmt.Println("FindBitmap------ ", x, fx, fx1, y, fy, fy1)
	//fmt.Println("FindBitmap------ ", x, fx/2, y, fy1/2)
}
