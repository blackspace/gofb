package main

import (
	"github.com/ungerik/go-cairo"
	"gofb/framebuffer"
	"fmt"
)


func main() {
	fb := framebuffer.NewFramebuffer()
	fb.Init()
	defer 	fb.Release()
	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, fb.Xres, fb.Yres)


	surface.SetSourceRGBA(0,0,0,1)
	surface.Rectangle(0,0,float64(fb.Xres), float64(fb.Yres))
	surface.Fill()
	surface.SelectFontFace("serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surface.SetFontSize(32.0)
	surface.SetSourceRGBA(1,1,1,1)
	surface.MoveTo(10.0, 50.0)
	surface.ShowText("道可道 非常道")
	data:=surface.GetData()
	fb.DrawData(0,0,data,fb.Xres, fb.Yres)


	surface.Finish()

	fmt.Scanln()


}
