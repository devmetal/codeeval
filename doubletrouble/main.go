package main

import "fmt"
import "log"
import "bufio"
import "os"

func getVariants(line string) int {
	n := len(line)
	m := n / 2
	v := 0

	for i, j := 0, m; i < m && j < n; i, j = i+1, j+1 {
		ci := line[i]
		cj := line[j]
		if ci == '*' || cj == '*' {
			if ci == cj {
				if v == 0 {
					v = 2
				} else {
					v = v * 2
				}
			} else {
				if v == 0 {
					v = 1
				}
			}
		} else {
			if ci != cj {
				return 0
			}
		}
	}

	return v
}

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

func variants(lines <-chan string) {
	for line := range lines {
		fmt.Println(getVariants(line))
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linesCh := lines(scanner)

	variants(linesCh)
}
