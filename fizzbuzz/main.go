package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

func fb(i int, x int, y int) string {
	ix := (i % x) == 0
	iy := (i % y) == 0

	if ix && iy {
		return "FB"
	} else if ix {
		return "F"
	} else if iy {
		return "B"
	} else {
		return strconv.Itoa(i)
	}
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

		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])
		n, _ := strconv.Atoi(fields[2])

		for i := 1; i <= n; i++ {
			fmt.Print(fb(i, x, y), " ")
		}

		fmt.Print("\n")
	}
}
