import ctypes
import time
from pathlib import Path

HERE = Path(__file__).parent

# ——— Load Go shared lib (built via `go build -buildmode=c-shared -o fib.so`) ———
go_lib = ctypes.CDLL(str(HERE / "fib.so"))
go_lib.Fib.argtypes = [ctypes.c_int]
go_lib.Fib.restype = ctypes.c_int

# ——— Load Rust shared lib (built via `cargo build --release`) ———
rust_lib = ctypes.CDLL(str(HERE / "libfibrust.so"))
rust_lib.fib.argtypes = [ctypes.c_int]
rust_lib.fib.restype = ctypes.c_int

# ——— Pure‑Python fib for comparison ———


def fib_py(n: int) -> int:
    if n < 2:
        return n
    return fib_py(n-1) + fib_py(n-2)

# ——— Wrappers around your FFI calls ———


def fib_go(n: int) -> int:
    return go_lib.Fib(n)


def fib_rust(n: int) -> int:
    return rust_lib.fib(n)

# ——— Timing helper ———


def time_func(fn, n, repeats=5):
    times = []
    for _ in range(repeats):
        t0 = time.perf_counter()
        res = fn(n)
        t1 = time.perf_counter()
        times.append(t1 - t0)
    return res, sum(times) / repeats


if __name__ == "__main__":
    n = 35
    runs = 5

    r_py,   t_py = time_func(fib_py,   n, repeats=runs)
    r_go,   t_go = time_func(fib_go,   n, repeats=runs)
    r_rust, t_rust = time_func(fib_rust, n, repeats=runs)

    print(f"\nfib({n}) over {runs} runs:")
    print(f" Python result={r_py}, avg={t_py:.6f}s")
    print(f" Go     result={r_go}, avg={t_go:.6f}s")
    print(f" Rust   result={r_rust}, avg={t_rust:.6f}s")
    print(f"\nSpeedups: Go  ≈ {t_py/t_go:.2f}×,  Rust ≈ {t_py/t_rust:.2f}×")
