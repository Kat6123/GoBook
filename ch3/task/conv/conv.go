package main

import (
	"fmt"
)

// convertTo returns interface with mass converted to mass or volume with error
func (m mass) convertTo(unit interface{}) (interface{}, error) {
	switch u := unit.(type) {
	case massUnit:
		return m.convert(u), nil
	case volumeUnit:
		return m.toVolume(u), nil
	default:
		return nil, fmt.Errorf("cann't convert value to unit type %T", u)
	}
}

// convertTo returns interface with volume converted to mass or volume with error
func (v volume) convertTo(unit interface{}) (interface{}, error) {
	switch u := unit.(type) {
	case massUnit:
		return v.toMass(u), nil
	case volumeUnit:
		return v.convert(u), nil
	default:
		return nil, fmt.Errorf("cann't convert value to unit type %T", u)
	}
}

func convert(value interface{}, unit interface{}) (interface{}, error) {
	switch v := value.(type) {
	case mass:
		return v.convertTo(unit)
	case volume:
		return v.convertTo(unit)
	default:
		return nil, fmt.Errorf("cann't convert value of type %T", v)
	}
}
