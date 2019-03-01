package main

/*
   #cgo LDFLAGS: -L/usr/lib/ -lwallaby
   #include <wallaby/button.h>
*/
import "C"
import "fmt"

func main() {
        fmt.Println(C.a_button());
}

