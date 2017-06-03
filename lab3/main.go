package main

import (
	"fmt"

	"github.com/ZzEeKkAa/numeric-methods/lab3/mid-rec"
)

func main() {
	var s mid_rec.Solver
	s.F(func(i int, x float64, y ...float64) float64 {
		switch i {
		case 0:
			return -2*y[0] + 4*y[1]
		case 1:
			return -y[0] + 3*y[1]
		}
		panic(fmt.Sprint("Wrong i:", i))
	}, 2)

	fmt.Println(s.Solve(-1, 100, 0, 3, 0))
}
