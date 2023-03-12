package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sync"

	"crypto_task_2_sub_task_1/cypher"
	"crypto_task_2_sub_task_1/tests"
)

var out = bufio.NewWriter(os.Stdout)

func RunTests(MSeq []byte, app *App) {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		serialSeq := make([]byte, len(MSeq))
		copy(serialSeq, MSeq)
		serial := tests.NewSerialTest(serialSeq, app.serialK)
		serial.Test(out, app.serialAlpha, &mutex)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		pokerSeq := make([]byte, len(MSeq))
		copy(pokerSeq, MSeq)
		poker := tests.NewPokerTest()
		poker.Test(out, pokerSeq, &mutex, app.pokerAlpha)
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		corrSeq := make([]byte, len(MSeq))
		copy(corrSeq, MSeq)
		corr := tests.NewCorrTest()
		corr.Test(corrSeq, out, &mutex)
	}()

	wg.Wait()
}

func main() {
	defer out.Flush()

	app := &App{}
	FlagParse(app)

	reg := NewRegister(app.L)
	if app.keyFile != "" {
		var err error
		reg.Digits, err = ReadFromFile(app.keyFile)
		reg.L = len(reg.Digits)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		reg.GenRegister()
	}
	reg.ParsePoly(app.poly)

	MSeq := make([]byte, 0, int(math.Pow(float64(2), float64(reg.L))))

	for {
		var val = reg.GetStringDigit()
		if _, ok := reg.uniqRegs[val]; ok {
			break
		} else {
			MSeq = append(MSeq, reg.FeedBackFunc())
			reg.uniqRegs[val] = struct{}{}
		}
	}

	if app.task == 1 {
		if app.N == -1 {
			for i := 0; i < len(MSeq); i++ {
				fmt.Fprint(out, MSeq[i], " ")
			}
		} else {
			for i := 0; i < app.N; i++ {
				fmt.Fprint(out, MSeq[i], " ")
			}
		}
		fmt.Fprintln(out)
	} else if app.task == 2 {
		RunTests(MSeq, app)
	} else if app.task == 3 {
		if err := cypher.Run(MSeq, "binfile"); err != nil {
			log.Fatal(err)
		}
	} else if app.task == 4 {
		bytes, err := cypher.ReadBinFile("binfile")
		if err != nil {
			log.Fatal(err)
		}

		source := make([]byte, 0, len(bytes)*8)
		for i := 0; i < len(bytes); i++ {
			bin := make([]byte, 0, 8)
			for bytes[i] > 0 {
				bit := bytes[i] % 2
				bytes[i] /= 2
				bin = append(bin, byte(bit))
			}

			for len(bin) < 8 {
				bin = append(bin, 0)
			}

			for i := 7; i >= 0; i-- {
				source = append(source, bin[i])
			}
		}
		
		RunTests(source, app)
	}

}
