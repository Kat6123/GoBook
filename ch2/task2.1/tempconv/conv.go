// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

import "fmt"

// Where I should check valid temperature?

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) (Fahrenheit, error) {
	if c < AbsoluteZeroC {
		return 0, fmt.Errorf("%s is lower than absolute zero\n", c)
	}
	return Fahrenheit(c*9/5 + 32), nil
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) (Celsius, error) {
	if f < AbsoluteZeroF {
		return 0, fmt.Errorf("%s is lower than absolute zero\n", f)
	}
	return Celsius((f - 32) * 5 / 9), nil
}

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) (Kelvin, error) {
	if c < AbsoluteZeroC {
		return 0, fmt.Errorf("%s is lower than absolute zero\n", c)
	}
	return Kelvin(c + 273.15), nil
}

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) (Celsius, error) {
	if k < AbsoluteZeroK {
		return 0, fmt.Errorf("%s is lower than absolute zero\n", k)
	}
	return Celsius(k - 273.15), nil
}

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) (Kelvin, error) {
	c, err := FToC(f)
	if err != nil {
		return 0, fmt.Errorf("error while converting from fahrenheit to celcius: %v", err)
	}
	// Add wrap
	k, err := CToK(c)
	if err != nil {
		return 0, fmt.Errorf("error while converting from celsius to kelvin: %v", err)
	}
	return k, err
}

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) (Fahrenheit, error) {
	c, err := KToC(k)
	if err != nil {
		return 0, fmt.Errorf("error while converting from kelvin to celcius: %v", err)
	}
	// dual error check?
	f, err := CToF(c)
	if err != nil {
		return 0, fmt.Errorf("error while converting from celcius to fahrenheit: %v", err)
	}
	return f, err
}

//!-
