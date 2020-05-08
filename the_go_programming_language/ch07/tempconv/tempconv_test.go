package tempconv_test

import (
	"fmt"
	"testing"
)

type Celsius float64
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

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func TestSet(t *testing.T) {
	var c celsiusFlag
	c.Set("10C")
	fmt.Println(c)
	v := Celsius(10)
	if c.Celsius != v {
		t.Errorf("%v = %v", c.Celsius, v)
	}
}
