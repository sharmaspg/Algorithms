package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sharmaspg/Algorithms/sorting"
	"github.com/sharmaspg/Algorithms/utils"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func Plotter(title string, timegraph []float64, input []int64) {

	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = "size"
	p.Y.Label.Text = "time"

	err := plotutil.AddLinePoints(p,
		"Merge", PlotPoints(timegraph, input))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "merge.png"); err != nil {
		panic(err)
	}
}

func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}

func PlotPoints(timegraph []float64, input []int64) plotter.XYs {
	pts := make(plotter.XYs, len(input))
	for i := range pts {
		pts[i].X = float64(input[i])
		pts[i].Y = timegraph[i]
		//	pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}

func main() {
	// d := utils.GetArray(10)
	// r := sorting.MergeSort(d)
	// fmt.Println(r)
	// return

	rows := utils.GetInputForGraph()
	timegraph := []float64{}
	inputrow := []int64{}
	for _, row := range rows {
		started := time.Now()
		inputrow = append(inputrow, int64(len(row)))
		sorting.MergeSort(row)
		ended := time.Now()
		duration := ended.Sub(started)
		timegraph = append(timegraph, duration.Seconds())
		fmt.Println(duration.Seconds())
	}

	Plotter("MergeSort", timegraph, inputrow)
}
