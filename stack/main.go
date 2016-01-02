package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strconv"
import "strings"

type Stack []int

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Push(item int) {
	slice := *s
	slice = append(slice, item)
	*s = slice
}

func (s *Stack) Pop() int {
	slice := *s
	item, slice := slice[len(slice)-1], slice[:len(slice)-1]
	*s = slice
	return item
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
		fields := strings.Fields(line)
		stack := new(Stack)

		for _, field := range fields {
			num, _ := strconv.Atoi(field)
			stack.Push(num)
		}

		for i := 0; stack.Len() > 0; i++ {
			num := stack.Pop()
			if i%2 == 0 {
				fmt.Print(num, " ")
			}
		}

		fmt.Print("\n")
	}
}
