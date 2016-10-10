package main

import (
	"gofb/framebuffer"
	"fmt"
	"github.com/fogleman/gg"
)

func main() {
	fb := framebuffer.NewFramebuffer()
	defer 	fb.Close()

	fb.Open()

	fb.Fill(0,0,0,0)

	dc := gg.NewContext(1000, 1000)

	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()

	fb.DrawImage(0,0,dc.Image())


	fmt.Scanln()
}

