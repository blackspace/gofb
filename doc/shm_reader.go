package main

/*
#cgo linux LDFLAGS: -lrt

#include <fcntl.h>
#include <unistd.h>
#include <sys/mman.h>

#define FILE_MODE (S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH)

int my_shm_open(char *name) {
    return shm_open(name, O_RDWR, FILE_MODE);
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
    fd, err := C.my_shm_open(C.CString(SHM_NAME))
    if err != nil {
        fmt.Println(err)
        return
    }

    ptr, err := C.mmap(nil, SHM_SIZE, C.PROT_READ|C.PROT_WRITE, C.MAP_SHARED, fd, 0)
    if err != nil {
        fmt.Println(err)
        return
    }
    C.close(fd)

    data := (*MyData)(unsafe.Pointer(ptr))

    fmt.Println(data)
}
