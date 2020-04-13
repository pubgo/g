package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	"github.com/pubgo/g/xerror"
	hook "github.com/robotn/gohook"
	"log"
	"testing"
	"time"
)

// 组件对象， 对象的坐标
// 获取组件的对象和坐标
// 关联鼠标事件和键盘事件等
// 下拉截止判断 通过获取当前鼠标的横竖区域的颜色，然后得到一个hash值，如果该hash值在一个范围内不变，则最后了

type Point struct {
	X float64
	Y float64
}

func point(x, y float64) Point {
	return Point{X: x, Y: y}
}

func move() {
	robotgo.Move(100, 200)

	// move the mouse to 100, 200
	robotgo.MoveMouse(100, 200)

	robotgo.Drag(10, 10)
	robotgo.Drag(20, 20, "right")
	//
	robotgo.DragSmooth(10, 10)
	robotgo.DragSmooth(100, 200, 1.0, 100.0)

	// smooth move the mouse to 100, 200
	robotgo.MoveSmooth(100, 200)
	robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)

	xerror.Panic(clipboard.WriteAll("日本語"))
	text, err := clipboard.ReadAll()
	if err != nil {
		log.Println("clipboard read all error: ", err)
	} else {
		if text != "" {
			log.Println("text is: ", text)
		}
	}

	for i := 0; i < 1080; i += 1000 {
		fmt.Println(i)
		robotgo.MoveMouse(800, i)
	}
}

func click() {

	// click the left mouse button
	robotgo.Click()

	// click the right mouse button
	robotgo.Click("right", false)

	// double click the left mouse button
	robotgo.MouseClick("left", true)
}

func toggleAndScroll() {
	// scrolls the mouse either up
	robotgo.ScrollMouse(2, "down")

	//for i:=0;i<2;i++{
	robotgo.Scroll(0, 10)
	//}

	// toggles right mouse button
	//robotgo.MouseToggle("down", "right")

	//robotgo.MouseToggle("up")
}

func get() {
	// gets the mouse coordinates
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
	if x == 456 && y == 586 {
		fmt.Println("mouse...", "586")
	}

	robotgo.MoveMouse(x, y)
}

func toggleAndScroll1() {
	// scrolls the mouse either up
	robotgo.ScrollMouse(gzhListStep, "up")
	//robotgo.Scroll(100, 200)

	// toggles right mouse button
	robotgo.MouseToggle("down", "right")

	//robotgo.MouseToggle("up")
}

func mouse() {
	fmt.Println("start")
	move()
	fmt.Println("mouse")

	click()

	fmt.Println("click")
	get()

	fmt.Println("get")

	toggleAndScroll()
	fmt.Println("toggleAndScroll")
}

func bitmaps() {
	x, y := robotgo.GetMousePos()
	bitmap := robotgo.CaptureScreen(x, y, 1000, 1)
	// use `defer robotgo.FreeBitmap(bit)` to free the bitmap
	defer robotgo.FreeBitmap(bitmap)

	//fmt.Println("...", bitmap)

	fx, fy := robotgo.FindBitmap(bitmap)

	bitmap1 := robotgo.CaptureScreen(x, y, 1, 1000)
	defer robotgo.FreeBitmap(bitmap1)

	// 把bitmap转化为str 然后再转化为bitmap
	bitmap1 = robotgo.BitmapStr(robotgo.TostringBitmap(bitmap1))

	fx1, fy1 := robotgo.FindBitmap(bitmap1)
	fmt.Println("FindBitmap------ ", x, fx, fx1, y, fy, fy1)
	fmt.Println("FindBitmap------ ", x, fx/2, y, fy1/2)
}

