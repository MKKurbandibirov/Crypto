package main

import (
	"bufio"
	"io"
	"os"
)

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
	file, err := os.Open(filename)
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
