package main

import (
	"gofb/framebuffer"
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"os"
)

func main() {

	fb := framebuffer.NewFramebuffer()
	defer 	fb.Release()

	fb.Init()

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


	fmt.Scanln()
}

