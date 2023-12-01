package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0
	// read each line
	for scanner.Scan() {
		i, err := strconv.Atoi(twoPtr(scanner.Text()))
		if err != nil {
			log.Fatalln("Msg: ", scanner.Text(), err)
		}
		total += i

	}

	fmt.Println(total)
}

func twoPtr(str string) string {
	lDig, rDig := "", ""
	l, r := 0, len(str)-1

	for l < r {

		if lDig == "" && unicode.IsDigit(rune(str[l])) {
			lDig = string(str[l])
		}

		if rDig == "" && unicode.IsDigit(rune(str[r])) {
			rDig = string(str[r])
		}

		if rDig != "" && lDig != "" {
			break
		}

		if lDig == "" {
			l++
			// find left first
			continue
		}

		if rDig == "" {
			r--
		}
	}

	if len(lDig)+len(rDig) == 1 {
		return (lDig + rDig) + (lDig + rDig)
	}
	return lDig + rDig
}
