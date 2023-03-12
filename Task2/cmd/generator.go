package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

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
