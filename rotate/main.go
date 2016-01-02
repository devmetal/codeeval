package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

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

func isRotated(left, right string) bool {
	if len(left) != len(right) {
		return false
	}

	rotateFrom := func(s string, from int) string {
		rotated := ""
		for i, j := 0, from; i < len(s); i, j = i+1, j+1 {
			if j >= len(s) {
				j = 0
			}
			rotated += string(s[j])
		}
		return rotated
	}

	n := len(left)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if left[i] == right[j] {
				rotated := rotateFrom(right, j)
				if left == rotated {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for line := range lines(scanner) {
		words := strings.Split(line, ",")
		left := words[0]
		right := words[1]
		if isRotated(left, right) {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}
	}
}
