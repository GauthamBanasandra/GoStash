package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	AbsoluteZeroK Kelvin  = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g⁰C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g⁰F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g⁰K", k)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + AbsoluteZeroC)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - AbsoluteZeroK)
}

func FtoK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}

func KtoF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

func main() {
	fmt.Println(FtoK(24))
}
