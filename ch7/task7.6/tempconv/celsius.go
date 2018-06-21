package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64

const AbsoluteZeroC Celsius = -273.15

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (c *Celsius) Set(s string) (err error) {
	var unit string
	var value float64

	if _, err := fmt.Sscanf(s, "%f%s", &value, &unit); err != nil {
		return fmt.Errorf("scan %s", s)
	}

	switch unit {
	case "C", "°C":
		*c = Celsius(value)
		return nil
	case "F", "°F":
		*c, err = FToC(Fahrenheit(value))
		if err != nil {
			return fmt.Errorf("parse %s as Fahrenheit to set Celsius: %v", s, err)
		}
	case "K":
		*c, err = KToC(Kelvin(value))
		if err != nil {
			return fmt.Errorf("parse %s as Kelvin to set Celsius: %v", s, err)
		}
	default:
		return fmt.Errorf("wrong temperature type %q in %s", unit, s)
	}
	return
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	temp := value
	flag.CommandLine.Var(&temp, name, usage)
	return &temp
}
