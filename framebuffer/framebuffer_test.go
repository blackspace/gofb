package framebuffer

import (
	"testing"
)

func BenchmarkFill(b *testing.B) {

	fb:=NewFramebuffer()

	fb.Init()

	fb.Fill(0,0,255,0)

	fb.Release()

}
