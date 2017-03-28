package main

import (
	"log"
	"math"

	"github.com/alanwj/go-poly"
)

func BuildPol1(pol poly.Poly, a, b float64) poly.Poly {
	fa := pol.Eval(a)
	fb := pol.Eval(b)
	c := (fa - fb) / (a - b)
	var d float64
	x1, x2 := Sol(pol.Der().Sub(poly.New(c)))
	if a < x1 && x1 < b {
		fx := pol.Eval(x1)
		d = (fa + fx - c*(a+x1)) / 2
	} else if a < x2 && x2 < b {
		fx := pol.Eval(x2)
		d = (fa + fx - c*(a+x2)) / 2
	} else {
		log.Println("No solution in [a,b]")
	}
	return poly.New(d, c)
}

func Sol(pol poly.Poly) (x1, x2 float64) {
	a, b, c := pol.Coeff(2), pol.Coeff(1), pol.Coeff(0)
	d := b*b - 4*a*c
	if d < 0 {
		log.Println("D < 0")
	} else {
		d = math.Sqrt(d)
		x1 = (-b - d) / 2 / a
		x2 = (-b + d) / 2 / a
	}
	return
}
