package framebuffer

import (
	"testing"
	"github.com/ungerik/go-cairo"
	"image"
)

func BenchmarkFill(b *testing.B) {
	fb:=NewFramebuffer()


	fb.Init()
	defer fb.Release()

	for i:=0;i<b.N;i++ {
		fb.Fill(0, 0, 255, 0)
	}


}

func BenchmarkMakeSurfaceImage(b *testing.B) {
	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, 1680, 1050)
	defer surface.Finish()
	for i:=0;i<b.N;i++ {
		surface.SetSourceRGBA(1, 1, 1, 1)
		surface.Rectangle(0, 0, 240, 80)
		surface.Fill()
		surface.SelectFontFace("serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
		surface.SetFontSize(32.0)
		surface.SetSourceRGBA(0, 0, 0, .1)
		surface.MoveTo(10.0, 50.0)
		surface.ShowText("Hello World")
		surface.GetImage()
	}

}

func BenchmarkDrawImage(b *testing.B) {
	fb:=NewFramebuffer()

	fb.Init()
	defer fb.Release()

	image:=image.NewRGBA(image.Rect(0,0,fb.Xres,fb.Yres))

	for i:=0;i<b.N;i++ {
		fb.DrawImage(0,0,image)
	}


}

func BenchmarkCairoGetData(b *testing.B) {
	fb:=NewFramebuffer()

	fb.Init()
	defer fb.Release()

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, fb.Xres, fb.Yres)
	defer surface.Finish()

	for i:=0;i<b.N;i++ {
		surface.GetData()
	}

}

func BenchmarkDrawData(b *testing.B) {
	fb:=NewFramebuffer()

	fb.Init()
	defer fb.Release()

	data:=make([]byte,fb.Xres*fb.Yres*4,fb.Xres*fb.Yres*4)

	for i:=0;i<b.N;i++ {
		fb.DrawData(0,0,data,fb.Xres,fb.Yres)
	}
}

func BenchmarkCopyDataByLine(b *testing.B) {
	fb:=NewFramebuffer()

	fb.Init()
	defer fb.Release()

	data1:=make([]byte,fb.Xres*fb.Yres*4,fb.Xres*fb.Yres*4)
	data2:=make([]byte,fb.Xres*fb.Yres*4,fb.Xres*fb.Yres*4)


	for i:=0;i<b.N;i++ {
		for y:=0;y<fb.Yres;y++ {
			copy(data1[y*fb.Xres:(y+1)*fb.Xres],data2[y*fb.Xres:(y+1)*fb.Xres])
		}
	}
}

func BenchmarkCopyDataAllInOne(b *testing.B) {
	fb:=NewFramebuffer()

	fb.Init()
	defer fb.Release()

	data1:=make([]byte,fb.Xres*fb.Yres*4,fb.Xres*fb.Yres*4)
	data2:=make([]byte,fb.Xres*fb.Yres*4,fb.Xres*fb.Yres*4)

	for i:=0;i<b.N;i++ {
		copy(data1, data2)
	}
}

func BenchmarkAnimation(b *testing.B) {
	fb:=NewFramebuffer()

	fb.Init()
	defer fb.Release()

	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, fb.Xres, fb.Yres)
	defer surface.Finish()

	for i:=0;i<b.N;i++ {
		surface.SetSourceRGBA(0,0,0,1)
		surface.Rectangle(0,0,float64(fb.Xres), float64(fb.Yres))
		surface.Fill()
		surface.SelectFontFace("serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
		surface.SetFontSize(32.0)
		surface.SetSourceRGBA(1,1,1,1)
		surface.MoveTo(10.0+float64(i), 50.0)
		surface.ShowText("道可道 非常道")
		data:=surface.GetData()
		fb.DrawData(0,0,data,fb.Xres, fb.Yres)
	}
}
