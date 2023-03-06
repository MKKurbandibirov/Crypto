package tests

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"sync"
)

type PokerTest struct {
	U []byte
	q int
	Freq map[int]float64
}

func NewPokerTest() *PokerTest {
	return &PokerTest{
		U: make([]byte, 0, 1000),
		q: 10,
		Freq: make(map[int]float64, 20),
	}
}

func (p *PokerTest) MakeSeries(Mseq []byte) {
	for len(Mseq)%32 != 0 {
		Mseq = Mseq[:len(Mseq)-1]
	}

	for i := 0; i < len(Mseq); i += 32 {
		tmp := Mseq[i : i+32]
		var n float64
		for j := 0; j < len(tmp); j++ {
			n += math.Pow(2, float64(j)) * float64(tmp[j])
		}

		n /= (math.Pow(2, 32) - 1)

		p.U = append(p.U, byte(n*float64(p.q)))
	}

	cvintets := make([][10]int, len(p.U)/5)

	for len(p.U) % 5 != 0 {
		p.U = p.U[:len(p.U)-1]
	}

	for i := 0; i < len(p.U)/5; i++ {
		for j := 0; j < len(p.U[i:i+5]); j++ {
			cvintets[i][p.U[i : i+5][j]]++
		}

		sort.SliceStable(cvintets[i][:], func(k, l int) bool {
			return cvintets[i][k] > cvintets[i][l]
		})

		if cvintets[i][0] == 1 {
			p.Freq[1]++
		} else if cvintets[i][0] == 2 {
			if cvintets[i][1] == 1 {
				p.Freq[2]++
			} else if cvintets[i][1] == 2 {
				p.Freq[3]++
			}
		} else if cvintets[i][0] == 3 {
			if cvintets[i][1] == 1 {
				p.Freq[4]++
			} else if cvintets[i][1] == 2 {
				p.Freq[5]++
			}
		} else if cvintets[i][0] == 4 {
			p.Freq[6]++
		} else {
			p.Freq[7]++
		}
	}

	for key, val := range p.Freq {
		p.Freq[key] = val / float64(len(cvintets))
	}
}

func (p *PokerTest) Test(out *bufio.Writer, Mseq []byte, mutex *sync.Mutex) {
	p.MakeSeries(Mseq)
	P := make([]float64, 7)

	P[0] = float64((p.q-1)*(p.q-2)*(p.q-3)*(p.q-4)) / math.Pow(float64(p.q), 4)
	P[1] = 10 * float64((p.q-1)*(p.q-2)*(p.q-3)) / math.Pow(float64(p.q), 4)
	P[2] = 15 * float64((p.q-1)*(p.q-2)) / math.Pow(float64(p.q), 4)
	P[3] = 10 * float64((p.q-1)*(p.q-2)) / math.Pow(float64(p.q), 4)
	P[4] = 10 * float64(p.q-1) / math.Pow(float64(p.q), 4)
	P[5] = 5 * float64(p.q-1) / math.Pow(float64(p.q), 4)
	P[6] = 1 / math.Pow(float64(p.q), 4)

	mutex.Lock()

	fmt.Fprintln(out, "-------------------------------")
	fmt.Fprintln(out, "--------- Poker Test ----------")
	fmt.Fprintln(out, "-------------------------------")

	for i := 1; i <= 7; i++ {
		fmt.Fprintf(out, "N_%d = %f    ---    P_%d = %f\n", i, p.Freq[i], i, P[i-1])
	}

	// TODO

	mutex.Unlock()
}
