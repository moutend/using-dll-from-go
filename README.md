# Using DLL from Go

In this example, `GetModuleFileNameW` function is called from Go code to obtain the path of running process.

> GetModuleFileName function (Windows)
> 
> Retrieves the fully qualified path for the file that contains the specified module. The module must have been loaded by the current process.
> https://msdn.microsoft.com/en-us/library/windows/desktop/ms683197(v=vs.85).aspx

If you are not familiar with Windows API, you might confuse that what the difference between `GetModuleFileName` function that is documented above and `GetModuleFileNameW`.
The suffix `W` means that the function receive string as UTF16 and `A` means that the function receives string as ASCII.
`GetModuleFileName` function has two version, and the document above shows general description for two functions.
Actually, if you did `strings kernel32.dll` command, you will find two functions.

```shell
$ strings kernel32.dll | grep "^GetModuleFileName.$" | sort | uniq
GetModuleFileNameA
GetModuleFileNameW
```

## Build an executable for Windows from another OS

For example, on Mac OSX, run the following command to build an executable for windows:

```shell
$ go generate generate.go
```

Now `syscall_windows.go` is generated and you can build an executable for Windows.

```shell
$ export GOOS=windows    # It assumes GOARCH is the same between Mac OSX and Windows.
$ go build
$ file using-dll-from-go.exe
using-dll-from-go.exe: PE32+ executable for MS Windows (console) Mono/.Net assembly
```

Copy an executable to Windows machine and try this:

```shell
C:\>using-dll-from-go.exe
[C:\using-dll-from-go.exe]
```

## LICENSE

MIT
