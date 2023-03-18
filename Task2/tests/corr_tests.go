package tests

import (
	"bufio"
	"fmt"
	"math"
	"sync"
)

type CorrTest struct {
	K []int
	R map[int]float64
}

func NewCorrTest() *CorrTest {
	return &CorrTest{
		K: []int{1, 2, 8, 9},
		R: make(map[int]float64, 4),
	}
}

func (c *CorrTest) AutoCorrFunc(Mseq []byte) {
	for _, k := range c.K {
		m_i := 0.0
		for j := 0; j < len(Mseq)-k; j++ {
			m_i += float64(Mseq[j])
		}
		m_i *= 1 / float64(len(Mseq)-k)

		m_i_k := 0.0
		for j := k; j < len(Mseq); j++ {
			m_i_k += float64(Mseq[j])
		}
		m_i_k *= 1 / float64(len(Mseq)-k)

		D_i := 0.0
		for j := 0; j < len(Mseq)-k; j++ {
			D_i += math.Pow(float64(Mseq[j])-m_i, 2)
		}
		D_i *= 1 / float64(len(Mseq)-k-1)

		D_i_k := 0.0
		for j := k; j < len(Mseq); j++ {
			D_i_k += math.Pow(float64(Mseq[j])-m_i_k, 2)
		}
		D_i *= 1 / float64(len(Mseq)-k-1)

		for i := 0; i < len(Mseq)-k; i++ {
			c.R[k] += (float64(Mseq[i]) - float64(m_i)) * (float64(Mseq[i+k]) - float64(m_i_k))
		}
		c.R[k] /= float64(len(Mseq)-k)
		c.R[k] /= math.Sqrt(D_i * D_i_k)
		c.R[k] = math.Abs(c.R[k])
	}

}

func (c *CorrTest) Test(Mseq []byte, out *bufio.Writer, mutex *sync.Mutex) {
	c.AutoCorrFunc(Mseq)

	N := float64(len(Mseq))
	R_cr := 1/(N-1) + 2/(N-2)*math.Sqrt(N*(N-3)/N+1)

	mutex.Lock()

	fmt.Fprintln(out, "--------------------------------")
	fmt.Fprintln(out, "---------- Corr Test -----------")
	fmt.Fprintln(out, "--------------------------------")

	pass := true
	fmt.Fprintln(out, "-- Автокорреляционная функция --")
	for k, r := range c.R {
		if r > R_cr {
			pass = false
		}

		fmt.Fprintf(out, "R[%d]: %f\n", k, r)
	}

	fmt.Fprintln(out, "----- Критическое значение -----")
	fmt.Fprintf(out, "R[k]_cr: %.6f\n", R_cr)

	fmt.Fprintln(out, "--------------------------------")
	if pass {
		fmt.Fprintln(out, "-----> Corr test passed! <------")
	} else {
		fmt.Fprintln(out, "-----> Corr test failed! <------")
	}
	fmt.Fprintln(out, "--------------------------------")

	mutex.Unlock()	 
}
