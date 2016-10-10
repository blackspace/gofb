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
	Data         []uint32
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

func (f *Framebuffer)Open()  {
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

	f.Data= *(*[]uint32)(unsafe.Pointer(&sl))
}

func (f *Framebuffer)Close() {
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

	f.Data[location/4]=a<<24+r<<16+g<<8+b
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

func  (f *Framebuffer)Fill(r,g,b,a uint32) {
	for y:=0;y<f.Yres;y++ {
		for x :=0; x < f.Xres; x++ {
			f.SetPixel(x,y, r,g,b,a)
		}
	}
}


