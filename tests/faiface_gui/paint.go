package main

import (
	"github.com/pubgo/g/xerror"
	"github.com/vcaesar/imgo"
	"image"
	"image/draw"

	"github.com/faiface/gui/win"
	"github.com/faiface/mainthread"
)



func DrawCentered(dst draw.Image, r image.Rectangle, src image.Image, op draw.Op) {
	if src == nil {
		return
	}
	bounds := src.Bounds()
	center := bounds.Min.Add(bounds.Max).Div(2)
	target := r.Min.Add(r.Max).Div(2)
	delta := target.Sub(center)
	draw.Draw(dst, bounds.Add(delta).Intersect(r), src, bounds.Min, op)
}


func run2() {
	w, err := win.New(win.Title("faiface/win"), win.Size(800, 600), win.Resizable())
	xerror.Panic(err)

	w.Draw() <- func(drw draw.Image) image.Rectangle {
		r := image.Rect(200, 200, 600, 400)
		//draw.Draw(drw, r, image.White, image.ZP, draw.Src)

		_img, _, _ := imgo.DecodeFile("1.png")
		DrawCentered(drw, r, _img, draw.Over)
		return r
	}

	for event := range w.Events() {
		switch event.(type) {
		case win.WiClose:
			close(w.Draw())
		}
	}
}

func main() {
	mainthread.Run(run2)
}
