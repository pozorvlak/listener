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

func main() {
        ps := upto(Primes(), 20)
        for p := range ps {
                fmt.Printf("%d  ", p)
        }
        fmt.Println()
}
