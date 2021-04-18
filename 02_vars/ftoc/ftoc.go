package main

import (
	"02_vars/tempconv"
	"fmt"
)

func main() {
	const freezing, boiling = 32.0, 212.0
	fmt.Printf("%g F %g C\n", freezing, fToC(freezing))
	fmt.Printf("%g F %g C\n", boiling, fToC(boiling))
	fmt.Printf("%g F %g C\n", boiling, tempconv.FToC(boiling))
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Temperature of boil: %g F of %g C\n", f, c)

	fmt.Printf("%g\n", BoilingC-FreezingC)

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
