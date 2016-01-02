package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

func max(i1, i2 int) int {
	if i1 >= i2 {
		return i1
	} else {
		return i2
	}
}

func makeLcsTable(n, m int) [][]int {
	lcsTable := make([][]int, n)
	for i := 0; i < n; i++ {
		lcsTable[i] = make([]int, m)
		for j := 0; j < m; j++ {
			lcsTable[i][j] = 0
		}
	}
	return lcsTable
}

func fillLcsTable(lcsTable [][]int, s1, s2 string, n, m int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				lcsTable[i][j] = lcsTable[i-1][j-1] + 1
			} else {
				lcsTable[i][j] = max(lcsTable[i][j-1], lcsTable[i-1][j])
			}
		}
	}
}

func backtrackLcs(table [][]int, s1, s2 string, i, j int) string {
	if i == 0 || j == 0 {
		return ""
	} else if s1[i-1] == s2[j-1] {
		return backtrackLcs(table, s1, s2, i-1, j-1) + string(s1[i-1])
	} else {
		if table[i][j-1] > table[i-1][j] {
			return backtrackLcs(table, s1, s2, i, j-1)
		} else {
			return backtrackLcs(table, s1, s2, i-1, j)
		}
	}
}

func lcs(s1, s2 string) string {
	n, m := len(s1), len(s2)
	lcsTable := makeLcsTable(n+1, m+1)
	fillLcsTable(lcsTable, s1, s2, n, m)
	return backtrackLcs(lcsTable, s1, s2, n, m)
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
		if len(line) == 0 {
			continue
		}

		ss := strings.Split(line, ";")
		fmt.Println(lcs(ss[0], ss[1]))
	}
}
