package primes

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

func Primes() chan int {
        ch := make(chan int)
        out := make(chan int)
        go generate(ch)
        go func() {
                for {
                        p := <-ch
                        out <- p
                        ch1 := make(chan int)
                        go filter(ch, ch1, p)
                        ch = ch1
                }
        }()
        return out
}
