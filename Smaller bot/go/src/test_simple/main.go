package main

/*
  #cgo CFLAGS: -g -Wall
  #include "test.h"
*/
import "C"
import "fmt"

func main() {
    fmt.Println(C.Foo());
}