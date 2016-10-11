package main

import (
	"gofb/framebuffer"
	"fmt"
)

func main() {

	fb := framebuffer.NewFramebuffer()
	defer 	fb.Release()

	fb.Init()
	fb.Fill(255,255,255,0)

	fmt.Scanln()
}

