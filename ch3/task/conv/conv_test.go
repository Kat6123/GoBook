package main

import (
	"testing"
)

func TestMass_convertTo(t *testing.T) {
	tt := []struct {
		testName       string
		m              mass
		unit           interface{}
		expectedValue  interface{}
		expectedErrMsg string
	}{
		{"mass to volume", mass{5, kilogr}, millim3,
			volume{5e6, millim3}, ""},
		{"mass to mass", mass{5, gram}, milligr,
			mass{5e3, milligr}, ""},
		{"mass to undefined unit", mass{5, gram}, 5,
			nil, "cann't convert value to unit type int"},
		{"mass to nil", mass{5, gram}, nil,
			nil, "cann't convert value to unit type <nil>"},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			val, err := tc.m.convertTo(tc.unit)
			// If return error with non nil value? How to test
			if err != nil {
				if err.Error() != tc.expectedErrMsg {
					t.Fatalf("error should be %q; get %q", tc.expectedErrMsg, err)
				}
			}

			// If val != nil?
			if val != nil {
				switch v := val.(type) {
				case mass:
					if v != tc.expectedValue {
						t.Fatalf("should be value %v of type %[1]T; get %v of type %[2]T", tc.expectedValue, v)
					}
				case volume:
					if v != tc.expectedValue {
						t.Fatalf("should be value %v of type %[1]T; get %v of type %[2]T", tc.expectedValue, v)
					}
				default:
					t.Fatalf("should be value %v of type %[1]T; get %v of type %[2]T", tc.expectedValue, v)
				}
			}
		})
	}
}

func TestVolume_convertTo(t *testing.T) {
	tt := []struct {
		testName       string
		m              volume
		unit           interface{}
		expectedValue  interface{}
		expectedErrMsg string
	}{
		{"volume to volume", volume{5, litr}, litr,
			volume{5, litr}, ""},
		{"volume to mass", volume{5, litr}, kilogr,
			mass{5, kilogr}, ""},
		{"volume to undefined unit", volume{5, millim3}, 5,
			nil, "cann't convert value to unit type int"},
		{"volume to nil", volume{5, millim3}, nil,
			nil, "cann't convert value to unit type <nil>"},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			val, err := tc.m.convertTo(tc.unit)
			if err != nil {
				if err.Error() != tc.expectedErrMsg {
					t.Fatalf("error should be %q; get %q", tc.expectedErrMsg, err)
				}
			}

			// If val != nil?
			if val != nil {
				switch v := val.(type) {
				case mass:
					if v != tc.expectedValue {
						t.Fatalf("should be value %v of type %[1]T; get %v of type %[2]T", tc.expectedValue, v)
					}
				case volume:
					if v != tc.expectedValue {
						t.Fatalf("should be value %v of type %[1]T; get %v of type %[2]T", tc.expectedValue, v)
					}
				default:
					t.Fatalf("should be value %v of type %[1]T; get %v of type %[2]T", tc.expectedValue, v)
				}
			}
		})
	}
}

func Test_convert(t *testing.T) {
	tt := []struct {
		testName       string
		v              interface{}
		unit           interface{}
		expectedValue  interface{}
		expectedErrMsg string
	}{
		{"volume to volume", volume{5, litr}, litr,
			volume{5, litr}, ""},
		{"undefined to mass", 5, kilogr,
			nil, "cann't convert value of type int"},
		{"volume to undefined unit", volume{5, millim3}, 5,
			nil, "cann't convert value to unit type int"},
		{"nil to nil", nil, nil,
			nil, "cann't convert value of type <nil>"},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			val, err := convert(tc.v, tc.unit)
			if err != nil {
				if err.Error() != tc.expectedErrMsg {
					t.Fatalf("error should be %q; get %q", tc.expectedErrMsg, err)
				}
			}

			// If val != nil?
			if val != nil {
				switch v := val.(type) {
				case mass:
					if v != tc.expectedValue {
						t.Fatalf("should be value %v of type %[1]T; get %v of type %[2]T", tc.expectedValue, v)
					}
				case volume:
					if v != tc.expectedValue {
						t.Fatalf("should be value %v of type %[1]T; get %v of type %[2]T", tc.expectedValue, v)
					}
				default:
					t.Fatalf("should be value %v of type %[1]T; get %v of type %[2]T", tc.expectedValue, v)
				}
			}
		})
	}
}
