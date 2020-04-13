package tests

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"testing"
)

func TestFyne(t *testing.T) {
	app := app.New()
	w := app.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))
	w.ShowAndRun()
}
