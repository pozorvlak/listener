package main

import (
        "fmt"
)

/* Take integers from a channel until you encounter one that's at least n */
func upto(in chan int, n int) chan int {
        out := make(chan int)
        go func() {
                for m := range in {
                        if m < n {
                                out <- m
                        } else {
                                break
                        }
                }
                close(out)
        }()
        return out
}

/* Take integers from a channel that satisfy the supplied predicate */
func filter(f func(int) bool, in chan int) chan int {
        out := make(chan int)
        go func() {
                for m := range in {
                        if f(m) {
                                out <- m
                        }
                }
                close(out)
        }()
        return out
}

/* Find length-n subslices of xs that sum to n */
func summingSubsets(xs []int, sum, n int) [][]int {
        if sum == 0 && n == 0 {
                return [][]int{ make([]int, 0) } // Only one way: take nothing
        } else if n == 0 || len(xs) == 0 {
                return make([][]int, 0)          // Can't be done!
        } else {
                out := summingSubsets(xs[1:], sum, n)
                x := xs[0]
                for _, s := range(summingSubsets(xs[1:], sum - x, n - 1)) {
                        out = append(out, append(s, x))
                }
                return out
        }
}

func main() {
        ps := filter(func (n int) bool { return n > 9 }, upto(Primes(), 100))
        pslice := make([]int, 30)
        i := 0
        for p := range ps {
                pslice[i] = p
                i++
        }
        pslice = pslice[0:i]    // Trim excess zeroes; gives a huge speedup
        for _, cand := range summingSubsets(pslice, 600, 8) {
                fmt.Println(cand)
        }
}
