package main

import (
	"fmt"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/kbinani/screenshot"
)

func main() {
	fmt.Println("Hello, playground")
	a := app.New()
	w := a.NewWindow("Images")
	bounds := screenshot.GetDisplayBounds(0)
	m := make([]uint32, bounds.Dx()*bounds.Dy()*4)

	for i := 0; i < 100; i++ {
		img, err := screenshot.CaptureRect(bounds)

		if err != nil {
			panic(err)
		}
		img_size := len(img.Pix)
		for i := 0; i < img_size; i += 4 {
			m[i/4] += (uint32(img.Pix[i]) + uint32(img.Pix[i+1]) + uint32(img.Pix[i+2])) / 3
		}
		time.Sleep(100 * time.Millisecond)
	}
	for i := 0; i < len(m); i++ {
		m[i] /= 100
	}

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}

	threshold := uint32(50)
	for i := 0; i < len(img.Pix); i += 4 {
		avg := (uint32(img.Pix[i]) + uint32(img.Pix[i+1]) + uint32(img.Pix[i+2])) / 3
		diff := uint32(math.Abs(float64(avg) - float64(m[i/4])))
		if diff > threshold {
			img.Pix[i] = 0
			img.Pix[i+1] = 0
			img.Pix[i+2] = 0
		}
	}

	img_img := img.SubImage(img.Rect)
	img_canvas := canvas.NewImageFromImage(img_img)
	w.SetContent(img_canvas)
	w.Resize(fyne.NewSize(600, 480))
	w.ShowAndRun()
}
