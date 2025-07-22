// src/lib.rs
/// A naïve recursive Fibonacci, exposed with C ABI.
#[no_mangle]
pub extern "C" fn fib(n: i32) -> i32 {
    if n < 2 {
        n
    } else {
        // two recursive calls
        fib(n - 1) + fib(n - 2)
    }
}


// inside src/lib.rs

#[cfg(test)]
mod tests {
    use super::fib;

    #[test]
    fn smoke() {
        // basic sanity checks
        assert_eq!(fib(0), 0);
        assert_eq!(fib(1), 1);
        assert_eq!(fib(5), 5);
        assert_eq!(fib(10), 55);
    }
}
