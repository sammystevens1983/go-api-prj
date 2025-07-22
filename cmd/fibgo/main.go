package main

import (
	"flag"
	"fmt"
	"os"

	"fibbench/internal/fibonacci"
)

func main() {
    n := flag.Int("n", 0, "which Fibonacci number to compute")
    flag.Parse()

    if *n < 0 {
        fmt.Fprintln(os.Stderr, "n must be non-negative")
        os.Exit(1)
    }

    fmt.Println(fibonacci.Compute(*n))
}
