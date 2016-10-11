package main

import (
	"gofb/framebuffer"
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"os"
	"github.com/ungerik/go-cairo"
)

func main() {

	fb := framebuffer.NewFramebuffer()
	defer 	fb.Release()

	fb.Init()

	fb.Fill(255,255,255,0)

	const S = 1024
	w:=1680
	h:=1050
	dc := gg.NewContext(w,h)
	dc.DrawRectangle(0,0,float64(w),float64(h))
	dc.SetRGB(1, 1, 1)
	dc.Fill()


	f,err:=os.Open("./flower.png")
	if err!=nil {
		panic(err.Error())
	}

	flower,_,err:=image.Decode(f)
	if err!=nil {
		panic(err.Error())
	}

	dc.DrawImage(flower,w-flower.Bounds().Max.X,h-flower.Bounds().Max.Y)

	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}


	fb.DrawImage(0,0,dc.Image())

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, 240, 80)
	surface.SetSourceRGBA(1,1,1,1)
	surface.Rectangle(0,0,240,80)
	surface.Fill()
	surface.SelectFontFace("serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surface.SetFontSize(32.0)
	surface.SetSourceRGBA(0,0,0,.1)
	surface.MoveTo(10.0, 50.0)
	surface.ShowText("Hello World")
	surface.Finish()

	fb.DrawImage(0,0,surface.GetImage())







	fmt.Scanln()


}

