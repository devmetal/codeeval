package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NumPair struct {
	x int
	n int
}

func readLines(scanner *bufio.Scanner) <-chan string {
	ch := make(chan string)

	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()

	return ch
}

func readPairs(in <-chan string) <-chan *NumPair {
	ch := make(chan *NumPair)

	go func() {
		for i := range in {
			fields := strings.Split(i, ",")
			x, _ := strconv.Atoi(fields[0])
			n, _ := strconv.Atoi(fields[1])
			ch <- &NumPair{x, n}
		}
		close(ch)
	}()

	return ch
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := readLines(scanner)
	pairs := readPairs(lines)

	for pair := range pairs {
		m := pair.n
		i := 1

		for m < pair.x {
			i++
			m = pair.n * i
		}

		fmt.Println(m)
	}
}
