package tests

import (
	"fmt"
	"math"
)

type SerialTest struct {
	MSeries map[string]int
	MFreq   map[string]float64
	Hi      float64
	MSeq    []byte
	K       int
	N       int
}

func NewSerialTest(MSeq []byte, k int) *SerialTest {
	return &SerialTest{
		MSeq:    MSeq,
		K:       k,
		N:       len(MSeq) / k,
		MSeries: make(map[string]int, 100),
		MFreq:   make(map[string]float64, 100),
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
	// --------------------- //
	for len(s.MSeq) % s.K != 0 {
		s.MSeq = s.MSeq[:len(s.MSeq)-1]
	}

	for i := 0; i < len(s.MSeq); i += s.K {
		s.MSeries[getStringSeries(s.MSeq[i:i+s.K])]++
	}

	for key, val := range s.MSeries {
		s.MFreq[key] = float64(val) / (float64(len(s.MSeq)) / float64(s.K))
	}

	fmt.Println(s.MSeries)
	fmt.Println(s.MFreq)
}

func (s *SerialTest) CountNs() {
	NTeor := (float64(s.N) / math.Pow(2, float64(s.K)))// / (float64(len(s.MSeq)) / float64(s.K))

	fmt.Println("N_teor =", NTeor)

	for _, val := range s.MSeries {
		s.Hi += math.Pow(float64(val) - NTeor, 2) / NTeor
	}

	fmt.Println(s.Hi)
}
