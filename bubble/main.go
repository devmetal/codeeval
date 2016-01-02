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

func numbers(line string) []int {
	fields := strings.Split(line, " ")
	vector := make([]int, len(fields))
	for index, field := range fields {
		vector[index], _ = strconv.Atoi(field)
	}
	return vector
}

func iteration(vector []int) bool {
	n := len(vector)
	swapped := false
	for i := 1; i < n; i++ {
		if vector[i-1] > vector[i] {
			vector[i-1], vector[i] = vector[i], vector[i-1]
			swapped = true
		}
	}
	return swapped
}

func bubble(vector []int, it int) {
	if it > 0 {
		for i := 0; i < it && iteration(vector); {
			i++
		}
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
		parts := strings.Split(line, "|")
		parts[0] = strings.Trim(parts[0], " ")
		parts[1] = strings.Trim(parts[1], " ")

		vector := numbers(parts[0])
		it, _ := strconv.Atoi(parts[1])

		bubble(vector, it)

		for _, item := range vector {
			fmt.Print(item, " ")
		}

		fmt.Print("\n")
	}
}
