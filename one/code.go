package one

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func RunFirst() {
	start := time.Now()

	data := util.ReadFile("/one/input.txt")

	inputArr := strings.Split(string(data), "\n")

out:
	for _, outer := range inputArr {
		i, _ := strconv.Atoi(outer)

		for _, inner := range inputArr {
			j, _ := strconv.Atoi(inner)

			if i+j == 2020 {
				fmt.Println(i * j)
				break out
			}
		}
	}

	fmt.Println(fmt.Sprintf("Elapsed time: %s", time.Since(start)))
}

func RunSecond() {
	start := time.Now()

	data := util.ReadFile("/one/input.txt")

	inputArr := strings.Split(string(data), "\n")

out:
	for _, outer := range inputArr {
		i, _ := strconv.Atoi(outer)

		for _, inner := range inputArr {
			j, _ := strconv.Atoi(inner)

			for _, inner := range inputArr {
				k, _ := strconv.Atoi(inner)

				if i+j+k == 2020 {
					fmt.Println(i * j * k)
					break out
				}
			}
		}
	}

	fmt.Println(fmt.Sprintf("Elapsed time: %s", time.Since(start)))
}
