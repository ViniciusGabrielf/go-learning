package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Print number of physical processors in current machine
	fmt.Println(runtime.NumCPU())
}
