package main

import "testing"

func TestMass_String(t *testing.T) {
	tt := []struct {
		m           mass
		expectedStr string
	}{
		{mass{5, microgr}, "5 mcgr"},
		{mass{5000, milligr}, "5000 mlgr"},
		{mass{35, massUnit(123)}, "35 not defined"},
	}

	for _, tc := range tt {
		tr := tc.m.String()
		if tr != tc.expectedStr {
			t.Errorf("string should be %v; got %v", tc.expectedStr, tr)
		}
	}
}

func TestMass_convert(t *testing.T) {
	// I don't handle negative values
	tt := []struct {
		testName     string
		m            mass
		unit         massUnit
		expectedMass mass
	}{
		{"to microgr", mass{5, gram}, microgr,
			mass{5e6, microgr}},
		{"to milligr", mass{5, gram}, milligr,
			mass{5e3, milligr}},
		{"to gram", mass{5, gram}, gram,
			mass{5, gram}},
		{"to kilogr", mass{5, gram}, kilogr,
			mass{5e-3, kilogr}},
		{"to centner", mass{5, gram}, centner,
			mass{5e-5, centner}},
		{"to ton", mass{5, gram}, ton,
			mass{5e-6, ton}},
		{"to undefined unit", mass{5, gram}, massUnit(123),
			mass{5. / 123, massUnit(123)}},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			m := tc.m
			mConv := m.convert(tc.unit)
			if mConv != tc.expectedMass {
				t.Fatalf("converted value should be %v; got %v", tc.expectedMass, mConv)
			}
		})
	}
}

func TestMass_toVolume(t *testing.T) {
	tt := []struct {
		testName    string
		m           mass
		unit        volumeUnit
		expectedVol volume
	}{
		{"to millim3", mass{5, kilogr}, millim3,
			volume{5e6, millim3}},
		{"to mll", mass{5, gram}, mll,
			volume{5.000000000000001, mll}},
		{"to litr", mass{5, kilogr}, litr,
			volume{5, litr}},
		{"to metr3", mass{5e3, kilogr}, metr3,
			volume{5, metr3}},
		{"to undefined unit", mass{5, kilogr}, volumeUnit(123),
			volume{4.065040650406504e-05, volumeUnit(123)}},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			v := tc.m.toVolume(tc.unit)
			if v != tc.expectedVol {
				t.Fatalf("converted value of %s should be %s; got %s", tc.m, tc.expectedVol, v)
			}
		})
	}
}
