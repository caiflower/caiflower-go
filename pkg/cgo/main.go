package main

import (
	"fmt"
	"unsafe"
)

/*
#cgo LDFLAGS: -L./rust/target/release -lcaiflower_rust
#include <stdint.h>
#include <stdlib.h>

void say_hello(const char *name, int age);
*/
import "C"

func main() {
	name := "Alice"
	age := 15

	nameCStr := C.CString(name)
	defer C.free(unsafe.Pointer(nameCStr))

	fmt.Println("\n=== Calling Rust function from Go ===")
	C.say_hello(nameCStr, C.int(age))
	fmt.Println("=== Call completed ===")
}
