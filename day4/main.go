package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines := strings.Fields(line)
		lines = lines[2:]

		fmt.Println(lines)
		winNums := make(map[string]bool, 0)

		idx := 0
		for i := range lines {
			if lines[i] == "|" {
				idx = i
				break
			}
			winNums[lines[i]] = true
		}

		firstMatch := true
		k := 1
		for i := idx; i < len(lines); i++ {
			if winNums[lines[i]] {
				if firstMatch {
					total++
					firstMatch = false
				} else {
					total = total + k
					k = k * 2
				}
			}
		}

		fmt.Println("TOTAL: ", total)
	}
}
