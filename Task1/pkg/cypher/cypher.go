package cypher

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	ALPH = "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Cypher struct {
	ReplaceMap         map[rune]rune
	ReversedReplaceMap map[rune]rune
}

func NewCypher() *Cypher {
	return &Cypher{
		ReplaceMap:         make(map[rune]rune, 100),
		ReversedReplaceMap: make(map[rune]rune, 100),
	}
}


func (c *Cypher) GenerateMapping() {
	alph := []rune(ALPH)
	
	for _, ch := range ALPH {
		ind1 := rand.Intn(len(alph))
		val1 := alph[ind1]
		for val1 == ch {
			ind1 = rand.Intn(len(alph))
			val1 = alph[ind1]
		}
		alph = append(alph[:ind1], alph[ind1+1:]...)
		
		c.ReplaceMap[ch] = val1
		c.ReversedReplaceMap[val1] = ch
	}
}

func (c *Cypher) String() string {
	var sb strings.Builder

	for _, ch := range ALPH {
		if ch != ' ' {
			sb.WriteString(fmt.Sprintf("%s : %s\n", string(ch), string(c.ReplaceMap[ch])))
		}
	}

	sb.WriteString(fmt.Sprintf("  : %s\n", string(c.ReplaceMap[' '])))

	return sb.String()
}

func (c *Cypher) Encrypt(text string) string {
	data := []rune(text)

	for i := 0; i < len(data); i++ {
		if _, ok := c.ReplaceMap[data[i]]; ok {
			data[i] = c.ReplaceMap[data[i]]
		}
	}

	return string(data)
}

func (c *Cypher) Decrypt(text string) string {
	data := []rune(text)

	for i := 0; i < len(data); i++ {
		if _, ok := c.ReversedReplaceMap[data[i]]; ok {
			data[i] = c.ReversedReplaceMap[data[i]]
		}
	}

	return string(data)
}
