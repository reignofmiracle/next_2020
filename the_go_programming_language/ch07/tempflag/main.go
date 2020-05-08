package main

import (
	"flag"
	"fmt"
)

// Celsius ...
type Celsius float64

// Fahrenheit ...
type Fahrenheit float64

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CToF ...
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC ...
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//CelsiusFlag ...
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temoerature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
