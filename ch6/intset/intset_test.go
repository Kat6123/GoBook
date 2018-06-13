// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "fmt"

func ExampleIntSet_UnionWith() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
}

func ExampleIntSet_String() {
	var x IntSet

	fmt.Println(&x)

	x.Add(1)
	x.Add(2)
	x.Add(3)
	fmt.Println(&x)

	// Output:
	// {}
	// {1 2 3}
}

func ExampleIntSet_Has() {
	var x IntSet

	fmt.Println(x.Has(5))

	x.Add(5)
	fmt.Println(x.Has(5))

	// Output:
	// false
	// true
}

func ExampleIntSet_Add() {
	var x IntSet

	fmt.Println(x.String()) // Empty set.

	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	// Output:
	// {}
	// {1 9 144}
}

func ExampleIntSet_Len() {
	var x IntSet
	fmt.Println(x.Len()) // Empty set. Length: 0.

	x.Add(1)
	fmt.Println(x.Len()) // { 1 }. Length: 1.

	x.Add(1) // { 1 }. Length: 1.
	fmt.Println(x.Len())

	x.Add(153) // { 1 153 }. Length: 2.
	fmt.Println(x.Len())

	// Output:
	// 0
	// 1
	// 1
	// 2
}

func ExampleIntSet_Remove() {
	var x IntSet // Empty set: {}.

	x.Remove(10)
	fmt.Println(&x) // Remove from empty set shouldn't change set: {}.

	x.Add(1)
	x.Add(2)
	x.Add(3) // {1 2 3}

	x.Remove(2)
	fmt.Println(&x) // {1 3}

	x.Remove(100000)
	fmt.Println(&x) // {1 3}

	// Output:
	// {}
	// {1 3}
	// {1 3}
}

func ExampleIntSet_Clear() {
	var x IntSet

	x.Clear()
	fmt.Println(&x) // Clear empty set should return empty set.

	x.Add(1)
	x.Add(2)
	x.Add(3) // Set: {1 2 3}
	x.Clear()

	fmt.Println(&x) // Empty set.

	// Output:
	// {}
	// {}
}

func ExampleIntSet_Copy() {
	var x, y *IntSet
	x = &IntSet{}

	y = x.Copy()

	fmt.Println(x) // Empty set.
	fmt.Println(y)
	fmt.Println(x == y) // Pointers should be different.

	x.Add(1)
	x.Add(12)
	x.Add(122)

	y = x.Copy()

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x == y) // Pointers should be different.

	// Output:
	// {}
	// {}
	// false
	// {1 12 122}
	// {1 12 122}
	// false
}
