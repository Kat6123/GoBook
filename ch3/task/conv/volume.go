package main

import "fmt"

type volume struct {
	value float64
	unit  volumeUnit
}

func (v volume) String() string {
	return fmt.Sprintf("%g %s", v.value, v.unit)
}

// convert returns volume in new units
func (v volume) convert(unit volumeUnit) volume {
	converted := volumeUnit(v.value) * v.unit / unit
	v.value = float64(converted)
	v.unit = unit
	return v
}

// toMass returns mass converted from volume
func (v volume) toMass(unit massUnit) mass {
	vConv := v.convert(metr3)
	m := mass{value: vConv.value * 1000, unit: kilogr}
	return m.convert(unit)
}
