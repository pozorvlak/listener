package main

import (
        "fmt"
)

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

func summingSubsets(xs []int, sum, n int) [][]int {
        if sum == 0 && n == 0 {
                return [][]int{ make([]int, 0) }
        } else if n == 0 || len(xs) == 0 {
                return make([][]int, 0)
        } else {
                out := summingSubsets(xs[1:], sum, n)
                x := xs[0]
                for _, s := range(summingSubsets(xs[1:], sum - x, n - 1)) {
                        out = append(out,
                                append(s, x))
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
        pslice = pslice[0:i]
        for _, cand := range summingSubsets(pslice, 600, 8) {
                fmt.Println(cand)
        }
}
