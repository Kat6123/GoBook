package main

import "fmt"

type byteSize float64

// Why in float64 precision is more than in int64? -
const (
	_ = iota
	kb
	mb
	gb
	tb
	pb
	eb
	zb
	yb
)

func main() {
	fmt.Println(kb)
	fmt.Println(mb)
	fmt.Println(gb)
	fmt.Println(tb)
	fmt.Println(pb)
	fmt.Println(eb)
	fmt.Println(zb)
	fmt.Println(yb)
}
