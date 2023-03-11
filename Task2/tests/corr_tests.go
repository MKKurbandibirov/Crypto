package tests

import (
	"fmt"
	"math"
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

		c.R[k] *= 1 / float64(len(Mseq)-k)
		c.R[k] /=  math.Sqrt(D_i*D_i_k)
	}

	fmt.Println(c.R)
}

func (c *CorrTest) Test(Mseq []byte) {
	c.AutoCorrFunc(Mseq)

	N := float64(len(Mseq))
	R_cr := 1 / (N-1) + 2 / (N-2) * math.Sqrt(N*(N-3) / N+1)

	// TODO
	fmt.Println(R_cr)
}
