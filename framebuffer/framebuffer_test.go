package framebuffer

import (
	"testing"
)

func BenchmarkFill(b *testing.B) {

	fb:=NewFramebuffer()

	fb.Open()

	fb.Fill(0,0,255,0)

	fb.Close()

}
