package main

import (
	"github.com/alanwj/go-poly"
	"fmt"
	"log"
	"math"
)

func BuildTrig(pol poly.Poly, a, b float64, r int) tPoly {
	p := tPoly{a:[]float64{},b:[]float64{}}

	for i:=0; i<r; i++{
		t:=0.
		for j:=0; j<pol.Deg(); j++{
			t+=pol.Coeff(j)*kpCos(i,j)
		}
		p.a=append(p.a,t/math.Pi)
		t=0.
		for j:=0; j<pol.Deg(); j++{
			t+=pol.Coeff(j)*kpSin(i,j)
		}
		p.b=append(p.a,t/math.Pi)
	}
	p.a[0]/=2


	log.Println(p)
	return p
}


type tPoly struct{
	a, b []float64
}

func(tp tPoly) String() string{
	var res string
	if len(tp.a)>0{
		res += fmt.Sprintf("%.3f",tp.a[0])
	}
	n:= len(tp.a)
	if n<len(tp.b){
		n = len(tp.b)
	}

	for i:=1; i<n; i++{
		if len(tp.a)>i{
			k:= tp.a[i]
			if k>=0{
				res += fmt.Sprintf(" + %.3fcos(%dx)",k,i)
			} else {
				res += fmt.Sprintf(" - %.3fcos(%dx)",-k,i)
			}
		}
		if len(tp.b)>i{
			k:= tp.b[i]
			if k>=0{
				res += fmt.Sprintf(" + %.3fsin(%dx)",k,i)
			} else {
				res += fmt.Sprintf(" - %.3fsin(%dx)",-k,i)
			}
		}
	}

	return res
}

func(tp tPoly) Eval(x float64) float64{
	res :=0.
	for i,a:=range tp.a{
		res+=a*math.Cos(float64(i)*x)
	}
	for i,b:=range tp.b{
		res+=b*math.Sin(float64(i)*x)
	}

	return res
}

func kpCos(k,p int) float64{
	if k ==0{
		res := 1.
		for i:=0; i<p; i++{
			res *= 2*math.Pi
		}
		return res
	}
	if p ==0{
		return 0
	}
	res := -1./float64(k)
	for i:=0; i<p; i++{
		res *= 2*math.Pi
	}
	return res + float64(p)/float64(k) * kpSin(k,p-1)
}

func kpSin(k,p int) float64{
	if p==0{
		return 0
	}
	res := -1./float64(k)
	for i:=0; i<p; i++{
		res *= 2*math.Pi
	}
	return -float64(p)/float64(k) * kpCos(k,p-1)
}