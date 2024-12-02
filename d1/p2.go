package d1

import (
	"aoc-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func P2() {
	lines := utils.GetLines("d1/p2.txt")
	lNums := []float64{}
	rNums := []float64{}
	multMap := map[float64]int{}
	score := 0.0

	for _, line := range lines {
		numStrs := strings.Split(line, "   ")
		n0, err := strconv.ParseFloat(numStrs[0], 64)
		if err != nil {
			panic(err)
		}
		n1, err := strconv.ParseFloat(numStrs[1], 64)
		if err != nil {
			panic(err)
		}
		lNums = append(lNums, n0)
		rNums = append(rNums, n1)

		_, ok := multMap[n1]
		if !ok {
			multMap[n1] = 0
		}
		multMap[n1]++
	}

	for _, lNum := range lNums {
		mult, ok := multMap[lNum]
		if !ok {
			mult = 0.0
		}

		score += lNum * float64(mult)
	}

	fmt.Printf("score = %f\n", score)
}
