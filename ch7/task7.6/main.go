package main

import (
	"flag"
	"fmt"

	"github.com/kat6123/GoBook/ch7/task7.6/tempconv"
)

var (
	c = tempconv.CelsiusFlag("celsius", 20, "set temperature in any unit")
	k = tempconv.KelvinFlag("kelvin", 20, "set temperature in any unit")
	f = tempconv.FahrenheitFlag("fahrenheit", 20, "set temperature in any unit")
)

func main() {
	flag.Parse()
	fmt.Println(c, k, f)
}
