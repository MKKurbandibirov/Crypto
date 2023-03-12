package tests

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"sync"
)

type PokerTest struct {
	U     []byte
	q     int
	Count map[int]int
	M     int
	Hi    float64

	CriticalHi map[float64]float64
}

func NewPokerTest() *PokerTest {
	return &PokerTest{
		U:     make([]byte, 0, 1000),
		q:     10,
		Count: make(map[int]int, 20),

		CriticalHi: map[float64]float64{
			0.95: 1.645, 0.9: 2.2, 0.8: 3.07, 0.2: 8.56, 0.1: 10.64, 0.05: 12.59,
		},
	}
}

func (p *PokerTest) MakeSeries(Mseq []byte) {
	for len(Mseq)%32 != 0 {
		Mseq = Mseq[:len(Mseq)-1]
	}
	p.M = len(Mseq)

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

	for len(p.U)%5 != 0 {
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
			p.Count[1]++
		} else if cvintets[i][0] == 2 {
			if cvintets[i][1] == 1 {
				p.Count[2]++
			} else if cvintets[i][1] == 2 {
				p.Count[3]++
			}
		} else if cvintets[i][0] == 3 {
			if cvintets[i][1] == 1 {
				p.Count[4]++
			} else if cvintets[i][1] == 2 {
				p.Count[5]++
			}
		} else if cvintets[i][0] == 4 {
			p.Count[6]++
		} else {
			p.Count[7]++
		}
	}
}

func (p *PokerTest) Test(out *bufio.Writer, Mseq []byte, mutex *sync.Mutex, alpha float64) {
	p.MakeSeries(Mseq)
	P := make([]int, 7)

	P[0] = int(float64((p.q-1)*(p.q-2)*(p.q-3)*(p.q-4)) / math.Pow(float64(p.q), 4) * float64(p.M) / 160)
	P[1] = int(10 * float64((p.q-1)*(p.q-2)*(p.q-3)) / math.Pow(float64(p.q), 4) * float64(p.M) / 160)
	P[2] = int(15 * float64((p.q-1)*(p.q-2)) / math.Pow(float64(p.q), 4) * float64(p.M) / 160)
	P[3] = int(10 * float64((p.q-1)*(p.q-2)) / math.Pow(float64(p.q), 4) * float64(p.M) / 160)
	P[4] = int(10 * float64(p.q-1) / math.Pow(float64(p.q), 4) * float64(p.M) / 160)
	P[5] = int(5 * float64(p.q-1) / math.Pow(float64(p.q), 4) * float64(p.M) / 160)
	P[6] = int(1 / math.Pow(float64(p.q), 4) * float64(p.M) / 160)
	
	for i := 0; i < 7; i++ {
		if P[i] != 0 {
			p.Hi += math.Pow(float64(p.Count[i+1] - P[i]), 2) / float64(P[i])
		} else {

		}
	}

	mutex.Lock()

	fmt.Fprintln(out, "------------------------------------")
	fmt.Fprintln(out, "------------ Poker Test ------------")
	fmt.Fprintln(out, "------------------------------------")

	fmt.Fprintln(out, "- Эмпирические и эталонные частоты -")
	for i := 1; i <= 7; i++ {
		fmt.Fprintf(out, "N_%d = %-4d  -\tP_%d = %-4d\n", i, p.Count[i], i, P[i-1])
	}

	// TODO
	fmt.Fprintln(out, "----------- Критерий 𝒳^2 -----------")
	fmt.Fprintf(out, "𝒳^2 = %f\n", p.Hi)

	fmt.Fprintln(out, "------------------------------------")
	if alpha == 0 {
		HiMin, HiMax := p.CriticalHi[0.1], p.CriticalHi[0.9]
		if HiMax <= p.Hi && HiMin >= p.Hi {
			fmt.Fprintln(out, "-------> Poker test passed! <-------")
		} else {
			fmt.Fprintln(out, "-------> Poker test failed! <-------")
		}
	} else {
		HiMin := p.CriticalHi[alpha]
		if HiMin >= p.Hi {
			fmt.Fprintln(out, "----> Poker test passed! <-------")
		} else {
			fmt.Fprintln(out, "-------> Poker test failed! <-------")
		}
	}
	fmt.Fprintln(out, "------------------------------------")

	fmt.Fprintln(out)

	mutex.Unlock()
}
