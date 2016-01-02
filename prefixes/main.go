package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"
import "unicode"

type Stack []float32

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Push(item float32) {
	slice := *s
	slice = append(slice, item)
	*s = slice
}

func (s *Stack) Pop() float32 {
	slice := *s
	item, slice := slice[len(slice)-1], slice[:len(slice)-1]
	*s = slice
	return item
}

func isDigit(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func resolve(line string) int {
	fields := strings.Fields(line)
	stack := new(Stack)

	for i := len(fields) - 1; i >= 0; i-- {
		c := fields[i]
		if isDigit(c) {
			num, _ := strconv.Atoi(c)
			stack.Push(float32(num))
		} else {
			var result float32
			op1 := stack.Pop()
			op2 := stack.Pop()
			if c == "+" {
				result = op1 + op2
			} else if c == "*" {
				result = op1 * op2
			} else {
				result = op1 / op2
			}
			stack.Push(result)
		}
	}

	return int(stack.Pop())
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
		if len(line) > 0 {
			res := resolve(line)
			fmt.Println(res)
		}
	}
}
