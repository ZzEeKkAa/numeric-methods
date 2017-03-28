package main

import (
	"log"

	"image/color"

	"github.com/alanwj/go-poly"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

func main() {
	a, b := 1., 8.
	pol := poly.New(2, 0, 2, 5)

	//a, b = 0., 1.
	//pol = poly.New(0, 0, 0, 1)
	pol0 := BuildPol0(pol, a, b)
	pol1 := BuildPol1(pol, a, b)
	pol2 := BuildPol2(pol, a, b)
	pol3 := BuildTrig(pol, a, b, 15)

	fPol := plotter.NewFunction(pol.Eval)
	fPol.Color = color.RGBA{R: 255, G: 0, B: 0, A: 1}
	fPol0 := plotter.NewFunction(pol0.Eval)
	fPol0.Color = color.RGBA{R: 255, G: 0, B: 255, A: 1}
	fPol1 := plotter.NewFunction(pol1.Eval)
	fPol1.Color = color.RGBA{R: 0, G: 255, B: 0, A: 1}
	fPol2 := plotter.NewFunction(pol2.Eval)
	fPol2.Color = color.RGBA{R: 0, G: 255, B: 255, A: 1}
	fPol3 := plotter.NewFunction(pol3.Eval)
	fPol3.Color = color.RGBA{R: 255, G: 255, B: 0, A: 1}

	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = a-1, b+1
	pl.Y.Min, pl.Y.Max = -1000, 3400
	pl.Add(fPol)
	pl.Add(fPol0)
	pl.Add(fPol1)
	pl.Add(fPol2)
	pl.Add(fPol3)

	if err := pl.Save(14*vg.Inch, 14*vg.Inch, "pol.png"); err != nil {
		panic(err.Error())
	}

	log.Println(pol)
	log.Println(pol0)
	log.Println(pol1)
	log.Println(pol2)
}
