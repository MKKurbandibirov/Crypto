package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Register struct {
	uniqRegs map[string]struct{}
	Digits   []int
	polynome []int
	L        int
}

func NewRegister(L int) *Register {
	return &Register{
		uniqRegs: make(map[string]struct{}, int(math.Pow(float64(2), float64(L)))),
		Digits:   make([]int, L),
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
		r.Digits[i] = rand.Intn(2)
	}

	var data string
	for i := 0; i < r.L; i++ {
		data += fmt.Sprint(r.Digits[i])
	}

	if err := WriteToFile("key.txt", data); err != nil {
		log.Fatal(err)
	}
}

func (r *Register) FeedBackFunc() int {
	var newDigit int
	for _, k_i := range r.polynome {
		newDigit ^= r.Digits[k_i-1]
	}

	rightLast := r.Digits[len(r.Digits)-1]
	r.Digits = append([]int{newDigit}, r.Digits[:len(r.Digits)-1]...)

	return rightLast
}

func WriteToFile(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(data); err != nil {
		return err
	}
	defer writer.Flush()

	return nil
}

func ReadFromFile(filename string) ([]int, error) {
	file, err := os.Open("key.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	
	key, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	result := make([]int, 0)
	for i := 0; i < len(key); i++ {
		result = append(result, int(key[i]-'0'))
	}

	return result, nil
}

func main() {
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
		if err != nil {
			log.Fatal(err)
		}
	} else {
		reg.GenRegister()
	}
	reg.ParsePoly(feedBackFunc)

	
	MSeq := make([]int, 0, int(math.Pow(float64(2), float64(reg.L))))
	
	for {
		MSeq = append(MSeq, reg.FeedBackFunc())
		if _, ok := reg.uniqRegs[reg.GetStringDigit()]; ok {
			break
		} else {
			reg.uniqRegs[reg.GetStringDigit()] = struct{}{}
		}
	}

	if N == -1 {
		fmt.Println(MSeq)
	} else {
		fmt.Println(MSeq[:N])
	}
}
