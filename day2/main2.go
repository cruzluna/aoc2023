package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	gameId := 0 // will function as index
	for scanner.Scan() {
		gameId++
		// fmt.Printf("--------------Game \033[31m%v \033[0m--------\n", gameId)
		line := scanner.Text()

		idx := strings.Index(line, ":")
		line = line[idx+1:]

		// split up the games. ; delimiter
		rounds := strings.Split(line, ";")
		games := make([][]string, len(rounds))

		for i, round := range rounds {
			games[i] = strings.Split(strings.TrimSpace(round), ",")
		}

		// ready to parsed
		// find the maximum of each color per game
		maxR, maxG, maxB := 0, 0, 0
		for _, round := range games {
			for j := range round {
				split := strings.Fields(round[j])

				if split[1] == "blue" {
					cnt, _ := strconv.Atoi(split[0])
					maxB = max(maxB, cnt)
				} else if split[1] == "green" {
					cnt, _ := strconv.Atoi(split[0])
					maxG = max(maxG, cnt)
				} else {
					cnt, _ := strconv.Atoi(split[0])
					maxR = max(maxR, cnt)
				}
			}
		}

		// store the sum of maxes
		total = total + (maxB * maxG * maxR)
	}

	fmt.Println("TOTAL: ", total)
}

// need to update my Go version...
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
