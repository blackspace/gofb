package main

/*
#cgo linux LDFLAGS: -lrt

#include <fcntl.h>
#include <unistd.h>
#include <sys/mman.h>

#define FILE_MODE (S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH)

int my_shm_new(char *name) {
    shm_unlink(name);
    return shm_open(name, O_RDWR|O_CREAT|O_EXCL, FILE_MODE);
}
*/
import "C"
import (
    "fmt"
    "unsafe"
)

const SHM_NAME = "my_shm"
const SHM_SIZE = 4 * 1000 * 1000 * 1000

type MyData struct {
    Col1 int
    Col2 int
    Col3 int
}

func main() {
    fd, err := C.my_shm_new(C.CString(SHM_NAME))
    if err != nil {
        fmt.Println(err)
        return
    }

    C.ftruncate(fd, SHM_SIZE)

    ptr, err := C.mmap(nil, SHM_SIZE, C.PROT_READ|C.PROT_WRITE, C.MAP_SHARED, fd, 0)
    if err != nil {
        fmt.Println(err)
        return
    }
    C.close(fd)

    data := (*MyData)(unsafe.Pointer(ptr))

    data.Col1 = 100
    data.Col2 = 876
    data.Col3 = 8021
}
