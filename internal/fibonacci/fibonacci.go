package fibonacci

// Compute returns the n-th Fibonacci number (naïve recursive).
func Compute(n int) int {
    if n < 2 {
        return n
    }
    return Compute(n-1) + Compute(n-2)
}
