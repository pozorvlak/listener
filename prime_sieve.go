package prime_sieve

func generate(ch chan<- int) {
        for i := 2; ; i++ {
                ch <- i
        }
}

func filter(in <-chan int, out chan<- int, prime int) {
        for {
                i := <-in
                if i % prime != 0 {
                        out <- i
                }
        }
}

func PrimesUpTo(n int) {
        ch := make(chan int)
        go generate(ch)
        var p int
        for p < n {
                p := <-ch
                ch1 := make(chan int)
                go filter(ch, ch1, p)
                ch = ch1
        }
}
