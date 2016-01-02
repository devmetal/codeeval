package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

func parseTestCase(line string) (int, int, int, int) {
    parts := strings.Split(line, "|")
    size := strings.Split(strings.Trim(parts[0], " "), "x")
    girl := strings.Split(strings.Trim(parts[1], " "), " ")
    n, _ := strconv.Atoi(size[0])
    m, _ := strconv.Atoi(size[1])
    x, _ := strconv.Atoi(girl[0])
    y, _ := strconv.Atoi(girl[1])
    
    return n, m, x, y
}

func nuts(n, m, x, y) int {
    
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
        n, m, x, y := parseTestCase(line)
        
        
        
        //do magix here
        /*
        Ha nincs a szélén, akkor (n - 1 + m - 1) * 2
        Mivel egy teljes kört kell bejárnia
        
        Tehát ha az n*m es mátrix szélén van akkor kiszámoljuk
        az eltolást. Egyébkét a diókhoz hozzáadjuk a fenti "képletet"
        Az új mátrix pedig (n - 1),(m - 1) es
        */
    }   
}