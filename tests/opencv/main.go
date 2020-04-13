package main

import (
	"fmt"
	"github.com/vcaesar/imgo"
	"gocv.io/x/gocv"
	"image"
	"time"
)

func main1() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello")
	defer window.Close()

	window.SetWindowProperty(gocv.WindowPropertyVisible, gocv.WindowNormal)

	img := gocv.NewMat()

	//_m := gocv.IMRead("1.png", gocv.IMReadAnyColor)
	for {
		webcam.Read(&img)
		//window.IMShow(_m)
		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}
	}
}

func main() {
	//si := contrib.NewSIFT()
	//_m1 := si.Detect(gocv.IMRead("1.png", gocv.IMReadColor))
	//_m2 := si.Detect(gocv.IMRead("2.png", gocv.IMReadColor))
	//fmt.Println(_m1)
	//fmt.Println(_m2)

	//xerror.Panic(robotgo.ActiveName("WeChat"))
	time.Sleep(time.Second)
	imgScene := gocv.IMRead("1.png", gocv.IMReadColor)
	if imgScene.Empty() {
		fmt.Printf("Invalid read of face.jpg in MatchTemplate test")
	}
	fmt.Println(imgScene.ToImage())
	defer imgScene.Close()

	//_name := fmt.Sprintf("%d.png", time.Now().Unix())
	//robotgo.SaveCapture(_name)
	//shutil.Execute(fmt.Sprintf("open %s", _name))
	imgTemplate := gocv.IMRead("2.png", gocv.IMReadColor)
	if imgTemplate.Empty() {
		fmt.Printf("Invalid read of toy.jpg in MatchTemplate test")
	}
	defer imgTemplate.Close()

	result := gocv.NewMat()
	defer result.Close()
	m := gocv.NewMat()
	gocv.MatchTemplate(imgScene, imgTemplate, &result, gocv.TmCcoeffNormed, m)
	m.Close()

	fmt.Println(gocv.MinMaxLoc(result))
	_, maxConfidence, min, max := gocv.MinMaxLoc(result)
	if maxConfidence < 0.95 {
		fmt.Printf("Max confidence of %f is too low. MatchTemplate could not find template in scene.", maxConfidence)
	}

	ss1 := imgTemplate.Region(image.Rect(min.X, min.Y, max.X, max.Y))
	ss2, _ := (&ss1).ToImage()
	imgo.SaveToPNG("sss.png", ss2)

	//os.Remove(_name)

	//robotgo.MoveClick(min.X/2+10, min.Y/2+10)

}
