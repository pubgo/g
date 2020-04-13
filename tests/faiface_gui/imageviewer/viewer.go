package main

import (
	"fmt"
	"github.com/faiface/gui"
	"github.com/vcaesar/imgo"
	"image"
	"image/draw"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/bmp"
)

func Viewer(env gui.Env, theme *Theme, view <-chan string) {
	redraw := func(r image.Rectangle, img image.Image) func(draw.Image) image.Rectangle {
		return func(drw draw.Image) image.Rectangle {
			draw.Draw(drw, r, &image.Uniform{theme.Empty}, image.ZP, draw.Src)
			DrawCentered(drw, r, img, draw.Over)
			return r
		}
	}

	var (
		r   image.Rectangle
		img image.Image
	)

	_img, _, _ := imgo.DecodeFile("1.png")
	env.Draw() <- redraw(r, _img)

	for {
		select {
		case path := <-view:
			_img, _, _ := imgo.DecodeFile("1.png")
			env.Draw() <- redraw(r, _img)
			fmt.Println(path)

			//func() {
			//	f, err := os.Open(path)
			//	if err != nil {
			//		img = invalid
			//		return
			//	}
			//	defer f.Close()
			//	img, _, err = image.Decode(f)
			//	if err != nil {
			//		img = invalid
			//		return
			//	}
			//}()
			//env.Draw() <- redraw(r, img)
		case e, ok := <-env.Events():
			if !ok {
				close(env.Draw())
				return
			}
			if resize, ok := e.(gui.Resize); ok {
				r = resize.Rectangle
				env.Draw() <- redraw(r, img)
			}
		}
	}
}
