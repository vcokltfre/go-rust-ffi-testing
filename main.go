package main

/*
#cgo LDFLAGS: ./lib/librustffi.a -ldl
#include "./lib/librustffi.h"
*/
import "C"

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command: <pass|time>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "pass":
		fmt.Println("Calling rustffi_pass")
		res := C.basic_string_pass(C.CString("testing"))
		fmt.Println("Result:", C.GoString(res))
	case "time":
		fmt.Println("Calling rustffi_time")
		go func() {
			for i := 0; i < 5; i++ {
				fmt.Printf("%d of %d (go)\n", i, 5)
				time.Sleep(time.Second)
			}
		}()
		C.long_running(C.int(5))
	default:
		fmt.Println("Unknown command")
	}
}
