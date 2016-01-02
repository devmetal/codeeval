package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
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

func pairs(seq []int, x int) <-chan []int {
	ch := make(chan []int)

	go func() {
		found := false
		for i := 0; i < len(seq)-1; i++ {
			for j := i + 1; j < len(seq); j++ {
				if seq[i]+seq[j] == x {
					pair := make([]int, 2)
					pair[0] = seq[i]
					pair[1] = seq[j]
					ch <- pair
					found = true
				}
			}
		}

		if !found {
			ch <- nil
		}

		close(ch)
	}()

	return ch
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linesCh := lines(scanner)

	for line := range linesCh {
		parts := strings.Split(line, ";")
		seq := strings.Split(parts[0], ",")
		x := parts[1]

		iseq := make([]int, len(seq))
		for i, s := range seq {
			iseq[i], _ = strconv.Atoi(s)
		}

		ix, _ := strconv.Atoi(x)

		first := true
		for pair := range pairs(iseq, ix) {
			if len(pair) < 2 {
				fmt.Printf("NULL")
			} else if first {
				first = false
				fmt.Print(pair[0], ",", pair[1])
			} else {
				fmt.Print(";", pair[0], ",", pair[1])
			}
		}

		fmt.Printf("\n")
	}
}
