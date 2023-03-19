package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

func ReadFromFile(filename string) (string, error) {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(text), nil
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
