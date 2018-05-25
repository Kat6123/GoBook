package main

import "testing"

func TestMassUnit_String(t *testing.T) {
	tt := []struct {
		unit        massUnit
		expectedStr string
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
		if tr != tc.expectedStr {
			t.Errorf("string should be %v; got %v", tc.expectedStr, tr)
		}
	}
}

func TestVolumeUnit_String(t *testing.T) {
	tt := []struct {
		unit        volumeUnit
		expectedStr string
	}{
		{millim3, "mm^3"},
		{mll, "mll"},
		{litr, "litr"},
		{metr3, "m^3"},
		{volumeUnit(123), "not defined"},
	}

	for _, tc := range tt {
		tr := tc.unit.String()
		if tr != tc.expectedStr {
			t.Errorf("string should be %v; got %v", tc.expectedStr, tr)
		}
	}
}
