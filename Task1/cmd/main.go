package main

import (
	"crypto_task_1/pkg/cypher"
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "\u001b[31m[ERROR]\u001b[0m\t", log.Lshortfile)

func main() {
	if len(os.Args) == 1 {
		logger.Fatalln("Don't have a text file!")
	}

	text, err := ReadFromFile(os.Args[1])
	if err != nil {
		logger.Fatalf("Couldn't open a file: %s: %s", os.Args[1], err)
	}

	fmt.Println(text)
	fmt.Println()

	cypher := cypher.NewCypher()
	cypher.GenerateMapping()

	encryptedText := cypher.Encrypt(text)
	if err := WriteToFile(encryptedText, "encrypted.txt"); err != nil {
		logger.Fatalln(err)
	}

	if err := WriteToFile(cypher.String(), "key.txt"); err != nil {
		logger.Fatalln(err)
	}

	decryptedText := cypher.Decrypt(encryptedText)
	fmt.Println(decryptedText)
	if err := WriteToFile(decryptedText, "decrypted.txt"); err != nil {
		logger.Fatalln(err)
	}
}
