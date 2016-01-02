package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "errors"
import "strconv"

func FieldsReverse(line string) []string {
	fields := strings.Fields(line)
	reversed := make([]string, len(fields))

	for i, j := len(fields)-1, 0; i >= 0; i, j = i-1, j+1 {
		reversed[j] = fields[i]
	}

	return reversed
}

func findMth(s []string, i int) (string, error) {
	abs := i - 1
	for i, v := range s {
		if i == abs {
			return v, nil
		}
	}
	return "", errors.New("Invalid index")
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
		if len(line) < 2 {
			continue
		}

		rev := FieldsReverse(line)
		index, _ := strconv.Atoi(rev[0])
		if index > len(rev) {
			continue
		}

		mth := rev[1:][index-1]
		fmt.Println(mth)
	}
}
