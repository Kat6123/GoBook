// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// Where I should check valid temperature?

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {return Fahrenheit(c*9/5 + 32)}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

func ValidTemp(t interface{}) bool {
	switch v := t.(type) {
	case Celsius:
		if v < AbsoluteZeroC{return false}
	case Fahrenheit:
		if v < AbsoluteZeroF{return false}
	case Kelvin:
		if v < AbsoluteZeroK{return false}
	default:
		return false
	}

	return true
}

//!-
