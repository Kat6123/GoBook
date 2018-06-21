package tempconv

import (
	"flag"
	"fmt"
)

type Kelvin float64

const AbsoluteZeroK = 0

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

func (k *Kelvin) Set(s string) (err error) {
	var unit string
	var value float64

	if _, err := fmt.Sscanf(s, "%f%s", &value, &unit); err != nil {
		return fmt.Errorf("scan %s", s)
	}

	switch unit {
	case "K":
		*k = Kelvin(value)
		return nil
	case "F", "°F":
		*k, err = FToK(Fahrenheit(value))
		if err != nil {
			return fmt.Errorf("parse %s as Fahrenheit to set Kelvin: %v", s, err)
		}
	case "C", "°C":
		*k, err = CToK(Celsius(value))
		if err != nil {
			return fmt.Errorf("parse %s as Celsius to set Kelvin: %v", s, err)
		}
	default:
		return fmt.Errorf("wrong temperature type %q in %s", unit, s)
	}
	return
}

func KelvinFlag(name string, value Kelvin, usage string) *Kelvin {
	temp := value
	flag.CommandLine.Var(&temp, name, usage)
	return &temp
}
