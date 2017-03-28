package main

import (
	"log"

	"image/color"

	"math"

	"github.com/alanwj/go-poly"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

func main() {
	a, b := 1., 8.
	//a, b := -math.Pi, math.Pi
	pol := poly.New(2, 0, 2, 5)

	//a, b = 0., 1.
	//pol = poly.New(0, 0, 0, 1)
	pol0, df0 := BuildPol0(pol, a, b)
	pol1, df1 := BuildPol1(pol, a, b)
	pol2, df2 := BuildPol2(pol, a, b)
	pol3 := BuildTrig(pol, a, b, 50)

	fPol := plotter.NewFunction(pol.Eval)
	fPol.Color = color.RGBA{R: 255, G: 0, B: 0, A: 1}
	fPol0 := plotter.NewFunction(pol0.Eval)
	fPol0.Color = color.RGBA{R: 255, G: 0, B: 255, A: 1}
	fPol1 := plotter.NewFunction(pol1.Eval)
	fPol1.Color = color.RGBA{R: 0, G: 255, B: 0, A: 1}
	fPol2 := plotter.NewFunction(pol2.Eval)
	fPol2.Color = color.RGBA{R: 0, G: 255, B: 255, A: 1}
	fPol3 := plotter.NewFunction(func(x float64) float64 {
		return pol3.Eval(2 * math.Pi * (x - a) / (b - a))
	})
	fPol3.Color = color.RGBA{R: 255, G: 255, B: 0, A: 1}

	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = a-1, b+1
	pl.Y.Min, pl.Y.Max = -500, 3500
	pl.Add(fPol)
	pl.Add(fPol0)
	pl.Add(fPol1)
	pl.Add(fPol2)
	pl.Add(fPol3)

	if err := pl.Save(14*vg.Inch, 14*vg.Inch, "pol.png"); err != nil {
		panic(err.Error())
	}

	log.Println("Original:", pol)
	log.Println("Q0:", pol0, df0)
	log.Println("Q1:", pol1, df1)
	log.Println("Q2:", pol2, df2)
	log.Println("T50:", pol3)
}
