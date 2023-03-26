package cypher

import (
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"time"

	"crypto_task_3/utils"
)

func Run() {
	text, err := utils.ReadFromFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	e, d, n := utils.ReadKeys()
	E, _ := big.NewInt(0).SetString(e, 10)
	D, _ := big.NewInt(0).SetString(d, 10)
	N, _ := big.NewInt(0).SetString(n, 10)

	source := utils.GetBin(text)
	fmt.Println("-------------- Source Text --------------")
	M := make([]*big.Int, len(source))
	for i := 0; i < len(source); i++ {
		M[i], _ = big.NewInt(0).SetString(string(source[i]), 2)
		fmt.Println(M[i])
	}

	C := make([]*big.Int, len(M))
	for i := 0; i < len(M); i++ {
		C[i] = big.NewInt(0).Exp(M[i], E, N)
	}

	var encrypted string
	for i := 0; i < len(C); i++ {
		encrypted += C[i].String() + "\n"
	}

	utils.WriteToFile(encrypted, "encrypted.txt")

	R := make([]*big.Int, len(M))
	for i := 0; i < len(M); i++ {
		R[i] = big.NewInt(0).Exp(C[i], D, N)
	}

	var decrypted string
	for i := 0; i < len(R); i++ {
		decrypted += R[i].String() + "\n"
	}

	utils.WriteToFile(decrypted, "decrypted.txt")
}

func PollardAttack(n *big.Int) *big.Int {
	random := big.NewInt(0).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), n)
	k := big.NewInt(1)
	gcd := big.NewInt(0)

	for f := big.NewInt(1); f.Cmp(n) == -1; f = f.Add(f, big.NewInt(1)) {
		x := make([]*big.Int, 0, 100)
		x = append(x, random)
		z := 0

		i := big.NewInt(0).Add(big.NewInt(0).Exp(big.NewInt(2), k, nil), big.NewInt(1))
		for i.Cmp(big.NewInt(0).Add(big.NewInt(0).Exp(big.NewInt(2), k.Add(k, big.NewInt(1)), nil), big.NewInt(1))) == -1 {
			tmp := big.NewInt(0).Mod(big.NewInt(0).Add(big.NewInt(0).Exp(x[z], big.NewInt(2), nil), big.NewInt(1)), n)
			x = append(x, tmp)
			gcd = big.NewInt(0).GCD(nil, nil, n, big.NewInt(0).Abs(big.NewInt(0).Sub(x[0], x[z])))
			if gcd.Cmp(big.NewInt(1)) == 1 {
				return gcd
			}
			z++
		}
		random = x[z-1]
		k = k.Add(k, big.NewInt(1))
	}

	return nil
}
