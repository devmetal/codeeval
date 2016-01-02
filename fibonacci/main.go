package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strconv"

func lines(scanner *bufio.Scanner) <-chan string {
	ch := make(chan string)

	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()

	return ch
}

func fib(n int) int {
	if n <= 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for line := range lines(scanner) {
		n, _ := strconv.Atoi(line)
		fmt.Println(fib(n))
	}
}
