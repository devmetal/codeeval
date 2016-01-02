package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

func reverse(line string) <-chan string {
	ch := make(chan string)

	go func() {
		fields := strings.Fields(line)
		for i := len(fields) - 1; i >= 0; i-- {
			ch <- fields[i]
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

	for scanner.Scan() {
		line := scanner.Text()
		for word := range reverse(line) {
			fmt.Print(word, " ")
		}
		fmt.Print("\n")
	}
}
