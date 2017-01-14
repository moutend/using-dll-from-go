package main

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -output syscall_windows.go generate.go
//sys GetModuleFileNameW(hModule syscall.Handle, buf unsafe.Pointer, nSize uint32) (dWord uint32) = kernel32.GetModuleFileNameA
