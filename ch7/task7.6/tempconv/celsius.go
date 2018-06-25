package tempconv

import (
	"flag"
	"fmt"
)

// Celsius type contains value of Celsius temp.
type Celsius float64

const (
	absoluteZeroC Celsius = -273.15
	celsius               = "C"
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// Set parses input string of temperature in Celsius, Kelvin on Fahrenheit and set value in Celsius.
func (c *Celsius) Set(s string) (err error) {
	var unit string
	var value float64

	if _, err = fmt.Sscanf(s, "%f%s", &value, &unit); err != nil {
		return fmt.Errorf("scan %s", s)
	}

	switch unit {
	case celsius, "°" + celsius:
		*c = Celsius(value)
		return nil
	case fahrenheit, "°" + fahrenheit:
		*c, err = FToC(Fahrenheit(value))
		if err != nil {
			return fmt.Errorf("parse %s as Fahrenheit to set Celsius: %v", s, err)
		}
	case kelvin:
		*c, err = KToC(Kelvin(value))
		if err != nil {
			return fmt.Errorf("parse %s as Kelvin to set Celsius: %v", s, err)
		}
	default:
		return fmt.Errorf("wrong temperature type %q in %s", unit, s)
	}
	return
}

// CelsiusFlag defines a tempconv.Celsius flag with specified name, default value, and usage string.
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	temp := value
	flag.CommandLine.Var(&temp, name, usage)
	return &temp
}
