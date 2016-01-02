package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strconv"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum := 0
		for _, c := range line {
			digit, _ := strconv.Atoi(string(c))
			sum += digit
		}
		fmt.Println(sum)
	}
}
