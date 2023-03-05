package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"crypto_task_2_sub_task_1/tests"
)

var out = bufio.NewWriter(os.Stdout)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Register struct {
	uniqRegs map[string]struct{}
	Digits   []byte
	polynome []int
	L        int
}

func NewRegister(L int) *Register {
	return &Register{
		uniqRegs: make(map[string]struct{}, int(math.Pow(float64(2), float64(L)))),
		Digits:   make([]byte, L),
		L:        L,
		polynome: make([]int, 0),
	}
}

func (r *Register) GetStringDigit() string {
	var str string
	for i := 0; i < len(r.Digits); i++ {
		str += fmt.Sprint(r.Digits[i])
	}

	return str
}

func (r *Register) ParsePoly(poly string) {
	polynome := strings.Split(poly, "+")
	polynome = polynome[:len(polynome)-1]

	for i := 0; i < len(polynome); i++ {
		val, _ := strconv.Atoi(polynome[i][1:])
		r.polynome = append(r.polynome, val)
	}
}

func (r *Register) GenRegister() {
	for i := 0; i < r.L; i++ {
		r.Digits[i] = byte(rand.Intn(2))
	}

	var data string
	for i := 0; i < r.L; i++ {
		data += fmt.Sprint(r.Digits[i])
	}

	if err := WriteToFile("key.txt", data); err != nil {
		log.Fatal(err)
	}
}

func (r *Register) FeedBackFunc() byte {
	var newDigit = r.Digits[r.L-r.polynome[0]]
	for i := 1; i < len(r.polynome); i++ {
		newDigit ^= r.Digits[r.L-r.polynome[i]]
	}

	rightLast := r.Digits[len(r.Digits)-1]
	r.Digits = append([]byte{newDigit}, r.Digits[:len(r.Digits)-1]...)

	return rightLast
}

func main() {
	defer out.Flush()

	var L int
	flag.IntVar(&L, "L", 4, "Register size")

	var feedBackFunc string
	flag.StringVar(&feedBackFunc, "poly", "x4+x1+1", "Polynom for feedback function")

	var N int
	flag.IntVar(&N, "N", -1, "Result size")

	var keyFile string
	flag.StringVar(&keyFile, "file", "", "File to store key")

	flag.Parse()

	reg := NewRegister(L)
	if keyFile != "" {
		var err error
		reg.Digits, err = ReadFromFile(keyFile)
		reg.L = len(reg.Digits)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		reg.GenRegister()
	}
	reg.ParsePoly(feedBackFunc)

	MSeq := make([]byte, 0, int(math.Pow(float64(2), float64(reg.L))))

	for {
		var val = reg.GetStringDigit()
		if _, ok := reg.uniqRegs[val]; ok {
			break
		} else {
			MSeq = append(MSeq, reg.FeedBackFunc())
			reg.uniqRegs[val] = struct{}{}
		}
	}

	// if N == -1 {
	// 	for i := 0; i < len(MSeq); i++ {
	// 		fmt.Fprint(out, MSeq[i], " ")
	// 	}
	// } else {
	// 	for i := 0; i < N; i++ {
	// 		fmt.Fprint(out, MSeq[i], " ")
	// 	}
	// }
	// fmt.Fprintln(out)

	serial := tests.NewSerialTest(MSeq, 4)

	serial.Test(0)
}
