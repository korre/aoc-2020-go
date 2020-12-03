package three

import (
	"aoc/util"
	"strings"
)

func RunFirst(right int, down int) int {
	data := util.ReadFile("/three/input.txt")

	inputArr := strings.Split(string(data), "\n")

	treeHits := 0
	row := down
	column := right
	for i, s := range inputArr {
		if column >= len(s) {
			column = column - len(s)
		}

		if i == row {
			if string(s[column]) == "#" {
				treeHits++
			}

			row += down
			column += right
		}
	}

	return treeHits
}

func RunSecond() int {
	return RunFirst(1, 1) *
		RunFirst(3, 1) *
		RunFirst(5, 1) *
		RunFirst(7, 1) *
		RunFirst(1, 2)
}
