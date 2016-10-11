package framebuffer

import (
	"testing"
	"github.com/ungerik/go-cairo"
)

func BenchmarkFill(b *testing.B) {

	fb:=NewFramebuffer()

	fb.Init()

	for i:=0;i<b.N;i++ {
		fb.Fill(0, 0, 255, 0)
	}

	fb.Release()
}

func BenchmarkMakeImage(b *testing.B) {
	for i:=0;i<b.N;i++ {
		surface := cairo.NewSurface(cairo.FORMAT_ARGB32, 1680, 1050)
		surface.SetSourceRGBA(1, 1, 1, 1)
		surface.Rectangle(0, 0, 240, 80)
		surface.Fill()
		surface.SelectFontFace("serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
		surface.SetFontSize(32.0)
		surface.SetSourceRGBA(0, 0, 0, .1)
		surface.MoveTo(10.0, 50.0)
		surface.ShowText("Hello World")
		surface.GetImage()
		surface.Finish()
	}

}

func BenchmarkDrawImage(b *testing.B) {
	fb:=NewFramebuffer()
	fb.Init()

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, fb.Xres, fb.Yres)
	surface.SetSourceRGBA(1,1,1,1)
	surface.Rectangle(0,0,240,80)
	surface.Fill()
	surface.SelectFontFace("serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surface.SetFontSize(32.0)
	surface.SetSourceRGBA(0,0,0,.1)
	surface.MoveTo(10.0, 50.0)
	surface.ShowText("Hello World")
	image:=surface.GetImage()
	surface.Finish()

	for i:=0;i<b.N;i++ {
		fb.DrawImage(0,0,image)
	}

	fb.Release()
}

