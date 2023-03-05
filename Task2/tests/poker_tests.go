package tests

import (
	"fmt"
	"math"
)

type PokerTest struct {
	U []byte
	q int
}

func NewPokerTest() *PokerTest {
	return &PokerTest{
		U: make([]byte, 0, 1000),
		q: 10,
	}
}

func (p *PokerTest) MakeSeries(Mseq []byte) {
	for len(Mseq) % 32 != 0 {
		Mseq = Mseq[:len(Mseq)-1]
	}

	for i := 0; i < len(Mseq); i += 32 {
		tmp := Mseq[i:i+32]
		var n float64
		for j := 0; j < len(tmp); j++ {
			n += math.Pow(2, float64(j)) * float64(tmp[j])
		}
		
		n /= (math.Pow(2, 32)-1)

		p.U = append(p.U, byte(n * float64(p.q)))
	}

	for i := 0; i < len(p.U); i++ {
		fmt.Println(p.U[i])
	}
}
