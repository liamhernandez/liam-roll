package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Scanln(&input)
	dsplit := strings.Split(input, "d")
	num, err := strconv.ParseFloat(dsplit[0], 10)
	if err != nil {
		fmt.Println(err)
	}
	tsplit := strings.Split(dsplit[1], "t")
	side, err := strconv.ParseFloat(tsplit[0], 10)
	if err != nil {
		fmt.Println(err)
	}
	var thresh float64 = 0
	if (len(tsplit) > 1) {
		thresh, err = strconv.ParseFloat(tsplit[1], 10)
		if err != nil {
			fmt.Println(err)
		}
	}

	var rolls []uint32
	for i := 0.0; i != num; i++ {
		rolls = append(rolls, uint32(rand.Int63n(int64(side)) + 1))
	}
	fmt.Println("ROLLS : ", rolls)

	var total uint32 = 0
	for _, v := range rolls {
		total += v
	}
	fmt.Println("TOTAL : ", total)

	mu := num * (1.0 + side) / 2.0
	sigma := math.Sqrt(float64(num) * (1.0 + math.Pow(float64(side), 2.0)) / 12.0)
	if thresh == 0 {
		zscore := (float64(total) - mu) / sigma
		cdf := math.Abs(0.5 * (1 + math.Erf(zscore / math.Sqrt2)))
		fmt.Println("Percentile of roll : ", math.Round(cdf * 100))
	} else {
		zscore := (float64(thresh) - mu) / sigma
		cdf := math.Abs(0.5 * (1 + math.Erf(zscore / math.Sqrt2)))
		fmt.Println("CDF of thresh : ", cdf)
	}
}
