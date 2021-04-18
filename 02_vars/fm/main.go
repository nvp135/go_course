package main

import (
	"02_vars/lengthconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fm: %v\n", err)
			os.Exit(1)
		}
		f := lengthconv.Feet(l)
		m := lengthconv.Meter(l)
		fmt.Printf("%s = %s, %s = %s\n", f, lengthconv.CFtM(f), m, lengthconv.CMtF(m))
	}
}
