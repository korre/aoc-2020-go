package two

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func RunFirst() {
	start := time.Now()

	data := util.ReadFile("/two/input.txt")

	inputArr := strings.Split(string(data), "\n")

	totalValid := 0
	for _, s := range inputArr {
		min, _ := strconv.Atoi(s[0:strings.Index(s, "-")])
		max, _ := strconv.Atoi(s[strings.Index(s, "-")+1 : strings.Index(s, " ")])
		char := s[strings.Index(s, " ")+1 : strings.Index(s, ":")]
		code := s[strings.Index(s, ":")+2:]

		charCount := strings.Count(code, char)

		if charCount >= min && charCount <= max {
			totalValid++
		}
	}

	fmt.Println(totalValid)

	fmt.Println(fmt.Sprintf("Elapsed time: %s", time.Since(start)))
}

func RunSecond() {
	start := time.Now()

	data := util.ReadFile("/two/input.txt")

	inputArr := strings.Split(string(data), "\n")

	totalValid := 0
	for _, s := range inputArr {
		first, _ := strconv.Atoi(s[0:strings.Index(s, "-")])
		second, _ := strconv.Atoi(s[strings.Index(s, "-")+1 : strings.Index(s, " ")])
		char := s[strings.Index(s, " ")+1 : strings.Index(s, ":")]
		code := s[strings.Index(s, ":")+2:]

		isValid := string(code[first-1]) == char

		if isValid && string(code[second-1]) != char {
			totalValid++
		} else if !isValid && string(code[second-1]) == char {
			totalValid++
		}
	}

	fmt.Println(totalValid)

	fmt.Println(fmt.Sprintf("Elapsed time: %s", time.Since(start)))
}
