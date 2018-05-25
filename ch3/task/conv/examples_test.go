package main

import "fmt"

func ExampleMass_convert() {
	m := mass{5, gram}
	m.convert(microgr)
	fmt.Println(m)
	// Output:
	// 5e+06 mcgr
}
