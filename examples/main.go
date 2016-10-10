package main

import (
	"gofb/framebuffer"
	"time"
	"fmt"
)

func main() {
	fb := framebuffer.NewFramebuffer()

	fb.Open()

	for y:=0;y<fb.Yres;y++ {
		for x :=0; x < fb.Xres; x++ {
			fb.SetPixel(x,y, framebuffer.Pixel{0, 255,255,255})
		}
	}

	time.Sleep(3*time.Second)

	for y:=0;y<fb.Yres;y++ {
		for x :=0; x < fb.Xres; x++ {
			fb.SetPixel(x,y, framebuffer.Pixel{0, 0,255,255})
		}
	}

	time.Sleep(3*time.Second)

	for y:=0;y<fb.Yres;y++ {
		for x :=0; x < fb.Xres; x++ {
			fb.SetPixel(x,y, framebuffer.Pixel{0, 255,0,255})
		}
	}

	time.Sleep(3*time.Second)

	for y:=0;y<fb.Yres;y++ {
		for x :=0; x < fb.Xres; x++ {
			fb.SetPixel(x,y, framebuffer.Pixel{0, 255,255,0})
		}
	}

	time.Sleep(3*time.Second)

	fmt.Scanln()

	fb.Close()
}

