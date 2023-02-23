package main

import (
	"flag"
	"fmt"

	"github.com/aniruddha2000/passpass/cmd"
)

func main() {
	var (
		len     int
		lower   bool
		upper   bool
		numner  bool
		special bool
	)

	flag.IntVar(&len, "len", 8, "Determine the length of the password")
	flag.BoolVar(&lower, "lower", true, "Include lower case letters")
	flag.BoolVar(&upper, "upper", false, "Include upper case letters")
	flag.BoolVar(&numner, "number", false, "Include number case letters")
	flag.BoolVar(&special, "special", false, "Include special case letters")

	flag.Parse()

	cfg := cmd.NewConfig(len, lower, upper, numner, special)
	res := cfg.Generate()
	fmt.Println(res)
}
