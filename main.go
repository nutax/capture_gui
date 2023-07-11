package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/kbinani/screenshot"
)

func main() {
	a := app.New()
	w := a.NewWindow("Images")
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	img_img := img.SubImage(img.Rect)
	img_canvas := canvas.NewImageFromImage(img_img)
	w.SetContent(img_canvas)
	w.Resize(fyne.NewSize(600, 480))

	w.ShowAndRun()
}
