package image_test

import (
	"github.com/pubgo/g/xerror"
	"gocv.io/x/gocv"

	"github.com/vcaesar/imgo"

	"testing"
)

func TestEnv(t *testing.T) {
	_img := imgo.ReadPNG("1.png")

	window := gocv.NewWindow("Capture Window")
	defer window.Close()

	_m := xerror.PanicErr(gocv.ImageToMatRGBA(_img)).(gocv.Mat)
	defer _m.Close()

	window.IMShow(_m)
}
