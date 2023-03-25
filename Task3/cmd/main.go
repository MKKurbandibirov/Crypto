package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"time"

	"crypto_task_3/cypher"
	"crypto_task_3/utils"
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
	// newP = newP.Add(newP, big.NewInt(1))

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

	var L int64
	flag.Int64Var(&L, "L", 128, "RSA key bit size")

	var task int
	flag.IntVar(&task, "task", 1, "Task number")

	flag.Parse()

	if task == 1 {
		gen := NewGenerator(L)

		publicKey := fmt.Sprintf("E: %s\nN: %s", gen.e, gen.n)
		privateKey := fmt.Sprintf("D: %s\nN: %s", gen.d, gen.n)

		utils.WriteToFile(publicKey, "public.txt")
		utils.WriteToFile(privateKey, "private.txt")
	} else if task == 2 {
		cypher.Run()
	} else if task == 3 {
		
	}
}
