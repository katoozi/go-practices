package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/Arafatk/glot"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a data file.")
		return
	}

	fileName := os.Args[1]
	_, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("Cannot stat: ", err)
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	xP, xY := []float64{}, []float64{}
	for _, rec := range allRecords {
		x, _ := strconv.ParseFloat(rec[0], 64)
		y, _ := strconv.ParseFloat(rec[1], 64)
		xP = append(xP, x)
		xY = append(xY, y)
	}

	points := [][]float64{}
	points = append(points, xP)
	points = append(points, xY)
	fmt.Println(points)

	dimentions := 2
	persist := true
	debug := false
	plot, _ := glot.NewPlot(dimentions, persist, debug)

	plot.SetTitle("Using Glot with CSV data")
	plot.SetXLabel("X-Axis")
	plot.SetYLabel("Y-Axis")
	style := "circle"
	plot.AddPointGroup("Circle: ", style, points)
	plot.SavePlot("output.png")
}
