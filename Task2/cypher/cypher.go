package cypher

import (
	"encoding/binary"
	"math"
	"os"
)

func ReadBinFile(filename string) ([]uint8, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	var bytes = make([]uint8, stats.Size())
	if err := binary.Read(file, binary.BigEndian, &bytes); err != nil {
		return nil, err
	}

	return bytes, nil
}

func Encode(source, key []uint8, outfile string) ([]uint8, error) {
	result := make([]uint8, 0, len(source))
	for i := 0; i < len(source); i++ {
		result = append(result, source[i]^key[i])
	}

	encrFile, err := os.Create(outfile)
	if err != nil {
		return nil, err
	}
	if err := binary.Write(encrFile, binary.BigEndian, result); err != nil {
		return nil, err
	}

	return result, nil
}

func Run(MSeq []byte, filename string) error {
	source, err := ReadBinFile(filename)
	if err != nil {
		return err
	}

	key := make([]uint8, 0, len(MSeq)/8)
	for i := 0; i < len(MSeq); i += 8 {
		tmp := MSeq[i : i+8]
		var n uint8
		for j := 0; j < len(tmp); j++ {
			n += uint8(math.Pow(2, float64(j)) * float64(tmp[j]))
		}

		key = append(key, n)
	}
	key = key[:len(source)]

	encrypted, err := Encode(source, key, "encrypted.txt")
	if err != nil {
		return err
	}

	_, err = Encode(encrypted, key, "decrypted.txt")
	if err != nil {
		return err
	}

	return nil
}
