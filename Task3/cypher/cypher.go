package cypher

import (
	"fmt"
	"log"
	"math/big"

	"crypto_task_3/utils"
)

func Run(cut int) {
	text, err := utils.ReadFromFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	e, d, n := utils.ReadKeys()
	E, _ := big.NewInt(0).SetString(e, 10)
	D, _ := big.NewInt(0).SetString(d, 10)
	N, _ := big.NewInt(0).SetString(n, 10)

	source := utils.GetBin(text, cut)
	fmt.Println("-------------- Source Text --------------")
	M := make([]*big.Int, len(source))
	for i := 0; i < len(source); i++ {
		M[i], _ = big.NewInt(0).SetString(string(source[i]), 2)
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
