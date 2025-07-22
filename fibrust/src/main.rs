// src/main.rs
use std::env;
use fibrust::fib;

fn main() {
    let n = env::args()
               .nth(1)
               .and_then(|s| s.parse().ok())
               .unwrap_or(10);
    println!("{}", fib(n));
}
