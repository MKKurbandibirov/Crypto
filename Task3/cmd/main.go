package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"os/exec"
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

func GeneratePrime(l int64) *big.Int {
	length := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(l), nil)

	newP := big.NewInt(0).Add(length, big.NewInt(0).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), length))
	for !newP.ProbablyPrime(100) {
		newP.Add(newP, big.NewInt(1))
	}

	return newP
} 

func NewGenerator(L int64) *Generator {
	gen := &Generator{}

	for {
		newP := GeneratePrime(L)
		// newP = newP.Add(newP, big.NewInt(1))

		newQ := GeneratePrime(L)

		gen.p = newP
		gen.q = newQ
		gen.n = big.NewInt(0).Mul(newP, newQ)
		gen.fi = big.NewInt(0).Mul(big.NewInt(0).Sub(newP, big.NewInt(1)), big.NewInt(0).Sub(newQ, big.NewInt(1)))
		gen.e = big.NewInt(65537)
		gen.d = big.NewInt(0).ModInverse(gen.e, gen.fi)

		var e, d *big.Int = big.NewInt(0).Set(gen.e), big.NewInt(0).Set(gen.d)
		if big.NewInt(0).GCD(nil, nil, e, d).Cmp(big.NewInt(1)) == 0 {
			break
		}
	}

	return gen
}

func main() {
	defer out.Flush()

	var L int64
	flag.Int64Var(&L, "L", 64, "RSA key bit size")

	var task int
	flag.IntVar(&task, "task", 1, "Task number")

	flag.Parse()

	if task == 1 {
		gen := NewGenerator(L)

		publicKey := fmt.Sprintf("E: %s\nN: %s", gen.e, gen.n)
		privateKey := fmt.Sprintf("P: %s\nQ: %s\nD: %s\nN: %s", gen.p, gen.q, gen.d, gen.n)

		utils.WriteToFile(publicKey, "public.txt")
		utils.WriteToFile(privateKey, "private.txt")
	} else if task == 2 {
		cypher.Run(int(L) / 4)
	} else if task == 4 {
		var Ns string
		for i := 35; i < 60; {
			gen := NewGenerator(int64(i))
			
			Ns += fmt.Sprintf("%s\n", gen.n)
			
			if i <= 45 {
				i += 5
			} else {
				i++
			}
		}
		utils.WriteToFile(Ns, "N.txt")
		cmd := exec.Command("python3", "attack/attack.py")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	} else if task == 5 {
		L := 55
		var Ns string
		for r := 0.25; r <= 0.5; r += 0.025 {
			p := GeneratePrime(int64(r*float64(L)))
			q := GeneratePrime(int64((1-r)*float64(L)))
			n := big.NewInt(0).Mul(p, q)

			Ns += fmt.Sprintf("%s\n", n)
		}
		utils.WriteToFile(Ns, "N.txt")
		cmd := exec.Command("python3", "attack/attack.py")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
