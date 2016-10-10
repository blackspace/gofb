package main

/*
struct foo {
	int a;
};
 */
import "C"
import "fmt"

func main() {
	var f C.struct_foo

	f.a=1

	fmt.Println(f)
}