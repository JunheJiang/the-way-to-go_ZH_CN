//package main
//
//func main() {
//	println("Hello", "world")
//}

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("%s", runtime.Version())
}
