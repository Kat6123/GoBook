package main

import (
	"fmt"
)

const (
	micro = 1e-6
	milli = 1e-3
	kilo  = 1e3
	mega  = 1e6
	giga  = 1e9
)

// If  use float 64 error doesn't accumulate as in float32, but there was written in book that float32 is better
type massUnit float64

// iota, do not export
const (
	microgr massUnit = micro
	milligr massUnit = milli
	gram    massUnit = 1
	kilogr  massUnit = kilo
	centner massUnit = 100 * kilogr
	ton     massUnit = mega
)

func (unit massUnit) String() string {
	switch unit {
	case microgr:
		return "mcgr"
	case milligr:
		return "mlgr"
	case gram:
		return "gr"
	case kilogr:
		return "kg"
	case centner:
		return "centner"
	case ton:
		return "ton"
	}

	return "not defined"
}

type mass struct {
	value float64
	unit  massUnit
}

func (m mass) String() string {
	return fmt.Sprintf("%g %s", m.value, m.unit)
}

func (m *mass) convert(unit massUnit) {
	converted := massUnit(m.value) * m.unit / unit
	m.value = float64(converted)
	m.unit = unit
}

func (m mass) toVolume(unit volumeUnit) volume {
	m.convert(kilogr)
	v := volume{value: m.value / 1000, unit: metr3}
	v.convert(unit)
	return v
}

type volumeUnit float64

const (
	millim3 volumeUnit = 1 / giga
	mll     volumeUnit = 1 / mega
	litr    volumeUnit = 1 / kilo
	metr3   volumeUnit = 1
)

type volume struct {
	value float64
	unit  volumeUnit
}

func (v volume) String() string {
	return fmt.Sprintf("%g %s", v.value, v.unit)
}

func (unit volumeUnit) String() string {
	switch unit {
	case millim3:
		return "mm^3"
	case mll:
		return "mll"
	case litr:
		return "litr"
	case metr3:
		return "m^3"
	}

	return "not defined"
}

func (v *volume) convert(unit volumeUnit) {
	converted := volumeUnit(v.value) * v.unit / unit
	v.value = float64(converted)
	v.unit = unit
}

func (v volume) toMass(unit massUnit) mass {
	v.convert(metr3)
	m := mass{value: v.value * 1000, unit: kilogr}
	m.convert(unit)
	return m
}

func (m mass) convertTo(unit interface{}) (interface{}, error) {
	switch u := unit.(type) {
	case massUnit:
		// If mass to mass then return the sam value?
		m.convert(u)
		return m, nil
	case volumeUnit:
		return m.toVolume(u), nil
	default:
		return nil, fmt.Errorf("cann't convert value of first unit type %T", u)
	}
}

func (v volume) convertTo(unit interface{}) (interface{}, error) {
	switch u := unit.(type) {
	case massUnit:
		return v.toMass(u), nil
	case volumeUnit:
		v.convert(u)
		return v, nil
	default:
		return nil, fmt.Errorf("cann't convert value of first unit type %T", u)
	}
}

func convert(unit interface{}, value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case mass:
		return v.convertTo(unit)
	case volume:
		return v.convertTo(unit)
	default:
		return nil, fmt.Errorf("cann't convert value of second unit type %T", v)
	}
}

// 100% coverage
//func main() {
//	//m, err := Convert(kilogr, 1, L)
//	//if err != nil {
//	//	log.Fatalf("smth get wrong: %v", err)
//	//}
//	//fmt.Printf("%v", m)
//	n := mass{12, kilogr}
//	fmt.Println(n)
//	n.convert(gram)
//	fmt.Println(n)
//
//	v := volume{12, metr3}
//	fmt.Println(v)
//	v.convert(litr)
//	fmt.Println(v)
//	m := v.toMass(kilogr)
//	fmt.Println(m)
//
//	i, err := convert(kilogr, v)
//	if err != nil {
//		log.Fatalf("smth get wrong: %v", err)
//	}
//	fmt.Printf("%T %[1]v", i)
//}
