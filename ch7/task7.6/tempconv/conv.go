// Package tempconv defines:
// types for different temperatures;
// methods to use temperatures as command-line flags;
// methods to convert from one temperature type to another.
package tempconv

import "fmt"

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) (Fahrenheit, error) {
	if c < absoluteZeroC {
		return 0, fmt.Errorf("%s is lower than absolute zero", c)
	}
	return Fahrenheit(c*9/5 + 32), nil
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) (Celsius, error) {
	if f < absoluteZeroF {
		return 0, fmt.Errorf("%s is lower than absolute zero", f)
	}
	return Celsius((f - 32) * 5 / 9), nil
}

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) (Kelvin, error) {
	if c < absoluteZeroC {
		return 0, fmt.Errorf("%s is lower than absolute zero", c)
	}
	return Kelvin(c + 273.15), nil
}

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) (Celsius, error) {
	if k < absoluteZeroK {
		return 0, fmt.Errorf("%s is lower than absolute zero", k)
	}
	return Celsius(k - 273.15), nil
}

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) (Kelvin, error) {
	c, err := FToC(f)
	if err != nil {
		return 0, fmt.Errorf("convert from fahrenheit to celcius: %v", err)
	}

	k, err := CToK(c)
	if err != nil {
		return 0, fmt.Errorf("convert from celsius to kelvin: %v", err)
	}
	return k, err
}

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) (Fahrenheit, error) {
	c, err := KToC(k)
	if err != nil {
		return 0, fmt.Errorf("convert from kelvin to celcius: %v", err)
	}

	f, err := CToF(c)
	if err != nil {
		return 0, fmt.Errorf("convert from celcius to fahrenheit: %v", err)
	}
	return f, err
}
