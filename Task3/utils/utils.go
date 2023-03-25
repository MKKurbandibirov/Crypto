package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadFromFile(filename string) (string, error) {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(text), nil
}

func ReadKeys() (E, D, N string) {
	public, err := ReadFromFile("public.txt")
	if err != nil {
		log.Fatal(err)
	}
	publKeys := strings.Split(public, "\n")
	E = strings.Split(publKeys[0], " ")[1]
	N = strings.Split(publKeys[1], " ")[1]

	private, err := ReadFromFile("private.txt")
	if err != nil {
		log.Fatal(err)
	}
	privKeys := strings.Split(private, "\n")
	D = strings.Split(privKeys[0], " ")[1]

	return E, D, N
}

func GetBin(text string) []string {
	tmp := make([]byte, 0, len(text)*8)
	for i := 0; i < len(text); i++ {
		var val = text[i]

		j := 0
		rev := make([]byte, 0, 8)
		for ; val > 0; j++ {
			rev = append(rev, val%2)
			val /= 2
		}
		for j := len(rev); j < 8; j++ {
			tmp = append(tmp, 0)
		}
		for j := len(rev) - 1; j >= 0; j-- {
			tmp = append(tmp, rev[j])
		}
	}

	cut := len(tmp) / 4
	tmp2 := make([][]byte, 0, 4)
	for i := 0; i < len(tmp); i += cut {
		tmp2 = append(tmp2, tmp[i:i+cut])
	}

	source := make([]string, 0, 4)
	for i := 0; i < len(tmp2); i++ {
		source = append(source, strings.ReplaceAll(strings.Trim(fmt.Sprint(tmp2[i]), "[]"), " ", ""))
	}



	return source
}

func WriteToFile(encryptedText, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	_, err = writer.WriteString(encryptedText)
	if err != nil {
		return err
	}

	return nil
}
