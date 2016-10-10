package main

import (
	"gofb/framebuffer"
	"time"
	"fmt"
)

func main() {
	fb := framebuffer.NewFramebuffer()

	fb.Open()


	a1:=time.Now()

	for y:=0;y<fb.Yres;y++ {
		for x :=0; x < fb.Xres; x++ {
			fb.SetPixel(x,y, framebuffer.Pixel{0, 255,255,255})
		}
	}

	d1:=time.Now().Sub(a1)



	for y:=0;y<fb.Yres;y++ {
		for x :=0; x < fb.Xres; x++ {
			fb.SetPixel(x,y, framebuffer.Pixel{0, 0,255,255})
		}
	}

	for y:=0;y<fb.Yres;y++ {
		for x :=0; x < fb.Xres; x++ {
			fb.SetPixel(x,y, framebuffer.Pixel{0, 255,0,255})
		}
	}


	for y:=0;y<fb.Yres;y++ {
		for x :=0; x < fb.Xres; x++ {
			fb.SetPixel(x,y, framebuffer.Pixel{0, 255,255,0})
		}
	}


	fmt.Println(d1)

	fb.Close()
}

