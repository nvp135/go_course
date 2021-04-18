// Package tempconv performs Celsius and Fahrenheit temperature computations
package tempconv

import (
	"fmt"
)

// Celsius type represents celsius
type Celsius float64

// Fahrenheit type represents fahrenheit
type Fahrenheit float64

// Kelvin type represents kelvin
type Kelvin float64

// Constants
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	boilingF      float64 = 212.0
)

// CToF convert celsius to fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK convert celsius to kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToC convert fahrenheit to celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK convert fahrenheit to kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin((f-32)*5/9 + 273.15) }

// KToC convert kelvin to celsius
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF convert kelvin to celsius
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*9/5 + 32) }

func (c Celsius) String() string    { return fmt.Sprintf("%g C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g F", f) }
