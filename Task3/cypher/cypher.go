package cypher

import (
	"fmt"
	"log"
	"math"
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
	M := make([]*big.Int, len(source))
	// var numText string
	for i := 0; i < len(source); i++ {
		M[i], _ = big.NewInt(0).SetString(string(source[i]), 2)
		// fmt.Println(M[i])
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

	// var bin = make([]byte, 0, len(R)*8)
	// for i := 0; i < len(R); i++ {
	// 	bin = append(bin, utils.GetBits(string(R[i].Bytes()))...)
	// }

	// res := make([]byte, 0, len(bin)/8)
	// for i := 0; i < len(bin); i += 8 {
	// 	tmp := bin[i : i+8]
	// 	var n byte
	// 	for j := 0; j < len(tmp); j++ {
	// 		n += byte(math.Pow(2, float64(j)) * float64(tmp[j]))
	// 	}

	// 	res = append(res, n)
	// }
	// // res = res[:len(source)]

	// fmt.Println(string(res))
}
