// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Kat6123/GoBook/ch2/task2.1/tempconv"
	"log"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		// Why error is printed during the first iteration? as it's the second param, smth with log buffer?
		if err != nil {
			log.Printf("can't parse float: %v\n", err)
			continue
		}

		c := tempconv.Celsius(t)

		cTof, err := tempconv.CToF(c)
		if err != nil {
			log.Printf("%s not valid temperature", c)
			continue
		}
		fmt.Printf("%s: %s\n", c, cTof)

		cTok, err := tempconv.CToK(c)
		if err != nil {
			log.Printf("%s not valid temperature", c)
			continue
		}
		fmt.Printf("%s: %s\n", c, cTok)

	}
}

//!-
