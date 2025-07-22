package fibonacci

import "testing"

func TestCompute(t *testing.T) {
    cases := []struct{
        n, want int
    }{
        {0, 0}, {1, 1}, {5, 5}, {10, 55},
    }
    for _, c := range cases {
        if got := Compute(c.n); got != c.want {
            t.Errorf("Compute(%d) = %d; want %d", c.n, got, c.want)
        }
    }
}
