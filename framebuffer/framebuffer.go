package framebuffer

/*
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <linux/fb.h>
#include <sys/mman.h>
#include <sys/ioctl.h>

int OpenFrameBuffer(char * name) {
	return open(name, O_RDWR);
}

static  int GetFixedScreenInformation(int fd,struct fb_fix_screeninfo *finfo)
{

     return ioctl(fd, FBIOGET_FSCREENINFO, finfo);
}


static  int GetVarScreenInformation(int fd,struct fb_var_screeninfo *vinfo) {
	return ioctl(fd, FBIOGET_VSCREENINFO, vinfo);
}
 */
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
	"image"
)

type Framebuffer struct {
	Fd           int
	BitsPerPixel int
	Xres         int
	Yres         int
	Data         []byte
	Xoffset      int
	Yoffset      int
	LineLength   int
	Screensize   int
}

func  NewFramebuffer() *Framebuffer {
	return &Framebuffer{}
}

func _HideCursor() {
	fmt.Print("\033[?25l")
}

func _ShowCursor() {
	fmt.Printf("\033[?25h")
}

func (f *Framebuffer)Init()  {
	_HideCursor()
	dev_file:=C.CString("/dev/fb0")
	fd,err:=C.OpenFrameBuffer(dev_file)
	C.free(unsafe.Pointer(dev_file))

	if err!= nil {
		panic(errors.New("Open the framebuffer failed"))
	}

	var finfo C.struct_fb_fix_screeninfo
	if _, err := C.GetFixedScreenInformation(fd, &finfo); err != nil {
		fmt.Println(err)
	}

	var vinfo C.struct_fb_var_screeninfo
	if _, err := C.GetVarScreenInformation(fd, &vinfo); err != nil {
		fmt.Println(err)
	}

	f.Xres=int(vinfo.xres)
	f.Yres=int(vinfo.yres)
	f.BitsPerPixel=int(vinfo.bits_per_pixel)
	f.Xoffset=int(vinfo.xoffset)
	f.Yoffset=int(vinfo.yoffset)
	f.LineLength=int(finfo.line_length)

	f.Screensize=int(finfo.smem_len)

	addr:= uintptr(C.mmap(nil, C.size_t(f.Screensize), C.PROT_READ | C.PROT_WRITE, C.MAP_SHARED, fd, 0))


	var sl = struct {
		addr uintptr
		len  int
		cap  int
	}{addr, f.Screensize, f.Screensize}

	f.Data= *(*[]byte)(unsafe.Pointer(&sl))


}

func (f *Framebuffer)Release() {
	C.munmap(unsafe.Pointer(&f.Data[0]), C.size_t(f.Screensize))
	C.close(C.int(f.Fd))
	_ShowCursor()
}

func  (f *Framebuffer)SetPixel(x int,y int,r uint32,g uint32,b uint32,a uint32) {
	if x<0 || x>f.Xres {
		panic(errors.New("X is too big or is negative"))
	}

	if y<0 || y>f.Yres {
		panic(errors.New("Y is too big or is negative"))
	}

	location := (x + f.Xoffset) *(f.BitsPerPixel / 8) + (y + f.Yoffset) * f.LineLength

	f.Data[location+3]=byte(a&0xff)
	f.Data[location+2]=byte(r&0xff)
	f.Data[location+1]=byte(g&0xff)
	f.Data[location]=byte(b&0xff)
}

func  (f *Framebuffer)DrawImage(xoffset int,yoffset int,image image.Image) {

	b:=image.Bounds()

	for y:=0;y<b.Max.Y;y++ {
		for x:=0;x<b.Max.X;x++ {
			r,g,b,a:=image.At(x,y).RGBA()
			f.SetPixel(x+xoffset,y+yoffset,r&0xff,g&0xff,b&0xff,a&0xff)
		}
	}
}

func  (f *Framebuffer)DrawData(xoffset int,yoffset int,data []byte,w int,h int) {

	if w>f.Xres {
		panic(errors.New("The width of data must NOT be bigger the Xres of the framebuffer"))
	}

	for y:=0;y<h;y++ {
		if (y+1)*w>len(data) {
			panic(errors.New("The length of image data is too small or w(h) argument is wrong"))
		}


		line_start := (xoffset + f.Xoffset) *(f.BitsPerPixel / 8) + (y + yoffset + f.Yoffset) * f.LineLength
		line_end := (xoffset + f.Xoffset) *(f.BitsPerPixel / 8) + (y +1+ yoffset + f.Yoffset) * f.LineLength-1

		if line_start +w>line_end {
			panic(errors.New("The lines is too long beyond the framebuffer"))
		}

		data_line :=data[y*w*4:(y+1)*w*4]
		copy(f.Data[line_start:line_start +w*4], data_line)
	}

}

func  (f *Framebuffer)Fill(r,g,b,a uint32) {
	for y:=0;y<f.Yres;y++ {
		for x :=0; x < f.Xres; x++ {
			f.SetPixel(x,y, r,g,b,a)
		}
	}
}




