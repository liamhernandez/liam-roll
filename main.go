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

	if thresh == 0 {
		denom := math.Pow(side, num)
		//TODO make this more efficient :)
		average := float64(total) / num
		lessthan := 0.0
		for i := 1; i <= int(num); i++ {
			roll := 0
			for j := 1; i <= int(side); i++ {
				if i + j > int(average) {
					roll += // CONFUSE
				}
			}
		}
		prob := lessthan / denom
		fmt.Println("Probability of roll : ", prob)
	} else {
	}
}
