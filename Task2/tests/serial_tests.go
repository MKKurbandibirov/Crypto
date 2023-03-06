package tests

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sync"
)

var out = bufio.NewWriter(os.Stdout)

type SerialTest struct {
	MSeries map[string]int
	MFreq   map[string]float64
	Hi      float64
	MSeq    []byte
	K       int
	N       int
	N_T     float64

	CriticalHi map[int]map[float64]float64
}

func NewSerialTest(MSeq []byte, k int) *SerialTest {
	return &SerialTest{
		MSeq:    MSeq,
		K:       k,
		N:       len(MSeq) / k,
		MSeries: make(map[string]int, 100),
		MFreq:   make(map[string]float64, 100),

		CriticalHi: map[int]map[float64]float64{
			2: {0.95: 0.352, 0.9: 0.584, 0.8: 1.005, 0.2: 4.6, 0.1: 6.251, 0.05: 7.815},
			3: {0.95: 2.167, 0.9: 2.833, 0.8: 3.28, 0.2: 9.8, 0.1: 12.017, 0.05: 14.057},
			4: {0.95: 7.261, 0.9: 8.547, 0.8: 10.31, 0.2: 19.31, 0.1: 22.307, 0.05: 24.996},
		},
	}
}

func getStringSeries(series []byte) string {
	var str string
	for i := 0; i < len(series); i++ {
		str += fmt.Sprint(series[i])
	}

	return str
}

func (s *SerialTest) CountSeries() {
	for len(s.MSeq)%s.K != 0 {
		s.MSeq = s.MSeq[:len(s.MSeq)-1]
	}

	for i := 0; i < len(s.MSeq); i += s.K {
		s.MSeries[getStringSeries(s.MSeq[i:i+s.K])]++
	}

	for key, val := range s.MSeries {
		s.MFreq[key] = float64(val) / (float64(len(s.MSeq)) / float64(s.K))
	}
}

func (s *SerialTest) CountNs() {
	s.N_T = (float64(s.N) / math.Pow(2, float64(s.K)))

	for _, val := range s.MSeries {
		s.Hi += math.Pow(float64(val)-s.N_T, 2) / s.N_T
	}
}

func (s *SerialTest) Test(out *bufio.Writer, alpha float64, mutex *sync.Mutex) {
	s.CountSeries()
	s.CountNs()

	mutex.Lock()

	fmt.Fprintln(out, "-------------------------------")
	fmt.Fprintln(out, "-------- Serial Test ----------")
	fmt.Fprintln(out, "-------------------------------")

	fmt.Fprintln(out, "----- Эмпирические частоты ----")
	for key, N_I := range s.MSeries {
		fmt.Fprintf(out, "%s: %d\n", key, N_I)
	}

	fmt.Fprintln(out, "------ Эталонная частота ------")
	fmt.Fprintln(out, "N_T =", s.N_T)

	fmt.Fprintln(out, "--------- Критерий 𝒳^2 --------")
	fmt.Fprintln(out, "𝒳^2 =", s.Hi)

	fmt.Fprintln(out, "-------------------------------")
	if alpha == 0 {
		HiMin, HiMax := s.CriticalHi[s.K][0.1], s.CriticalHi[s.K][0.9]
		if HiMax <= s.Hi && HiMin >= s.Hi {
			fmt.Fprintln(out, "----> Serial test passed! <----")
		} else {
			fmt.Fprintln(out, "----> Serial test failed! <----")
		}
	} else {
		HiMin := s.CriticalHi[s.K][alpha]
		if HiMin >= s.Hi {
			fmt.Fprintln(out, "----> Serial test passed! <----")
		} else {
			fmt.Fprintln(out, "----> Serial test failed! <----")
		}
	}
	fmt.Fprintln(out, "-------------------------------")
	fmt.Fprintln(out)

	mutex.Unlock()
}
