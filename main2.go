package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var wordsNumbers = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	f, err := os.Open("test2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0
	// read each line
	for scanner.Scan() {
		// parse each line now
		left, right := -1, -1
		line := scanner.Text()

		// annoying edge cases bc no positive lookahead
		e1 := regexp.MustCompile(`twone`)
		line = e1.ReplaceAllString(line, "2one")

		e2 := regexp.MustCompile(`threeight`)
		line = e2.ReplaceAllString(line, "3eight")

		e3 := regexp.MustCompile(`fiveight`)
		line = e3.ReplaceAllString(line, "5eight")

		e4 := regexp.MustCompile(`nineight`)
		line = e4.ReplaceAllString(line, "9eight")

		e5 := regexp.MustCompile(`oneight`)
		line = e5.ReplaceAllString(line, "1eight")

		e6 := regexp.MustCompile(`eightwo`)
		line = e6.ReplaceAllString(line, "8two")

		e7 := regexp.MustCompile(`eighthree`)
		line = e7.ReplaceAllString(line, "8three")

		//--------------------
		r := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|\d`)

		matches := r.FindAllString(line, -1)
		if len(matches) == 1 {
			i, err := strconv.Atoi(matches[0])
			if err != nil {
				i = wordsNumbers[matches[0]]
			}
			left = i
			right = left
		} else {
			left, err = strconv.Atoi(matches[0])
			if err != nil {
				left = wordsNumbers[matches[0]]
			}
			right, err = strconv.Atoi(matches[len(matches)-1])

			if err != nil {
				right = wordsNumbers[matches[len(matches)-1]]
			}
		}
		total = total + (left*10 + right)
	}

	fmt.Println(total)
}
