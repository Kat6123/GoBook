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

	"github.com/Kat6123/GoBook/tempconv"
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
		if !tempconv.ValidTemp(c){
			log.Printf("%s not valid temperature", c)
			continue
		}
		fmt.Printf("%s: %s, %s\n",
			c, tempconv.CToF(c), tempconv.CToK(c))

		k := tempconv.Kelvin(t)
		if !tempconv.ValidTemp(k){
			log.Printf("%s not valid temperature", k)
			continue
		}
		fmt.Printf("%s: %s, %s\n",
			k, tempconv.KToC(k), tempconv.KToF(k))

		f := tempconv.Fahrenheit(t)
		if !tempconv.ValidTemp(f){
			log.Printf("%s not valid temperature", f)
			continue
		}
		fmt.Printf("%s: %s, %s\n\n",
			f, tempconv.FToC(f), tempconv.FToK(f))
	}
}

//!-