// 测试选中
func xuanzhong() {
	//s := robotgo.Start()
	//defer robotgo.End()
}
func TestA1(t *testing.T) {
	//nps, err := robotgo.Process()
	//errors.Panic(err)
	//for _, p := range nps {
	//	fmt.Println(p)
	//}

	title := robotgo.GetTitle()
	fmt.Println("title@@@ ", title)

	xerror.Panic(robotgo.ActiveName(activeName))

	//errors.Panic(clipboard.WriteAll("日本語"))
	//text, err := clipboard.ReadAll()
	//if err != nil {
	//	log.Println("clipboard read all error: ", err)
	//} else {
	//	if text != "" {
	//		log.Println("text is: ", text)
	//	}
	//}

	//mouse()

	//_t := point(537, 411)

	//fmt.Println(robotgo.KeyTap("f", "lctrl", "lcmd"))

	robotgo.MoveMouseSmooth(int(gzhFirstList.X), int(gzhFirstList.Y))

	_ii := 0

	// 开始事件监听
	s := robotgo.Start()
	defer robotgo.End()

	go func() {
		fmt.Println(robotgo.AddEvents("s", "command"))
	}()

	go func() {
		a := 0
		for e := range s {
			//fmt.Println(e.String())
			//fmt.Println(e.Keycode, e.Keychar)

			if e.Kind == hook.MouseDown {
				a += int(hook.MouseDown) * int(hook.MouseDown)
			}

			if e.Kind == hook.MouseUp {
				a += int(hook.MouseUp) * int(hook.MouseUp)
			}

			if a == 100 {
				fmt.Println(e.String())
				a = 0
			}

			//robotgo.ShowAlert("sss", "dd")
			//robotgo.TypeString("okkkkk")
		}
	}()

	for i := 0; i < 100000; i++ {

		if _ii == 0 {
			_ii = 4
			//x, y := robotgo.GetMousePos()
			//robotgo.MoveMouseSmooth(x, y+_ii*105)
		}

		//robotgo.ScrollMouse(10, "down")
		//robotgo.SetMouseDelay(1000)
		//robotgo.Scroll(0, gzhListStep)

		// 获取当前的坐标
		//x, y := robotgo.GetMousePos()
		//fmt.Println(x, y)

		//robotgo.MoveMouse(int(gzhFirstList.X), int(gzhFirstList.Y)+105)
		//robotgo.MouseToggle("down")
		//robotgo.DragMouse(int(gzhFirstList.X), int(gzhFirstList.Y))
		//robotgo.MouseToggle("up")

		//bitmaps()

		//fmt.Println("history")
		//fmt.Println(robotgo.GetMousePos())
		//fmt.Println(robotgo.FindPic("/Users/barry/gopath/src/github.com/pubgo/mycli/internal/robot/history.png"))

		//robotgo.SaveCapture("a")
		//robotgo.SaveCapture()

		//x, y := robotgo.GetMousePos()
		//bitmap1 := robotgo.CaptureScreen(x, y, 60, 15)
		//defer robotgo.FreeBitmap(bitmap1)
		//robotgo.SaveBitmap(bitmap1, "test.png")

		//toggleAndScroll()

		time.Sleep(time.Millisecond * 10000)
		//fmt.Println(robotgo.IsValid())
		//fmt.Println(robotgo.GetActive())
		//fmt.Println(robotgo.FindNames())
		_ii -= 1
	}

	//s := robotgo.Start()
	//defer robotgo.End()
	//for e := range s {
	//errors.Panic(clipboard.WriteAll(e.String()))
	//text, err := clipboard.ReadAll()
	//if err != nil {
	//	log.Println("clipboard read all error: ", err)
	//} else {
	//	if text != "" {
	//cs, _ := charsetutil.GuessString(text)
	//text = charsetutil.MustDecodeString(text, cs.Charset())
	//fmt.Println("text is: ", text)
	//}
	//}

	// color := robotgo.GetMouseColor()
	// fmt.Println("color---- ", color)

	//

	// bitmap := robotgo.CaptureScreen(10, 20, 1500, 30)
	// defer robotgo.FreeBitmap(bitmap)

	// s := sha1.New()
	// s.Write([]byte(robotgo.TostringBitmap(bitmap)))
	// fmt.Println(hex.EncodeToString(s.Sum(nil)))

	//robotgo.CountColor()
	//robotgo.GoCaptureScreen()

	//robotgo.ShowAlert("dd","ggg")

	//robotgo.ScrollMouse(10, "down")

	//fmt.Println("hook: ", e.String())

	//}
}
