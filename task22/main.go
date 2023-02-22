package main

import (
	"fmt"
	"math/big"
)

func main() {
	var aStr string = "342213324453455443423424223"
	var bStr string = "10101011101110102929101291"

	a := new(big.Float)
	a.SetString(aStr)

	b := new(big.Float)
	b.SetString(bStr)

	r := new(big.Float)

	fmt.Printf("%s * %s = %f\n", aStr, bStr, r.Mul(a, b))
	fmt.Printf("%s / %s = %f\n", aStr, bStr, r.Quo(a, b))
	fmt.Printf("%s + %s = %f\n", aStr, bStr, r.Add(a, b))
	fmt.Printf("%s - %s = %f\n", aStr, bStr, r.Sub(a, b))
}
