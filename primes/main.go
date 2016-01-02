package main

import "fmt"
import "sync"

func genPrimes(n int) <-chan int {
	primes := make(chan int)

	go func() {
		primes <- 2

		i := 3
		var c int

		count := 2
		for count <= n {
			for c = 2; c <= i-1; c++ {
				if i%c == 0 {
					break
				}
			}

			if c == i {
				primes <- i
				count++
			}

			i++
		}
		close(primes)
	}()

	return primes
}

func sumWorker(in <-chan int) <-chan int {
	res := make(chan int)
	go func() {
		sum := 0
		for i := range in {
			sum += i
		}
		res <- sum
	}()
	return res
}

func sumOfSums(workers ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	res := make(chan int)
	sum := 0

	readWorker := func(worker <-chan int) {
		workerRes := <-worker
		fmt.Println(workerRes)
		sum += workerRes
		wg.Done()
	}

	wg.Add(len(workers))

	for _, w := range workers {
		go readWorker(w)
	}

	go func() {
		wg.Wait()
		res <- sum
		close(res)
	}()

	return res
}

func main() {
	primes := genPrimes(10000)

	w1 := sumWorker(primes)
	w2 := sumWorker(primes)
	w3 := sumWorker(primes)
	w4 := sumWorker(primes)

	fmt.Println(<-sumOfSums(w1, w2, w3, w4))
}
