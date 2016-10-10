package main

import (
	"gofb/framebuffer"
	"fmt"
	"github.com/ungerik/go-cairo"
)

func main() {
	fb := framebuffer.NewFramebuffer()
	defer 	fb.Close()

	fb.Open()

	fb.Fill(0,0,0,0)

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, 240, 80)
	surface.SelectFontFace("serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surface.SetFontSize(32.0)
	surface.SetSourceRGB(0.0, 0.0, 1.0)
	surface.MoveTo(10.0, 50.0)
	surface.ShowText("Hello World")

	fb.DrawImage(0,0,surface.GetImage())

	surface.Finish()


	fmt.Scanln()
}

