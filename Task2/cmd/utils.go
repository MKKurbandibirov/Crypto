package main

import (
	"bufio"
	"flag"
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

func ReadFromFile(filename string) ([]byte, error) {
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

	result := make([]byte, 0)
	for i := 0; i < len(key); i++ {
		result = append(result, byte(key[i]-'0'))
	}

	return result, nil
}

type App struct {
	task        int
	L           int
	poly        string
	N           int
	keyFile     string
	serialK     int
	serialAlpha float64
	pokerAlpha  float64
}

func FlagParse(app *App) {
	flag.IntVar(&app.task, "task", 1, "Register size")

	flag.IntVar(&app.L, "L", 4, "Register size")

	flag.StringVar(&app.poly, "poly", "x4+x1+1", "Polynom for feedback function")

	flag.IntVar(&app.N, "N", -1, "Result size")

	flag.StringVar(&app.keyFile, "file", "", "File to store key")

	flag.IntVar(&app.serialK, "serialK", 2, "Serial Test K")

	flag.Float64Var(&app.serialAlpha, "serialAlpha", 0, "Serial Test alpha")

	flag.Float64Var(&app.pokerAlpha, "pokerAlpha", 0, "Poker Test alpha")

	flag.Parse()
}
