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

	// read each line
	gameId := 0

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameId++

		idx := strings.Index(line, ":")
		line = line[idx+1:]

		// split up the games. ; delimiter
		rounds := strings.Split(line, ";")
		games := make([][]string, len(rounds))

		for i, round := range rounds {
			games[i] = strings.Split(strings.TrimSpace(round), ",")
		}

		gameFlag := true
		for _, round := range games {
			// now iterate the round
			r, g, b := 12, 13, 14
			for j := range round {
				// split it based on white space
				split := strings.Fields(round[j])
				if split[1] == "blue" {
					cnt, _ := strconv.Atoi(split[0])
					b = b - cnt
				} else if split[1] == "green" {
					cnt, _ := strconv.Atoi(split[0])
					g = g - cnt
				} else {
					cnt, _ := strconv.Atoi(split[0])
					r = r - cnt
				}
				if r < 0 || g < 0 || b < 0 {
					break
				}
			}
			if r < 0 || g < 0 || b < 0 {
				gameFlag = false
				break
			}
		}
		fmt.Printf("Total {%v} , gameId {%v}", total, gameId)

		if gameFlag {
			total = total + gameId
		}

		// iterate each game and check

	}
	fmt.Println("TOTAL: ", total)
}
