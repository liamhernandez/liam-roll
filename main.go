package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

//TODO Want to make an html template in a `string like this` to just insert our beastly figures into. This way, we won't have to have any files besides our executable.
// And, we won't have to call fmt.Fprint so often, especially if we run into an error.
//TODO Also want to check my math. Still not sure if it's right, such as when rolling 1d6.

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, `<!DOCTYPE HTML>
		<html lang="en-US">
		<head>
		<meta charset="UTF-8">
		<title>liam-roll</title>
		<link rel="stylesheet" type="text/css" href="/liam-roll.css" />
		</head>
		<body>
		<header>
		<h1>liam-roll</h1>
		</header>
		<p>Welcome to liam-roll</p>
		<form action="/roll">
		<label for="num">Number of dice:</label>
		<input type="number" id="num" name="num" required>
		<br>
		<label for="sides">Number of sides per die:</label>
		<input type="number" id="sides" name="sides" required>
		<br>
		<label for="thresh">Threshold to meet (leave blank to ignore):</label>
		<input type="number" id="thresh" name="thresh">
		<br>
		<input type="submit">
		</form>
		</body>
		</html>`)
	})
	http.HandleFunc("/liam-roll.css", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/css")
		fmt.Fprint(res, `
		body {
			margin: 0px;
			padding: 0px;
			background-color: #000;
			color: #fff;
			text-align: center;
		}
		header {
			width: 100%;
			height: 10rem;
			display: flex;
			justify-content: center;
			align-items: center;
			margin: 0;
		}
		header h1 {
			padding: 5%;
			position: absolute;
			font-size: 44px;
			font-weight: bolder;
			color: transparent;
			background-image: linear-gradient(to right, #c99, #9c9, #99c, #c99);
			background-clip: text;
			background-size: 200%;
			background-position: -200%;
			animation: animated-gradient 5s infinite normal linear;
		}
		@keyframes animated-gradient {
			to {
				background-position: 200%;
			}
		}
		`)
	})
	http.HandleFunc("/roll", func(res http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			fmt.Fprint(res, "Failed to parse input...")
			return
		}
		num, err := strconv.ParseFloat(req.Form.Get("num"), 10)
		if err != nil {
			fmt.Fprint(res, "Failed to parse input...")
			return
		}
		sides, err := strconv.ParseFloat(req.Form.Get("sides"), 10)
		if err != nil {
			fmt.Fprint(res, "Failed to parse input...")
			return
		}
		fmt.Fprint(res, "<p>Number of dice: ", num, "<br>")
		fmt.Fprint(res, "Number of sides per die: ", sides, "</p>")

		var rolls []uint32
		for i:= 0.0; i != num; i++ {
			rolls = append(rolls, uint32(rand.Int63n(int64(sides)) + 1))
		}
		fmt.Fprint(res ,"ROLLS : ", rolls)
		var total uint32 = 0
		for _, v := range rolls {
			total += v
		}
		fmt.Fprint(res ,"TOTAL : ", total)

		mu := num * (1.0 + sides) / 2.0
		sigma := math.Sqrt(float64(num) * (1.0 + math.Pow(float64(sides), 2.0)) / 12.0)

		if len(req.Form.Get("thresh")) > 0 {
			thresh, err := strconv.ParseFloat(req.Form.Get("thresh"), 10)
			if err != nil {
				fmt.Fprint(res, "Failed to parse input...")
				return
			}
			fmt.Fprint(res, "<p>Has thresh: ", thresh, "</p>")

			zscore := (float64(thresh) - mu) / sigma
			cdf := math.Abs(0.5 * (1 + math.Erf(zscore / math.Sqrt2)))
			fmt.Fprint(res, "Chance of meeting threshold : ", math.Round((1 - cdf) * 100), "%")
		} else {
			fmt.Fprint(res, "<p>Does not have thresh</p>")

			zscore := (float64(total) - mu) / sigma
			cdf := math.Abs(0.5 * (1 + math.Erf(zscore / math.Sqrt2)))
			fmt.Fprint(res, "Percentile of roll : ", math.Round(cdf * 100))
		}
	})

	log.Fatal(http.ListenAndServe(":4900", nil))
}
