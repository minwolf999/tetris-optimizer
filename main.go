package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Error: no file give")
		fmt.Println("go run . [text file]")
		return
	}

	arg := os.Args[1]

	data, err := os.ReadFile(arg)
	if err != nil {
		panic(err)
	}

	listTemp := strings.Split(string(data), "\n")
	var listTetris [][]string

	//Range loop to verify format

	for i := range listTemp {
		if i%5 == 4 && i != 0 && listTemp[i] != "" {
			fmt.Println("ERROR: Bad Format !")
			return
		}
	}

	//Isolating each forms on TabTab-String

	listTetris, err = GetData(listTetris, listTemp)
	if err != nil {
		fmt.Println(err)
		return
	}

	Rename(listTetris)
	DeleteUselessPoint(listTetris)

	//Define lenght_Square tab (format)
	number := int(math.Sqrt(float64(len(listTetris) * 4)))

	if len(listTetris) != 1 {
		number += 1
	}

	res.NbrOfPoints = math.MaxInt

	newtab := FillTabByPoints(number)

	//First element became the last.
	listTetris = append(listTetris[1:], listTetris[0])

	//See commits
	Optimize(listTetris, newtab)

	if res.result == nil {
		fmt.Println("ERROR: there is an invalid Tetromino !")
		return
	}

	for _, v := range res.result {
		fmt.Println(v)
	}
}
