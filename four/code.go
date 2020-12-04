package four

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func RunFirst() {
	data := util.ReadFile("/four/input.txt")

	inputArr := strings.Split(string(data), "\n\n")

	var re = regexp.MustCompile(`(byr)|(iyr)|(eyr)|(hgt)|(hcl)|(ecl)|(pid)`)

	valid := 0
	for _, s := range inputArr {
		if len(re.FindAllStringIndex(s, -1)) == 7 {
			valid++
		}
	}

	fmt.Println(valid)
}

func RunSecond() {
	data := util.ReadFile("/four/input.txt")

	inputArr := strings.Split(string(data), "\n\n")

	var re = regexp.MustCompile(`(byr)|(iyr)|(eyr)|(hgt)|(hcl)|(ecl)|(pid)`)

	valid := 0
	for _, a := range inputArr {

		line := strings.ReplaceAll(a, "\n", " ")

		if len(re.FindAllStringIndex(line, -1)) != 7 {
			continue
		}

		fields := strings.Split(line, " ")

		validCount := 0

		for _, s := range fields {
			switch strings.Split(s, ":")[0] {
			case "byr":
				if validRange(strings.Split(s, ":")[1], 1920, 2002) {
					validCount++
				}

			case "iyr":
				if validRange(strings.Split(s, ":")[1], 2010, 2020) {
					validCount++
				}
			case "eyr":
				if validRange(strings.Split(s, ":")[1], 2020, 2030) {
					validCount++
				}
			case "hgt":
				if validHgt(strings.Split(s, ":")[1]) {
					validCount++
				}
			case "hcl":
				if validHcl(strings.Split(s, ":")[1]) {
					validCount++
				}
			case "ecl":
				if validEcl(strings.Split(s, ":")[1]) {
					validCount++
				}
			case "pid":
				if validPid(strings.Split(s, ":")[1]) {
					validCount++
				}
			}
		}

		if validCount == 7 {
			valid++
		}
	}

	fmt.Println(valid)
}

func validRange(input string, min int, max int) bool {
	var re = regexp.MustCompile(`[0-9]{` + strconv.Itoa(len(strconv.Itoa(min))) + `}`)
	val, _ := strconv.Atoi(input)

	return re.MatchString(input) && val >= min && val <= max
}

func validHgt(input string) bool {
	result := false
	if strings.Contains(input, "cm") {
		result = validRange(input[0:strings.Index(input, "cm")], 150, 193)
	} else if strings.Contains(input, "in") {
		result = validRange(input[0:strings.Index(input, "in")], 59, 76)
	}

	return result
}

func validHcl(input string) bool {
	var re = regexp.MustCompile(`[a-f0-9]{6}`)
	return string(input[0]) == "#" && re.MatchString(input[1:])
}

func validEcl(input string) bool {
	validValues := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}

	_, ok := validValues[input]
	return ok
}

func validPid(input string) bool {
	var re = regexp.MustCompile(`[0-9]{9}`)
	return re.MatchString(input)
}
