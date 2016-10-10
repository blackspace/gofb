package framebuffer

/*
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
)

type Framebuffer struct {
	Fd int
	BitsPerPixel int
	Xres int
	Yres int
	Fbp unsafe.Pointer
	Xoffset int
	Yoffset int
	LineLength int
	Screensize int
}

func  NewFramebuffer() *Framebuffer {
	return &Framebuffer{}
}
func (f *Framebuffer)Open()  {
	fd,err:=C.OpenFrameBuffer(C.CString("/dev/fb0"))

	if err!=nil {
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

	f.Fbp = C.mmap(nil, C.size_t(f.Screensize), C.PROT_READ | C.PROT_WRITE, C.MAP_SHARED, fd, 0)

}

func (f *Framebuffer)Close() {
	C.munmap(f.Fbp, C.size_t(f.Screensize))
	C.close(C.int(f.Fd))
}

func  (f *Framebuffer)SetPixel(x int,y int,p Pixel) {
	location := (x + f.Xoffset) *(f.BitsPerPixel / 8) + (y + f.Yoffset) * f.LineLength

	start:=uintptr(f.Fbp) + uintptr(location)

	*(*uint32)(unsafe.Pointer(start)) = p.ToU32()

}


