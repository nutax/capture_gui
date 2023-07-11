package main

import (
	"fmt"

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
	// Print the image Pix length
	fmt.Println(len(img.Pix))
	fmt.Println(img.Bounds())
	fmt.Println(img.Rect.Max.X * img.Rect.Max.Y * 4)
	// Print stride
	fmt.Println(img.Stride)
	fmt.Println(img.Rect.Max.X * 4)
	// Iterate over the image Pix and change RGB values to the average of the three
	for i := 0; i < len(img.Pix); i += 4 {
		avg := (img.Pix[i] + img.Pix[i+1] + img.Pix[i+2]) / 3
		img.Pix[i] = avg
		img.Pix[i+1] = avg
		img.Pix[i+2] = avg
		img.Pix[i+3] = 255
	}
	if err != nil {
		panic(err)
	}
	img_img := img.SubImage(img.Rect)
	img_canvas := canvas.NewImageFromImage(img_img)
	w.SetContent(img_canvas)
	w.Resize(fyne.NewSize(600, 480))
	w.ShowAndRun()
}
