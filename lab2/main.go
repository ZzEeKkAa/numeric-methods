package main

import (
	"log"

	"flag"

	"math"

	poly "github.com/alanwj/go-poly"
)

var (
	step   = flag.Float64("s", 0.05, "Step for searching root")
	its    = flag.Int("it", 100000, "Max iterrations for searching root")
	gCount = flag.Int("gc", 10000000, "Count of iterration to calculate gamma function by Eiler")
	eps    = flag.Float64("eps", 0.000001, "Precision for roots of polinomial")
)

func main() {
	flag.Parse()

	n := 4
	// Set function to calculate // int[0,inf](x^alpha * e^(-x) * f(x) dx
	alpha := 1.
	f := func(x float64) float64 {
		return 1 / (1 + math.Exp(-x) + math.Exp(-2*x))
	}

	l := getLager(n, alpha)
	log.Printf("L[%d,%.6f] = %s\n", n, alpha, l.String())
	roots := findRoots(l, *eps)
	log.Println(roots)
	res := 0.

	for _, xk := range roots {
		ck := getLagerC(l, xk, alpha)
		log.Print(ck, " * ", f(xk), " + ")
		res += ck * f(xk)
	}

	log.Println("I = ", res)
}

func getLagerC(p poly.Poly, x float64, alpha float64) float64 {
	//log.Println(x)
	res := gamma(alpha+float64(p.Deg())+1) / x
	//log.Println(res)
	for i := 1; i <= p.Deg(); i++ {
		res *= float64(i)
	}
	//log.Println(res)
	tp := p.Der()
	res /= tp.Mul(tp).Eval(x)
	//log.Println(res)
	return res
}

func gamma(z float64) float64 {
	res := 1 / z
	for n := 1; n <= *gCount; n++ {
		res *= math.Pow(1+1/float64(n), z) / (1 + z/float64(n))
	}
	//log.Printf("G(%f)=%f\n", z, res)
	return res
}

func getLager(n int, alpha float64) poly.Poly {
	switch n {
	case -1:
		return poly.New(0)
	case 0:
		return poly.New(1)
	}
	return getLager(n-1, alpha).Mul(poly.New(float64(2*n-1)+alpha, -1)).Add(getLager(n-2, alpha).Mul(poly.New(-(float64(n-1) + alpha)))).Mul(poly.New(1 / float64(n)))
}

func findRoots(p poly.Poly, eps float64) []float64 {
	var res []float64
	x := 0.
	positive := p.Eval(x) > 0
	for i := 0; i < *its+1; i++ {
		if p.Deg() == len(res) {
			return res
		}
		x += *step
		tp := p.Eval(x) > 0
		if positive == !tp {
			res = append(res, findRoot(p, x-*step, x, eps))
			positive = tp
		}
	}
	log.Fatal("Can't find all roots! Found only:", res)
	return nil
}

func findRoot(p poly.Poly, a, b, eps float64) float64 {
	if p.Eval(a)*p.Eval(b) > 0 {
		log.Fatal("There is no roots of polinomial at [", a, b, "]")
	}
	if p.Eval(a) > 0 {
		a, b = b, a
	}
	m := (b + a) / 2
	for 2*(b-a) > eps {
		m = (b + a) / 2
		if p.Eval(m) > 0 {
			b = m
		} else {
			a = m
		}
	}
	return m
}
