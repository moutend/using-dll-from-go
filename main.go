package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	const nSize = 100
	var buf [nSize]byte
	var hModule syscall.Handle

	// Get running process name
	size := GetModuleFileNameW(hModule, unsafe.Pointer(&buf), nSize)
	fmt.Printf("[%s]\n", string(buf[0:size]))
}
