package tempconv

import (
	"flag"
	"fmt"
)

// Kelvin type contains value of Kelvin temp.
type Kelvin float64

const (
	absoluteZeroK = 0
	kelvin        = "K"
)

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

// Set parses input string of temperature in Celsius, Kelvin on Fahrenheit and set value in Kelvin.
func (k *Kelvin) Set(s string) (err error) {
	var unit string
	var value float64

	if _, err = fmt.Sscanf(s, "%f%s", &value, &unit); err != nil {
		return fmt.Errorf("scan %s", s)
	}

	switch unit {
	case kelvin:
		*k = Kelvin(value)
		return nil
	case fahrenheit, "°" + fahrenheit:
		*k, err = FToK(Fahrenheit(value))
		if err != nil {
			return fmt.Errorf("parse %s as Fahrenheit to set Kelvin: %v", s, err)
		}
	case celsius, "°" + celsius:
		*k, err = CToK(Celsius(value))
		if err != nil {
			return fmt.Errorf("parse %s as Celsius to set Kelvin: %v", s, err)
		}
	default:
		return fmt.Errorf("wrong temperature type %q in %s", unit, s)
	}
	return
}

// KelvinFlag defines a tempconv.Kelvin flag with specified name, default value, and usage string.
func KelvinFlag(name string, value Kelvin, usage string) *Kelvin {
	temp := value
	flag.CommandLine.Var(&temp, name, usage)
	return &temp
}
