package cmd

import (
	"math/rand"
	"time"
)

const (
	LoweCaseValues    string = "abcdefghijklmnopqrstuvwxyz"
	UpperCaseValues   string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumberValues      string = "0123456789"
	SpecialCharValues string = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

type Config struct {
	Length      int
	Lowercase   bool
	Uppercase   bool
	Number      bool
	SpecialChar bool
}

func NewConfig(len int, l, u, n, s bool) *Config {
	return &Config{
		Length:      len,
		Lowercase:   l,
		Uppercase:   u,
		Number:      n,
		SpecialChar: s,
	}
}

func (c *Config) Generate() string {
	seedRandom := rand.New(rand.NewSource(time.Now().UnixNano()))

	letters := LoweCaseValues

	if c.Uppercase {
		letters += UpperCaseValues
	}
	if c.Number {
		letters += NumberValues
	}
	if c.SpecialChar {
		letters += SpecialCharValues
	}

	res := make([]byte, c.Length)
	for i := range res {
		res[i] = letters[seedRandom.Intn(len(letters))]
	}

	return string(res)
}
