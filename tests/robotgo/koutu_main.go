package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/pubgo/g/pkg/shutil"
	hook "github.com/robotn/gohook"
	"time"
)

type Bitmap struct {
	ImgBuf        string
	Width         int
	Height        int
	Bytewidth     int
	BitsPixel     uint8
	BytesPerPixel uint8
}

func main1() {

	var s1 = false
	s := robotgo.Start()

	for {
		fmt.Println("ok")

		ct := false
		k := 0
		arr := []string{"ctrl", "command"}
		key := "a"
		for {
			e := <-s

			l := len(arr)
			if l > 0 {
				for i := 0; i < l; i++ {
					ukey := uint16(robotgo.Keycode[arr[i]].(int))

					if e.Kind == hook.KeyHold && e.Keycode == ukey {
						k++
					}

					if k == l {
						ct = true
					}

					if e.Kind == hook.KeyUp && e.Keycode == ukey {
						if k > 0 {
							k--
						}
						// time.Sleep(10 * time.Microsecond)
						ct = false
					}
				}
			} else {
				ct = true
			}

			if ct && e.Kind == hook.KeyUp && e.Keycode == uint16(robotgo.Keycode[key].(int)) {
				s1 = true
				break
			}
		}

		fmt.Println("ctrl + a")

		if s1 {
			func() {
				var h hook.Event
				var up hook.Event
				for ev := range s {
					if ev.Kind == hook.MouseHold {
						fmt.Println(ev)
						if h.X == 0 {
							h = ev
						}
					}

					if ev.Kind == hook.MouseUp {
						fmt.Println(ev)
						if up.X == 0 {
							up = ev

							fmt.Println(h, up, int(h.X), int(h.Y), int(up.X-h.X), int(up.Y-h.Y))
							_bm := robotgo.CaptureScreen(int(h.X), int(h.Y), int(up.X-h.X), int(up.Y-h.Y))
							robotgo.SaveCapture("sss.png", int(h.X), int(h.Y), int(up.X-h.X), int(up.Y-h.Y))
							fmt.Println("CaptureScreen")
							tbm := robotgo.ToBitmap(_bm)
							fmt.Println("ToBitmap")

							bbm := &Bitmap{}
							bbm.Height = tbm.Height
							bbm.Width = tbm.Width
							bbm.BitsPixel = tbm.BitsPixel
							bbm.BytesPerPixel = tbm.BytesPerPixel
							bbm.Bytewidth = tbm.Bytewidth
							//bbm.ImgBuf = robotgo.TostringBitmap(_bm)
							fmt.Print("\n\n\n\n")
							x1, y1 := robotgo.FindBitmap(_bm)
							fmt.Println(x1, y1)
							robotgo.MoveClick(int(x1/2)+bbm.Width/4, int(y1/2)+bbm.Height/4)

							s1 = false
							break
						}
					}
				}
				fmt.Println("end")
			}()

		}
	}

	//s := robotgo.Start()
	//defer robotgo.End()
	//var ok []int

	//for ev := range s {
	//	fmt.Println(ev)
	//	if ev.Rawcode == 55 {
	//		if ev.Kind == hook.KeyHold {
	//			x, y := robotgo.GetMousePos()
	//			ok = append(ok, x, y)
	//			fmt.Println(ok)
	//			//defer w.Exit()
	//		}
	//
	//		//if ev.Kind == hook.KeyUp {
	//		//	x, y := robotgo.GetMousePos()
	//		//	ok = append(ok, x, y)
	//		//
	//		//	fmt.Println(ok)
	//		//	fmt.Println("ok\n")
	//		//
	//		//	bbm := &Bitmap{}
	//		//
	//		//	_bm := robotgo.CaptureScreen(ok[0], ok[1], ok[2]-ok[0], ok[3]-ok[1])
	//		//	tbm := robotgo.ToBitmap(_bm)
	//		//	bbm.Height = tbm.Height
	//		//	bbm.Width = tbm.Width
	//		//	bbm.BitsPixel = tbm.BitsPixel
	//		//	bbm.BytesPerPixel = tbm.BytesPerPixel
	//		//	bbm.Bytewidth = tbm.Bytewidth
	//		//	bbm.ImgBuf = robotgo.TostringBitmap(_bm)
	//		//
	//		//	x1, y1 := robotgo.FindBitmap(_bm)
	//		//	fmt.Println(x1/2, y1/2, x1, y1)
	//		//	robotgo.MoveClick(x1/2+bbm.Width/4, y1/2+bbm.Height/4)
	//		//	xerror.Panic(ioutil.WriteFile(fmt.Sprintf("%d.png", time.Now().Unix()), xerror.PanicBytes(json.Marshal(bbm)), 0644))
	//		//	ok = ok[0:0]
	//		//}
	//	}
	//
	//	// _name := fmt.Sprintf("%d.png", time.Now().Unix())
	//	// robotgo.SaveCapture(_name, x1, x2, x3-x1, x4-x2)
	//	// robotgo.Sleep(2)
	//	// if xerror.PanicErr(dlgs.Question("dd", "is ok", false)).(bool) {
	//	// 	ok = true
	//	// }
	//	//}
	//}

	//openbit := robotgo.OpenBitmap()
	//xerror.Panic(robotgo.ActiveName("WeChat"))
	// [79 142 130 191]

	//fmt.Println(robotgo.FindPic("test.png"))

	//fmt.Println(robotgo.FindEveryBitmap(robotgo.OpenBitmap("test.png")))
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

	// _img1 := imageutil.Rgb2Gray(imgo.ReadPNG("3.png"))
	// _img2 := imageutil.Rgb2Gray(imgo.ReadPNG("sss.png"))

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

	// x, y := robotgo.FindPic("2.png")
	// fmt.Println(x, y)
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

func main() {
	s := robotgo.Start()
	defer robotgo.End()
	var ok []int
	var cmd int
	for ev := range s {
		if ev.Rawcode == 36 {
			cmd++
		} else {
			cmd = 0
		}

		if cmd == 4 {
			x, y := robotgo.GetMousePos()
			ok = append(ok, x, y)
			fmt.Println(ok)
		}

		if len(ok) == 4 {
			bbm := &Bitmap{}

			_bm := robotgo.CaptureScreen(ok[0], ok[1], ok[2]-ok[0], ok[3]-ok[1])
			tbm := robotgo.ToBitmap(_bm)
			bbm.Height = tbm.Height
			bbm.Width = tbm.Width
			bbm.BitsPixel = tbm.BitsPixel
			bbm.BytesPerPixel = tbm.BytesPerPixel
			bbm.Bytewidth = tbm.Bytewidth

			_name := fmt.Sprintf("%d.png", time.Now().Unix())
			robotgo.SaveBitmap(_bm, _name, )
			//robotgo.SaveCapture(_name, ok[0], ok[1], ok[2]-ok[0], ok[3]-ok[1])
			//bbm.ImgBuf = robotgo.SaveBitmap(_bm, _name)

			x1, y1 := robotgo.FindBitmap(robotgo.OpenBitmap(_name, ))
			//x1, y1 := robotgo.FindPic(_name)
			//x1, y1 := robotgo.FindBitmap(_bm)
			fmt.Println(x1/2, y1/2, x1, y1)
			robotgo.MoveClick(x1/2+bbm.Width/4, y1/2+bbm.Height/4)
			//xerror.Panic(ioutil.WriteFile(_name, xerror.PanicBytes(json.Marshal(bbm)), 0644))
			shutil.Execute(fmt.Sprintf("open %s", _name))
			//fmt.Println(dlgs.Entry("文件名", "图片", _name))
			//fmt.Println(dlgs.File("文件名","",false))

			ok = ok[0:0]
		}
	}
}
