package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

func RemoveChars(s string, chars string) string {
	f := func(c rune) bool {
		for _, char := range chars {
			if char == c {
				return true
			}
		}
		return false
	}

	output := make([]byte, 0)
	for _, char := range s {
		if !f(char) {
			output = append(output, byte(char))
		}
	}
	return string(output)
}

func getSequences(s string) (string, string) {
	split := strings.Split(s, ",")
	return split[0], strings.Trim(split[1], " ")
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
		word, rem := getSequences(line)
		fmt.Println(RemoveChars(word, rem))
	}
}
