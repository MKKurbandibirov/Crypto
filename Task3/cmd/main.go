package main

import (
	"bufio"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"time"
)

var (
	out = bufio.NewWriter(os.Stdout)
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Generator struct {
	p, q *big.Int
	n    *big.Int
	fi   *big.Int
	e    *big.Int
	d    *big.Int
}

func NewGenerator(L int64) *Generator {
	length := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(L), nil)

	newP := big.NewInt(0).Add(length, big.NewInt(0).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), length))
	for !newP.ProbablyPrime(100) {
		newP.Add(newP, big.NewInt(1))
	}

	newQ := big.NewInt(0).Add(length, big.NewInt(0).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), length))
	for !newQ.ProbablyPrime(100) {
		newQ.Add(newQ, big.NewInt(1))
	}

	gen := &Generator{
		p:  newP,
		q:  newQ,
		n:  big.NewInt(0).Mul(newP, newQ),
		fi: big.NewInt(0).Mul(newP.Sub(newP, big.NewInt(1)), newQ.Sub(newQ, big.NewInt(1))),
		e:  big.NewInt(65537),
	}

	gen.d = big.NewInt(0).ModInverse(gen.e, gen.fi)

	return gen
}

func main() {
	defer out.Flush()

	gen := NewGenerator(10)

	fmt.Fprintln(out, gen.p)
	fmt.Fprintln(out, gen.q)
	fmt.Fprintln(out, gen.n)
	fmt.Fprintln(out, gen.fi)
	fmt.Fprintln(out, gen.d)
}
