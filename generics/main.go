package main

import "fmt"

// declare two non-generic functions
func SumInt(m map[string]int64) int64 {

	var s int64

	for _, v := range m {
		s += v
	}

	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}
	return s
}

/*
In this code, you:
Declare two functions to add together the values of a map and return the sum

 1. SumInts takes a map of string to int64 values
 2. SumFloats takes a map of string to float64 values
*/

/*  */

func main() {

	ints := map[string]int64{
		"first":  34,
		"second": 38,
		"third":  42,
	}

	floats := map[string]float64{
		"first":  34.10,
		"second": 38.28,
		"third":  42.65,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n", SumInt(ints), SumFloats(floats))
	fmt.Println("*************************generic********************")
	fmt.Printf("Generic Sums: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
}

/* Add a generic function to handle multiple types */
/* we'll add a single generic function that can receive a map containing either integer or float values,
effectively replacing the two functions you just wrote with a single function

*/

/* keep in mind that a type parameter must support all the operations the generic code is performing on it.
For example, if your function's code were to try to perform string operations (such as indexing) on a type parametr whose constraint included numeric types,
the code wouln't compile.
*/

type Number interface {
	int64 | float64
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, value := range m {
		s += value
	}

	return s
}

/*
1. Declare SumIntsOrFloats function with two type parameters(inside the square brackets), K and V and one argument that uses the type parameters, m of type map m[K]V.
The function returns a value of type V.
2. Specify for the K type parameter the type constraint comparable. Intended specifically for cases like these,
the comparable constraint is predeclared in Go. It allows any type whose values may be used as an operand of the comparision operators == and !=.

*******************************imp********************
Specify that the m argument is of type map[K]V,
where K and V are the types already specified for the type parameters.
Note , we know map[K]V is a valid  map type because K is a comrable type.
If we hadn't declared K comrable, the compiler would reject the reference to map[K]V
*/
