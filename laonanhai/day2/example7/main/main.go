package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var gopath string = os.Getenv("GOPATH")
	fmt.Printf("The GOPATH is %s\n", gopath)
	fmt.Println("Current system is ", runtime.GOOS)
	fmt.Println("Current system is ", runtime.GOARCH)
}
