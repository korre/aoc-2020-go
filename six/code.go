package six

import (
	"aoc/util"
	"fmt"
	"strings"
)

func RunFirst() {
	data := util.ReadFile("/six/input.txt")

	inputArr := strings.Split(string(data), "\n\n")

	result := 0
	for _, s := range inputArr {
		s = strings.ReplaceAll(s, "\n", "")

		m := make(map[string]int)

		for _, c := range s {
			_, ok := m[string(c)]

			if !ok {
				m[string(c)] = 1
			}
		}

		for _, value := range m {
			result += value
		}
	}

	fmt.Println(result)
}

func RunSecond() {
	data := util.ReadFile("/six/input.txt")

	inputArr := strings.Split(string(data), "\n\n")

	result := 0
	for _, s := range inputArr {
		personArr := strings.Split(s, "\n")

		m := make(map[string]int)

		for _, p := range personArr {
			for _, c := range p {
				val, ok := m[string(c)]

				if !ok {
					m[string(c)] = 1
				} else {
					m[string(c)] = val + 1
				}
			}
		}

		for _, value := range m {
			if value == len(personArr) {
				result++
			}
		}
	}

	fmt.Println(result)
}
