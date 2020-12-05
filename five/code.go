package five

import (
	"aoc/util"
	"math"
	"sort"
	"strings"
)

func RunFirst() int {
	allSeats := getSeatIdsSorted()
	return allSeats[len(allSeats)-1]
}

func RunSecond() int {
	allSeats := getSeatIdsSorted()

	expectingSeat := allSeats[0]
	for _, i := range allSeats {
		if i != expectingSeat {
			return expectingSeat
		}

		expectingSeat++
	}

	return 0
}

func getSeatIdsSorted() []int {
	data := util.ReadFile("/five/input.txt")

	inputArr := strings.Split(string(data), "\n")

	allSeats := make([]int, 0)
	for _, s := range inputArr {

		row, col := calcSeat(s)

		total := row*8 + col

		allSeats = append(allSeats, total)
	}

	sort.Ints(allSeats)

	return allSeats
}

func calcSeat(s string) (int, int) {
	rowMin := 0.0
	rowMax := 127.0

	for i := 0; i < 7; i++ {

		if string(s[i]) == "B" {
			rowMin = math.Ceil(rowMax - float64(rowMax-rowMin)/2.0)
		} else {
			rowMax = math.Floor(rowMax - float64(rowMax-rowMin)/2.0)
		}
	}

	colMin := 0.0
	colMax := 7.0

	for i := 7; i < 10; i++ {
		if string(s[i]) == "R" {
			colMin = math.Ceil(colMax - float64(colMax-colMin)/2.0)
		} else {
			colMax = math.Floor(colMax - float64(colMax-colMin)/2.0)
		}
	}

	return int(rowMin), int(colMin)
}
