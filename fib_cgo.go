package main

// #include <stdint.h>
import "C"

//export Fib
func Fib(n C.int) C.int {
    if n < 2 {
        return n
    }
    return Fib(n-1) + Fib(n-2)
}

func main() {}  // required but never called
