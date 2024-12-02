package d1

import (
	"aoc-2024/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func P1() {
	lines := utils.GetLines("d1/p1.txt")
	lNums := []float64{}
	rNums := []float64{}
	diff := 0.0

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

	}
	sort.Float64s(lNums)
	sort.Float64s(rNums)

	for i := range lNums {
		diff += math.Abs(lNums[i] - rNums[i])
	}

	fmt.Printf("diff = %f\n", diff)
}
