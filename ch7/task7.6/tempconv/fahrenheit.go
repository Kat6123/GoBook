package tempconv

import (
	"flag"
	"fmt"
)

// Fahrenheit type contains value of Fahrenheit temp.
type Fahrenheit float64

const (
	absoluteZeroF = -459.67
	fahrenheit    = "F"
)

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

// Set parses input string of temperature in Celsius, Kelvin on Fahrenheit and set value in Fahrenheit.
func (f *Fahrenheit) Set(s string) (err error) {
	var unit string
	var value float64

	if _, err = fmt.Sscanf(s, "%f%s", &value, &unit); err != nil {
		return fmt.Errorf("scan %s", s)
	}

	switch unit {
	case fahrenheit, "°" + fahrenheit:
		*f = Fahrenheit(value)
		return nil
	case kelvin:
		*f, err = KToF(Kelvin(value))
		if err != nil {
			return fmt.Errorf("parse %s as Kelvin to set Fahrenheit: %v", s, err)
		}
	case celsius, "°" + celsius:
		*f, err = CToF(Celsius(value))
		if err != nil {
			return fmt.Errorf("parse %s as Celsius to set Fahrenheit: %v", s, err)
		}
	default:
		return fmt.Errorf("wrong temperature type %q in %s", unit, s)
	}
	return
}

// FahrenheitFlag defines a tempconv.Fahrenheit flag with specified name, default value, and usage string.
func FahrenheitFlag(name string, value Fahrenheit, usage string) *Fahrenheit {
	temp := value
	flag.CommandLine.Var(&temp, name, usage)
	return &temp
}
