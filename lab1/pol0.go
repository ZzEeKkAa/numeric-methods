package main

import (
	poly "github.com/alanwj/go-poly"
)

func BuildPol0(pol poly.Poly, a, b float64) poly.Poly {
	return poly.New((pol.Eval(a) + pol.Eval(b)) / 2)
}
