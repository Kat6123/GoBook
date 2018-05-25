package main

import (
	"testing"
)

// Need I test for mass string if massUnit already pass tests?
func TestMass_String(t *testing.T) {
	tt := []struct {
		m   mass
		str string
	}{
		{mass{5, microgr}, "5 mcgr"},
		{mass{10, milligr}, "10 mlgr"},
		{mass{15, gram}, "15 gr"},
		{mass{20, kilogr}, "20 kg"},
		{mass{25, centner}, "25 centner"},
		{mass{30, ton}, "30 ton"},
		{mass{35, massUnit(123)}, "35 not defined"},
	}

	for _, tc := range tt {
		tr := tc.m.String()
		if tr != tc.str {
			t.Errorf("string should be %v; got %v", tc.str, tr)
		}
	}
}

func TestMassUnit_String(t *testing.T) {
	tt := []struct {
		unit massUnit
		str  string
	}{
		{microgr, "mcgr"},
		{milligr, "mlgr"},
		{gram, "gr"},
		{kilogr, "kg"},
		{centner, "centner"},
		{ton, "ton"},
		{massUnit(123), "not defined"},
	}

	for _, tc := range tt {
		tr := tc.unit.String()
		if tr != tc.str {
			t.Errorf("string should be %v; got %v", tc.str, tr)
		}
	}
}

func TestMass_convert(t *testing.T) {
	tt := []struct {
		m     mass
		unit  massUnit
		value float64
	}{
		{mass{5, gram}, microgr, 5e6},
		{mass{5, gram}, milligr, 5e3},
		{mass{5, gram}, gram, 5},
		{mass{5, gram}, kilogr, 5e-3},
		{mass{5, gram}, centner, 5e-5},
		{mass{5, gram}, ton, 5e-6},
		{mass{5, gram}, massUnit(123), 5. / 123},
	}

	for _, tc := range tt {
		m := tc.m
		m.convert(tc.unit)
		if m.value != tc.value {
			t.Errorf("converted value should be %v; got %v", tc.value, m.value)
		}
	}
}

func TestMass_toVolume(t *testing.T) {
	tt := []struct {
		m     mass
		unit  volumeUnit
		value float64
	}{
		{mass{5, kilogr}, millim3, 5e6},
		{mass{5, gram}, mll, 5.},
		{mass{5, kilogr}, litr, 5.},
		{mass{5, gram}, metr3, 5e-6},
	}

	for _, tc := range tt {
		v := tc.m.toVolume(tc.unit)
		// XXX should get 5 by get 5.0001
		if float32(v.value) != float32(tc.value) {
			t.Errorf("converted value of %s should be %v in %s; got %v", tc.m, tc.value, tc.unit, v.value)
		}
	}
}

func TestMass_convertTo(t *testing.T) {
	tt := []struct {
		m     mass
		unit  interface{}
		value interface{}
		err   string
	}{
		{mass{5, kilogr}, millim3,
			volume{5e6, millim3}, ""},
		{mass{5, gram}, milligr,
			mass{5e3, milligr}, ""},
		{mass{5, gram}, 5,
			nil, "cann't convert value of first unit type int"},
	}

	for _, tc := range tt {
		val, err := tc.m.convertTo(tc.unit)
		if err != nil{
			if err.Error() != tc.err{
				// XXX
				t.Errorf("wrong error msg %s", err.Error())
			}
			continue
		}

		switch v := val.(type) {
		case mass:
				if v != tc.value{
					t.Errorf("wrong mass value %s", v)
				}
		case volume:
			if v != tc.value{
				t.Errorf("wrong volume value %s", v)
			}
		default:
			t.Errorf("wrong type %T", v)
		}
	}
}
