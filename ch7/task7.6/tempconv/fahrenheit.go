package tempconv

import (
	"flag"
	"fmt"
)

type Fahrenheit float64

const AbsoluteZeroF = -459.67

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (f *Fahrenheit) Set(s string) (err error) {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", value, unit)
	switch unit {
	case "F":
		*f = Fahrenheit(value)
		return nil
	case "K":
		*f, err = KToF(Kelvin(value))
		if err != nil {
			return fmt.Errorf("parse %s as Kelvin to set Fahrenheit: %v", s, err)
		}
	case "C", "°C":
		*f, err = CToF(Celsius(value))
		if err != nil {
			return fmt.Errorf("parse %s as Celsius to set Fahrenheit: %v", s, err)
		}
	default:
		return fmt.Errorf("wrong temperature %g", value)
	}
	return
}

func FahrenheitFlag(name string, value Fahrenheit, usage string) *Fahrenheit {
	temp := value
	flag.CommandLine.Var(&temp, name, usage)
	return &temp
}
