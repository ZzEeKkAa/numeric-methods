package mid_rec

type Solver struct {
	f func(int, float64, ...float64) float64
	n int
}

// Define function f at u'(x)=f(x,u) where n - is the number of dimensions
// First argument is the number of function at i'th dimension
func (s *Solver) F(f func(int, float64, ...float64) float64, n int) {
	s.f = f
	s.n = n
}

// N is the number of steps
func (s *Solver) Solve(x float64, N int, x0 float64, y0 ...float64) []float64 {
	h := (x - x0) / float64(N)
	y1 := make([]float64, len(y0))
	for i := 0; i < N; i, x0 = i+1, x0+h {
		// Predictor
		for j := 0; j < s.n; j++ {
			y1[j] = h / 2 * s.f(j, x0, y0...)
		}
		for j := 0; j < s.n; j++ {
			y1[j] += y0[j]
		}

		// Connector
		for j := 0; j < s.n; j++ {
			y0[j] += h * s.f(j, x0+h/2, y1...)
		}
	}

	return y0
}
