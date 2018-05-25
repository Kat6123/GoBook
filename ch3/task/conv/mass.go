package main

import "fmt"

// Where is better to define type
type mass struct {
	value float64
	unit  massUnit
}

func (m mass) String() string {
	return fmt.Sprintf("%g %s", m.value, m.unit)
}

// convert returns mass in new units
func (m mass) convert(unit massUnit) mass {
	converted := massUnit(m.value) * m.unit / unit
	m.value = float64(converted)
	m.unit = unit
	return m
}

// toVolume returns volume converted from mass
func (m mass) toVolume(unit volumeUnit) volume {
	mConv := m.convert(kilogr)
	v := volume{value: mConv.value / 1000, unit: metr3}
	return v.convert(unit)
}
