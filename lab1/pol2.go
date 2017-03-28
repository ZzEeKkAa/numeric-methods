package main

import (
	"github.com/alanwj/go-poly"
)

func BuildPol2(pol poly.Poly, a, b float64) poly.Poly {
	t := poly.New(0, -3, 0, 4)
	t = MovePol(t, poly.New((a-b)/(a+b), 2/(a+b)))
	d := pol.Coeff(t.Deg()) / t.Coeff(t.Deg())
	coeff := make([]float64, 0, t.Deg()+1)
	for j := 0; j <= t.Deg(); j++ {
		coeff = append(coeff, d*t.Coeff(j))
	}

	res := pol.Sub(poly.New(coeff...))
	return res
}

func MovePol(p, q poly.Poly) poly.Poly {
	rp := poly.New(p.Coeff(0))
	for i := p.Deg(); i > 0; i-- {
		tp := poly.New(1)
		for j := 0; j < i; j++ {
			tp = tp.Mul(q)
		}

		coeff := make([]float64, 0, tp.Deg()+1)
		for j := 0; j <= tp.Deg(); j++ {
			coeff = append(coeff, tp.Coeff(j))
		}
		for k, c := range coeff {
			coeff[k] = c * p.Coeff(i)
		}

		rp = rp.Add(poly.New(coeff...))
	}
	return rp
}
